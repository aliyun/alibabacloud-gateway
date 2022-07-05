<?php

// This file is auto-generated, don't edit it. Thanks.
namespace Darabonba\GatewayPds;

use Darabonba\GatewaySpi\Client as DarabonbaGatewaySpiClient;
use AlibabaCloud\Tea\Tea;
use AlibabaCloud\Tea\Utils\Utils;
use AlibabaCloud\OpenApiUtil\OpenApiUtilClient;
use AlibabaCloud\Tea\Exception\TeaError;
use AlibabaCloud\Darabonba\String\StringUtil;
use AlibabaCloud\Darabonba\SignatureUtil\SignatureUtil;
use AlibabaCloud\Darabonba\EncodeUtil\EncodeUtil;
use AlibabaCloud\Darabonba\MapUtil\MapUtil;
use AlibabaCloud\Darabonba\ArrayUtil\ArrayUtil;

use Darabonba\GatewaySpi\Models\InterceptorContext;
use Darabonba\GatewaySpi\Models\AttributeMap;

class Client extends DarabonbaGatewaySpiClient {
    public function __construct(){
        parent::__construct();
    }

    /**
     * @param InterceptorContext $context
     * @param AttributeMap $attributeMap
     * @return void
     */
    public function modifyConfiguration($context, $attributeMap){
        $config = $context->configuration;
        $config->endpoint = $this->getEndpoint($config->network, $config->endpoint);
    }

    /**
     * @param InterceptorContext $context
     * @param AttributeMap $attributeMap
     * @return void
     */
    public function modifyRequest($context, $attributeMap){
        $request = $context->request;
        $config = $context->configuration;
        $request->headers = Tea::merge([
            "date" => Utils::getDateUTCString(),
            "host" => $config->endpoint,
            "x-acs-version" => $request->version,
            "x-acs-action" => $request->action,
            "user-agent" => $request->userAgent,
            "x-acs-signature-nonce" => Utils::getNonce(),
            "x-acs-signature-method" => "HMAC-SHA1",
            "x-acs-signature-version" => "1.0",
            "accept" => "application/json"
        ], $request->headers);
        if (!Utils::isUnset($request->stream)) {
            $tmp = Utils::readAsBytes($request->stream);
            $request->stream = $tmp;
            $request->headers["content-type"] = "application/octet-stream";
        }
        else {
            if (!Utils::isUnset($request->body)) {
                if (Utils::equalString($request->reqBodyType, "json")) {
                    $jsonObj = Utils::toJSONString($request->body);
                    $request->stream = $jsonObj;
                    $request->headers["content-type"] = "application/json; charset=utf-8";
                }
                else {
                    $m = Utils::assertAsMap($request->body);
                    $formObj = OpenApiUtilClient::toForm($m);
                    $request->stream = $formObj;
                    $request->headers["content-type"] = "application/x-www-form-urlencoded";
                }
            }
        }
        if (!Utils::equalString($request->authType, "Anonymous")) {
            $credential = $request->credential;
            $accessKeyId = $credential->getAccessKeyId();
            $accessKeySecret = $credential->getAccessKeySecret();
            $securityToken = $credential->getSecurityToken();
            $bearerToken = $credential->getBearerToken();
            if (!Utils::empty_($bearerToken)) {
                $request->headers["x-acs-bearer-token"] = $bearerToken;
                $request->headers["Authorization"] = "Bearer " . $bearerToken . "";
            }
            else {
                if (!Utils::empty_($securityToken)) {
                    $request->headers["x-acs-security-token"] = $securityToken;
                }
                $request->headers["Authorization"] = $this->getAuthorization($request->pathname, $request->method, $request->query, $request->headers, $accessKeyId, $accessKeySecret);
            }
        }
    }

    /**
     * @param InterceptorContext $context
     * @param AttributeMap $attributeMap
     * @return void
     * @throws TeaError
     */
    public function modifyResponse($context, $attributeMap){
        $request = $context->request;
        $response = $context->response;
        if (Utils::is4xx($response->statusCode) || Utils::is5xx($response->statusCode)) {
            $_res = Utils::readAsJSON($response->body);
            $err = Utils::assertAsMap($_res);
            $requestId = $this->defaultAny(@$err["RequestId"], @$err["requestId"]);
            @$err["statusCode"] = $response->statusCode;
            throw new TeaError([
                "code" => "" . (string) ($this->defaultAny(@$err["Code"], @$err["code"])) . "",
                "message" => "code: " . (string) ($response->statusCode) . ", " . (string) ($this->defaultAny(@$err["Message"], @$err["message"])) . " request id: " . (string) ($requestId) . "",
                "data" => $err
            ]);
        }
        if (!Utils::isUnset($response->body) && !Utils::equalNumber($response->statusCode, 204)) {
            if (Utils::equalString($request->bodyType, "binary")) {
                $response->deserializedBody = $response->body;
            }
            else if (Utils::equalString($request->bodyType, "byte")) {
                $byt = Utils::readAsBytes($response->body);
                $response->deserializedBody = $byt;
            }
            else if (Utils::equalString($request->bodyType, "string")) {
                $str = Utils::readAsString($response->body);
                $response->deserializedBody = $str;
            }
            else if (Utils::equalString($request->bodyType, "json")) {
                $response->deserializedBody = Utils::readAsJSON($response->body);
            }
            else {
                $response->deserializedBody = Utils::readAsString($response->body);
            }
        }
    }

