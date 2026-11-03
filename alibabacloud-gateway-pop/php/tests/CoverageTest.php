<?php

namespace Darabonba\GatewayPop\Tests;

use AlibabaCloud\Credentials\Credential\CredentialModel;
use AlibabaCloud\Tea\Exception\TeaError;
use Darabonba\GatewayPop\Client;
use Darabonba\GatewaySpi\Models\AttributeMap;
use Darabonba\GatewaySpi\Models\InterceptorContext;
use Darabonba\GatewaySpi\Models\InterceptorContext\configuration;
use Darabonba\GatewaySpi\Models\InterceptorContext\request;
use Darabonba\GatewaySpi\Models\InterceptorContext\response;
use PHPUnit\Framework\TestCase;

class MemoryStream
{
    public $data;

    public function __construct($data)
    {
        $this->data = $data;
    }

    public function isSeekable()
    {
        return true;
    }

    public function rewind()
    {
    }

    public function getContents()
    {
        return $this->data;
    }
}

class BoomStream
{
    public function isSeekable()
    {
        return false;
    }

    public function rewind()
    {
    }

    public function getContents()
    {
        throw new \RuntimeException('read failed');
    }
}

class FixedCredential
{
    public $model;
    public $error;

    public function getCredential()
    {
        if ($this->error) {
            throw $this->error;
        }

        return $this->model;
    }
}

final class CoverageTest extends TestCase
{
    private function model($type)
    {
        $model = new CredentialModel();
        $model->type = $type;
        $model->accessKeyId = 'ak';
        $model->accessKeySecret = 'secret';
        $model->providerName = 'static_ak';

        return $model;
    }

    private function credential($model)
    {
        $credential = new FixedCredential();
        $credential->model = $model;

        return $credential;
    }

    private function request($credential)
    {
        $request = new request();
        $request->headers = array();
        $request->query = array();
        $request->pathname = '/';
        $request->productId = 'ecs';
        $request->action = 'DescribeRegions';
        $request->version = '2014-05-26';
        $request->protocol = 'HTTPS';
        $request->method = 'POST';
        $request->authType = 'AK';
        $request->bodyType = 'json';
        $request->reqBodyType = 'json';
        $request->credential = $credential;
        $request->userAgent = 'test-agent';

        return $request;
    }

    private function configuration($endpoint = 'ecs.cn-hangzhou.aliyuncs.com')
    {
        $config = new configuration();
        $config->regionId = 'cn-hangzhou';
        $config->endpoint = $endpoint;

        return $config;
    }

    private function context($request, $config)
    {
        $context = new InterceptorContext();
        $context->request = $request;
        $context->configuration = $config;
        $context->response = new response();

        return $context;
    }

    public function testEndpointAndConfiguration()
    {
        $client = new Client();
        $this->assertEquals('ecs.example.com', $client->getEndpoint('ecs', 'cn-hangzhou', 'regional', '', '', null, 'ecs.example.com'));
        $this->assertEquals('ecs.mapped.example.com', $client->getEndpoint('ecs', 'cn-hangzhou', 'regional', '', '', array('cn-hangzhou' => 'ecs.mapped.example.com'), null));
        $this->assertEquals('ecs.cn-hangzhou.aliyuncs.com', $client->getEndpoint('Ecs', 'cn-hangzhou', 'regional', '', '', null, null));
        $this->assertEquals('ecs.aliyuncs.com', $client->getEndpoint('Ecs', null, 'central', 'public', null, null, ''));
        try {
            $client->getEndpoint('Ecs', '', 'regional', '', '', null, null);
            $this->fail();
        } catch (\Exception $e) {
        }

        $built = new configuration();
        $built->regionId = 'cn-hangzhou';
        $built->endpointRule = 'regional';
        $built->network = '';
        $built->suffix = '';
        $attributes = new AttributeMap();
        $attributes->key = array(
            'signatureTypePrefix' => 'ACS4-',
            'signPrefix' => 'aliyun_v4',
            'endpointSuffix' => 'aliyuncs.com',
        );
        $client->modifyConfiguration($this->context($this->request($this->credential($this->model('access_key'))), $built), $attributes);
        $this->assertEquals('ecs.cn-hangzhou.aliyuncs.com', $built->endpoint);

        $kept = $this->configuration('already.example.com');
        $client->modifyConfiguration($this->context($this->request($this->credential($this->model('access_key'))), $kept), new AttributeMap());
        $this->assertEquals('already.example.com', $kept->endpoint);
    }

