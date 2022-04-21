<?php

// This file is auto-generated, don't edit it. Thanks.
namespace Darabonba\GatewayPop;

use Darabonba\GatewaySpi\Client as DarabonbaGatewaySpiClient;
use AlibabaCloud\Tea\Tea;
use AlibabaCloud\OpenApiUtil\OpenApiUtilClient;
use AlibabaCloud\Tea\Utils\Utils;
use AlibabaCloud\Darabonba\EncodeUtil\EncodeUtil;
use AlibabaCloud\Tea\Exception\TeaError;
use AlibabaCloud\Endpoint\Endpoint;
use AlibabaCloud\Darabonba\ArrayUtil\ArrayUtil;
use AlibabaCloud\Darabonba\String\StringUtil;
use AlibabaCloud\Darabonba\SignatureUtil\SignatureUtil;
use AlibabaCloud\Darabonba\MapUtil\MapUtil;

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
        $request = $context->request;
        $config = $context->configuration;
        $config->endpoint = $this->getEndpoint($request->productId, $config->regionId, $config->endpointRule, $config->network, $config->suffix, $config->endpointMap, $config->endpoint);
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
            "host" => $config->endpoint,
            "x-acs-version" => $request->version,
            "x-acs-action" => $request->action,
            "user-agent" => $request->userAgent,
            "x-acs-date" => OpenApiUtilClient::getTimestamp(),
            "x-acs-signature-nonce" => Utils::getNonce(),
            "accept" => "application/json"
        ], $request->headers);
        $signatureAlgorithm = "ACS3-HMAC-SHA256";
        if (!Utils::isUnset($request->signatureAlgorithm)) {
            $signatureAlgorithm = $request->signatureAlgorithm;
        }
        $hashedRequestPayload = EncodeUtil::hexEncode(EncodeUtil::hash(Utils::toBytes(""), $signatureAlgorithm));
        if (!Utils::isUnset($request->stream)) {
            $tmp = Utils::readAsBytes($request->stream);
            $hashedRequestPayload = EncodeUtil::hexEncode(EncodeUtil::hash($tmp, $signatureAlgorithm));
            $request->stream = $tmp;
            $request->headers["content-type"] = "application/octet-stream";
        }
        else {
            if (!Utils::isUnset($request->body)) {
                if (Utils::equalString($request->reqBodyType, "json")) {
                    $jsonObj = Utils::toJSONString($request->body);
                    $hashedRequestPayload = EncodeUtil::hexEncode(EncodeUtil::hash(Utils::toBytes($jsonObj), $signatureAlgorithm));
                    $request->stream = $jsonObj;
                    $request->headers["content-type"] = "application/json; charset=utf-8";
                }
                else {
                    $m = Utils::assertAsMap($request->body);
                    $formObj = OpenApiUtilClient::toForm($m);
                    $hashedRequestPayload = EncodeUtil::hexEncode(EncodeUtil::hash(Utils::toBytes($formObj), $signatureAlgorithm));
                    $request->stream = $formObj;
                    $request->headers["content-type"] = "application/x-www-form-urlencoded";
                }
            }
        }
        $request->headers["x-acs-content-sha256"] = $hashedRequestPayload;
        if (!Utils::equalString($request->authType, "Anonymous")) {
            $credential = $request->credential;
            $accessKeyId = $credential->getAccessKeyId();
            $accessKeySecret = $credential->getAccessKeySecret();
            $securityToken = $credential->getSecurityToken();
            if (!Utils::empty_($securityToken)) {
                $request->headers["x-acs-accesskey-id"] = $accessKeyId;
                $request->headers["x-acs-security-token"] = $securityToken;
            }
            $request->headers["Authorization"] = $this->getAuthorization($request->pathname, $request->method, $request->query, $request->headers, $signatureAlgorithm, $hashedRequestPayload, $accessKeyId, $accessKeySecret);
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
            throw new TeaError([
                "code" => "" . (string) ($this->defaultAny(@$err["Code"], @$err["code"])) . "",
                "message" => "code: " . (string) ($response->statusCode) . ", " . (string) ($this->defaultAny(@$err["Message"], @$err["message"])) . " request id: " . (string) ($requestId) . "",
                "data" => $err
            ]);
        }
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
            $obj = Utils::readAsJSON($response->body);
            $res = Utils::assertAsMap($obj);
            $response->deserializedBody = $res;
        }
        else if (Utils::equalString($request->bodyType, "array")) {
            $arr = Utils::readAsJSON($response->body);
            $response->deserializedBody = $arr;
        }
        else {
            $response->deserializedBody = Utils::readAsString($response->body);
        }
    }

    /**
     * @param string $productId
     * @param string $regionId
     * @param string $endpointRule
     * @param string $network
     * @param string $suffix
     * @param string[] $endpointMap
     * @param string $endpoint
     * @return string
     */
    public function getEndpoint($productId, $regionId, $endpointRule, $network, $suffix, $endpointMap, $endpoint){
        if (!Utils::empty_($endpoint)) {
            return $endpoint;
        }
        if (!Utils::isUnset($endpointMap) && !Utils::empty_(@$endpointMap[$regionId])) {
            return @$endpointMap[$regionId];
        }
        return Endpoint::getEndpointRules($productId, $regionId, $endpointRule, $network, $suffix);
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
     * @param string $signatureAlgorithm
     * @param string $payload
     * @param string $ak
     * @param string $secret
     * @return string
     */
    public function getAuthorization($pathname, $method, $query, $headers, $signatureAlgorithm, $payload, $ak, $secret){
        $signature = $this->getSignature($pathname, $method, $query, $headers, $signatureAlgorithm, $payload, $secret);
        $signedHeaders = $this->getSignedHeaders($headers);
        $signedHeadersStr = ArrayUtil::join($signedHeaders, ";");
        return "" . $signatureAlgorithm . "  Credential=" . $ak . ",SignedHeaders=" . $signedHeadersStr . ",Signature=" . $signature . "";
    }

    /**
     * @param string $pathname
     * @param string $method
     * @param string[] $query
     * @param string[] $headers
     * @param string $signatureAlgorithm
     * @param string $payload
     * @param string $secret
     * @return string
     */
    public function getSignature($pathname, $method, $query, $headers, $signatureAlgorithm, $payload, $secret){
        $canonicalURI = "/";
        if (!Utils::empty_($pathname)) {
            $canonicalURI = $pathname;
        }
        $stringToSign = "";
        $canonicalizedResource = $this->buildCanonicalizedResource($query);
        $canonicalizedHeaders = $this->buildCanonicalizedHeaders($headers);
        $signedHeaders = $this->getSignedHeaders($headers);
        $signedHeadersStr = ArrayUtil::join($signedHeaders, ";");
        $stringToSign = "" . $method . "\n" . $canonicalURI . "\n" . $canonicalizedResource . "\n" . $canonicalizedHeaders . "\n" . $signedHeadersStr . "\n" . $payload . "";
        $hex = EncodeUtil::hexEncode(EncodeUtil::hash(Utils::toBytes($stringToSign), $signatureAlgorithm));
        $stringToSign = "" . $signatureAlgorithm . "\n" . $hex . "";
        $signature = Utils::toBytes("");
        if (StringUtil::equals($signatureAlgorithm, "ACS3-HMAC-SHA256")) {
            $signature = SignatureUtil::HmacSHA256Sign($stringToSign, $secret);
        }
        else if (StringUtil::equals($signatureAlgorithm, "ACS3-HMAC-SM3")) {
            $signature = SignatureUtil::HmacSM3Sign($stringToSign, $secret);
        }
        else if (StringUtil::equals($signatureAlgorithm, "ACS3-RSA-SHA256")) {
            $signature = SignatureUtil::SHA256withRSASign($stringToSign, $secret);
        }
        return EncodeUtil::hexEncode($signature);
    }

    /**
     * @param string[] $query
     * @return string
     */
    public function buildCanonicalizedResource($query){
        $canonicalizedResource = "";
        if (!Utils::isUnset($query)) {
            $queryArray = MapUtil::keySet($query);
            $sortedQueryArray = ArrayUtil::ascSort($queryArray);
            $separator = "";
            foreach($sortedQueryArray as $key){
                $canonicalizedResource = "" . $canonicalizedResource . "" . $separator . "" . EncodeUtil::percentEncode($key) . "";
                if (!Utils::empty_(@$query[$key])) {
                    $canonicalizedResource = "" . $canonicalizedResource . "=" . EncodeUtil::percentEncode(@$query[$key]) . "";
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
        $canonicalizedHeaders = "";
        $sortedHeaders = $this->getSignedHeaders($headers);
        foreach($sortedHeaders as $header){
            $canonicalizedHeaders = "" . $canonicalizedHeaders . "" . $header . ":" . StringUtil::trim(@$headers[$header]) . "\n";
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
            if (StringUtil::hasPrefix($lowerKey, "x-acs-") || StringUtil::equals($lowerKey, "host") || StringUtil::equals($lowerKey, "content-type")) {
                if (!StringUtil::contains($tmp, $lowerKey)) {
                    $tmp = "" . $tmp . "" . $separator . "" . $lowerKey . "";
                    $separator = ";";
                }
            }
        }
        return StringUtil::split($tmp, ";", 0);
    }
}