    /**
     * @param string $network
     * @param string $endpoint
     * @return string
     */
    public function getEndpoint($network, $endpoint){
        $realEndpoint = "api.aliyunpds.com";
        if (!Utils::empty_($endpoint)) {
            $realEndpoint = $endpoint;
        }
        if (!Utils::empty_($network) && StringUtil::equals($network, "vpc")) {
            $realEndpoint = StringUtil::replace($realEndpoint, "api.aliyunpds.com", "api-vpc.aliyunpds.com", null);
            $realEndpoint = StringUtil::replace($realEndpoint, "admin.aliyunpds.com", "admin-vpc.aliyunpds.com", null);
        }
        return $realEndpoint;
    }

    /**
     * @param mixed $inputValue
     * @param mixed $defaultValue
     * @return any
     */
    public function defaultAny($inputValue, $defaultValue){
        if (Utils::isUnset($inputValue)) {
            return $defaultValue;
        }
        return $inputValue;
    }

    /**
     * @param string $pathname
     * @param string $method
     * @param string[] $query
     * @param string[] $headers
     * @param string $ak
     * @param string $secret
     * @return string
     */
    public function getAuthorization($pathname, $method, $query, $headers, $ak, $secret){
        $signature = $this->getSignature($pathname, $method, $query, $headers, $secret);
        return "acs " . $ak . ":" . $signature . "";
    }

    /**
     * @param string $pathname
     * @param string $method
     * @param string[] $query
     * @param string[] $headers
     * @param string $secret
     * @return string
     */
    public function getSignature($pathname, $method, $query, $headers, $secret){
        $stringToSign = "";
        $canonicalizedResource = $this->buildCanonicalizedResource($pathname, $query);
        $canonicalizedHeaders = $this->buildCanonicalizedHeaders($headers);
        $stringToSign = "" . $method . "\n" . $canonicalizedHeaders . "" . $canonicalizedResource . "";
        $signature = SignatureUtil::HmacSHA1Sign($stringToSign, $secret);
        return EncodeUtil::base64EncodeToString($signature);
    }

    /**
     * @param string $pathname
     * @param string[] $query
     * @return string
     */
    public function buildCanonicalizedResource($pathname, $query){
        $canonicalizedResource = $pathname;
        if (!Utils::isUnset($query)) {
            $queryArray = MapUtil::keySet($query);
            $sortedQueryArray = ArrayUtil::ascSort($queryArray);
            $separator = "?";
            foreach($sortedQueryArray as $key){
                $canonicalizedResource = "" . $canonicalizedResource . "" . $separator . "" . $key . "";
                if (!Utils::empty_(@$query[$key])) {
                    $canonicalizedResource = "" . $canonicalizedResource . "=" . @$query[$key] . "";
                }
                $separator = "&";
            }
        }
        return $canonicalizedResource;
    }

    /**
     * @param string[] $headers
     * @return string
     */
    public function buildCanonicalizedHeaders($headers){
        $accept = @$headers["accept"];
        if (Utils::isUnset($accept)) {
            $accept = "";
        }
        $contentMd5 = @$headers["content-md5"];
        if (Utils::isUnset($contentMd5)) {
            $contentMd5 = "";
        }
        $contentType = @$headers["content-type"];
        if (Utils::isUnset($contentType)) {
            $contentType = "";
        }
        $date = @$headers["date"];
        if (Utils::isUnset($date)) {
            $date = "";
        }
        $canonicalizedHeaders = "" . $accept . "\n" . $contentMd5 . "\n" . $contentType . "\n" . $date . "\n";
        $sortedHeaders = $this->getSignedHeaders($headers);
        foreach($sortedHeaders as $header){
            $value = @$headers[$header];
            $valueTrim = StringUtil::trim($value);
            $canonicalizedHeaders = "" . $canonicalizedHeaders . "" . $header . ":" . $valueTrim . "\n";
        }
        return $canonicalizedHeaders;
    }

    /**
     * @param string[] $headers
     * @return array
     */
    public function getSignedHeaders($headers){
        $headersArray = MapUtil::keySet($headers);
        $sortedHeadersArray = ArrayUtil::ascSort($headersArray);
        $tmp = "";
        $separator = "";
        foreach($sortedHeadersArray as $key){
            $lowerKey = StringUtil::toLower($key);
            if (StringUtil::hasPrefix($lowerKey, "x-acs-")) {
                if (!StringUtil::contains($tmp, $lowerKey)) {
                    $tmp = "" . $tmp . "" . $separator . "" . $lowerKey . "";
                    $separator = ";";
                }
            }
        }
        return StringUtil::split($tmp, ";", 0);
    }
}