    public function testCanonicalAndSigning()
    {
        $client = new Client();
        $this->assertEquals('', $client->buildCanonicalizedResource(null));
        $resource = $client->buildCanonicalizedResource(array('b' => 'x y', 'a' => 'm*n', 'empty' => ''));
        $this->assertTrue(strpos($resource, 'a=m%2An') !== false);
        $this->assertTrue(strpos($resource, 'b=x%20y') !== false);
        $canonical = $client->buildCanonicalizedHeaders(array(
            'Host' => '  example.com  ',
            'X-Acs-Action' => 'One',
            'x-acs-action' => 'Two',
        ));
        $this->assertTrue(strpos($canonical, 'host:example.com') !== false);
        $this->assertEquals('fallback', $client->defaultAny(null, 'fallback'));
        $this->assertEquals('kept', $client->defaultAny('kept', 'fallback'));
        $signingkey = $client->getSigningkey('ACS4-HMAC-SHA256', 'secret', 'ecs', 'cn-hangzhou', '20240101');
        $this->assertNotEmpty($signingkey);
        $headers = array('host' => 'example.com', 'x-acs-date' => '2024-01-01T00:00:00Z');
        $signature = $client->getSignature('', 'GET', array('q' => 'a b'), $headers, 'ACS4-HMAC-SHA256', 'payload', $signingkey);
        $this->assertEquals(64, strlen($signature));
        $authorization = $client->getAuthorization('/', 'GET', null, $headers, 'ACS4-HMAC-SHA256', 'payload', 'ak', $signingkey, 'ecs', 'cn-hangzhou', '20240101');
        $this->assertTrue(strpos($authorization, 'ACS4-HMAC-SHA256 Credential=ak/20240101/cn-hangzhou/ecs/aliyun_v4_request,SignedHeaders=') !== false);
    }

    public function testModifyRequest()
    {
        $client = new Client();
        $json = $this->request($this->credential($this->model('access_key')));
        $json->body = array('name' => 'a b');
        $json->query = array('InstanceId' => 'i-1');
        $client->modifyRequest($this->context($json, $this->configuration()), new AttributeMap());
        $this->assertEquals('application/json; charset=utf-8', $json->headers['content-type']);
        $this->assertEquals('DescribeRegions', $json->headers['x-acs-action']);
        $this->assertTrue(strpos($json->headers['Authorization'], 'ACS4-HMAC-SHA256 Credential=ak/') !== false);
        $this->assertEquals('static_ak', $json->headers['x-acs-credentials-provider']);

        $bytes = $this->request($this->credential($this->model('access_key')));
        $bytes->reqBodyType = 'byte';
        $bytes->body = array_values(unpack('C*', 'hello'));
        $bytes->action = '';
        $client->modifyRequest($this->context($bytes, $this->configuration()), new AttributeMap());
        $this->assertArrayNotHasKey('x-acs-action', $bytes->headers);

        $form = $this->request($this->credential($this->model('access_key')));
        $form->reqBodyType = 'form';
        $form->body = array('b' => 'x y', 'a' => array('m', 'n'));
        $client->modifyRequest($this->context($form, $this->configuration()), new AttributeMap());
        $this->assertEquals('application/x-www-form-urlencoded', $form->headers['content-type']);

        $stream = $this->request($this->credential($this->model('access_key')));
        $stream->stream = new MemoryStream('stream-body');
        $client->modifyRequest($this->context($stream, $this->configuration()), new AttributeMap());
        $this->assertEquals('application/octet-stream', $stream->headers['content-type']);

        $client->modifyRequest($this->context($this->request($this->credential($this->model('access_key'))), $this->configuration()), new AttributeMap());

        $sm3 = $this->request($this->credential($this->model('access_key')));
        $sm3->signatureAlgorithm = 'ACS4-HMAC-SM3';
        $sm3->body = array('k' => 'v');
        $client->modifyRequest($this->context($sm3, $this->configuration()), new AttributeMap());
        $this->assertNotEmpty($sm3->headers['x-acs-content-sm3']);

        $bearerModel = $this->model('bearer');
        $bearerModel->bearerToken = 'bearer-token';
        $bearerModel->providerName = 'bearer';
        $bearer = $this->request($this->credential($bearerModel));
        $client->modifyRequest($this->context($bearer, $this->configuration()), new AttributeMap());
        $this->assertEquals('Bearer bearer-token', $bearer->headers['Authorization']);

        $stsModel = $this->model('sts');
        $stsModel->accessKeyId = 'sts-ak';
        $stsModel->securityToken = 'token';
        $stsModel->providerName = '';
        $sts = $this->request($this->credential($stsModel));
        $client->modifyRequest($this->context($sts, $this->configuration()), new AttributeMap());
        $this->assertEquals('token', $sts->headers['x-acs-security-token']);

        $idModel = $this->model('id_token');
        $idModel->securityToken = 'id-token';
        $idModel->providerName = null;
        $idToken = $this->request($this->credential($idModel));
        $client->modifyRequest($this->context($idToken, $this->configuration()), new AttributeMap());
        $this->assertEquals('id-token', $idToken->headers['x-acs-zero-trust-idtoken']);

        $anonymous = $this->request(null);
        $anonymous->authType = 'Anonymous';
        $client->modifyRequest($this->context($anonymous, $this->configuration()), new AttributeMap());
        $this->assertArrayNotHasKey('Authorization', $anonymous->headers);

        try {
            $client->modifyRequest($this->context($this->request(null), $this->configuration()), new AttributeMap());
            $this->fail();
        } catch (TeaError $e) {
            $this->assertEquals('ParameterMissing', $e->getErrorInfo()['code']);
        }

        $broken = new FixedCredential();
        $broken->error = new \RuntimeException('credential down');
        try {
            $client->modifyRequest($this->context($this->request($broken), $this->configuration()), new AttributeMap());
            $this->fail();
        } catch (\RuntimeException $e) {
        }

        $badBytes = $this->request($this->credential($this->model('access_key')));
        $badBytes->reqBodyType = 'byte';
        $badBytes->body = 'not-bytes';
        try {
            $client->modifyRequest($this->context($badBytes, $this->configuration()), new AttributeMap());
            $this->fail();
        } catch (\Exception $e) {
        }

        $badStream = $this->request($this->credential($this->model('access_key')));
        $badStream->stream = new BoomStream();
        try {
            $client->modifyRequest($this->context($badStream, $this->configuration()), new AttributeMap());
            $this->fail();
        } catch (\Exception $e) {
        }
    }

    private function responseContext($status, $contentType, $body, $bodyType)
    {
        $context = $this->context($this->request($this->credential($this->model('access_key'))), $this->configuration());
        $context->request->bodyType = $bodyType;
        $context->response->statusCode = $status;
        $context->response->headers = array();
        if ($contentType !== null) {
            $context->response->headers['content-type'] = $contentType;
        }
        $context->response->body = new MemoryStream($body);

        return $context;
    }

    public function testModifyResponse()
    {
        $client = new Client();
        try {
            $client->modifyResponse($this->responseContext(400, 'application/json', '{"Code":"Invalid","Message":"no","RequestId":"req-1","Description":"d"}', 'json'), new AttributeMap());
            $this->fail();
        } catch (TeaError $e) {
            $this->assertEquals('Invalid', $e->getErrorInfo()['code']);
            $this->assertTrue(strpos($e->getMessage(), 'request id: req-1') !== false);
        }

        $xml = '<ErrorResponse><Error><Code>XmlCode</Code><Message>xml</Message><RequestId>xml-id</RequestId><Description>xd</Description></Error></ErrorResponse>';
        try {
            $client->modifyResponse($this->responseContext(500, 'text/xml', $xml, 'xml'), new AttributeMap());
            $this->fail();
        } catch (TeaError $e) {
            $this->assertEquals('XmlCode', $e->getErrorInfo()['code']);
        }

        $client->modifyResponse($this->responseContext(204, null, '', 'string'), new AttributeMap());
        $binary = $this->responseContext(200, null, 'raw', 'binary');
        $client->modifyResponse($binary, new AttributeMap());
        $this->assertNotNull($binary->response->deserializedBody);
        $text = $this->responseContext(200, null, 'hello', 'string');
        $client->modifyResponse($text, new AttributeMap());
        $this->assertEquals('hello', $text->response->deserializedBody);
        $parsed = $this->responseContext(200, null, '{"ok":true}', 'json');
        $client->modifyResponse($parsed, new AttributeMap());
        $this->assertEquals(true, $parsed->response->deserializedBody['ok']);
        $array = $this->responseContext(200, null, '[1,2]', 'array');
        $client->modifyResponse($array, new AttributeMap());
        $this->assertCount(2, $array->response->deserializedBody);
        $other = $this->responseContext(200, null, 'plain', 'other');
        $client->modifyResponse($other, new AttributeMap());
        $this->assertEquals('plain', $other->response->deserializedBody);
        $bytes = $this->responseContext(200, null, 'raw', 'byte');
        $client->modifyResponse($bytes, new AttributeMap());
        $this->assertNotEmpty($bytes->response->deserializedBody);
    }
}
