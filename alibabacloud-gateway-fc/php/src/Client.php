<?php

// This file is auto-generated, don't edit it. Thanks.
namespace Darabonba\GatewayFc;

use Darabonba\GatewaySpi\Client as DarabonbaGatewaySpiClient;
use AlibabaCloud\Darabonba\String\StringUtil;
use AlibabaCloud\Tea\Exception\TeaError;
use AlibabaCloud\Tea\Utils\Utils;
use AlibabaCloud\Endpoint\Endpoint;
use AlibabaCloud\Tea\Tea;
use AlibabaCloud\Darabonba\SignatureUtil\Signer;
use AlibabaCloud\Darabonba\EncodeUtil\Encoder;
use AlibabaCloud\OpenApiUtil\OpenApiUtilClient;
use AlibabaCloud\Darabonba\Array_\ArrayUtil;
use AlibabaCloud\Darabonba\Map\MapUtil;

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
        $config = $context->configuration;
        if (!StringUtil::hasSuffix($config->endpoint, "aliyuncs.com")) {
            $this->signRequestForFc($context);
        }
        else {
            $this->signRequestForPop($context);
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
        $config = $context->configuration;
        $response = $context->response;
        if (Utils::is4xx($response->statusCode) || Utils::is5xx($response->statusCode)) {
            $_headers = Utils::assertAsMap($response->headers);
            $_res = Utils::readAsJSON($response->body);
            $err = Utils::assertAsMap($_res);
            throw new TeaError([
                "code" => @$err["ErrorCode"],
                "message" => "code: " . (string) ($response->statusCode) . ", " . (string) (@$err["ErrorMessage"]) . " request id: " . (string) (@$_headers["x-fc-request-id"]) . "",
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
     * @param InterceptorContext $context
     * @return void
     */
    public function signRequestForFc($context){
        $request = $context->request;
        $config = $context->configuration;
        $request->headers = Tea::merge([
            "host" => $config->endpoint,
            "date" => Utils::getDateUTCString(),
            "accept" => "application/json",
            "user-agent" => $request->userAgent
        ], $request->headers);
        $request->headers["content-type"] = "application/json";
        if (!Utils::isUnset($request->stream)) {
            $tmp = Utils::readAsBytes($request->stream);
            $request->stream = $tmp;
            $request->headers["content-type"] = "application/octet-stream";
            $request->headers["content-md5"] = Encoder::base64EncodeToString(Signer::MD5Sign(Utils::toString($tmp)));
        }
        else {
            if (!Utils::isUnset($request->body)) {
                if (Utils::equalString($request->reqBodyType, "json")) {
                    $jsonObj = Utils::toJSONString($request->body);
                    $request->stream = $jsonObj;
                    $request->headers["content-type"] = "application/json";
                    $request->headers["content-md5"] = Encoder::base64EncodeToString(Signer::MD5Sign($jsonObj));
                }
                else {
                    $m = Utils::assertAsMap($request->body);
                    $formObj = OpenApiUtilClient::toForm($m);
                    $request->stream = $formObj;
                    $request->headers["content-type"] = "application/x-www-form-urlencoded";
                    $request->headers["content-md5"] = Encoder::base64EncodeToString(Signer::MD5Sign($formObj));
                }
            }
        }
        $credential = $request->credential;
        $accessKeyId = $credential->getAccessKeyId();
        $accessKeySecret = $credential->getAccessKeySecret();
        $securityToken = $credential->getSecurityToken();
        if (!Utils::empty_($securityToken)) {
            $request->headers["x-fc-security-token"] = $securityToken;
        }
        $request->headers["Authorization"] = $this->getAuthorizationForFc($request->pathname, $request->method, $request->query, $request->headers, $accessKeyId, $accessKeySecret);
    }

    /**
     * @param InterceptorContext $context
     * @return void
     */
    public function signRequestForPop($context){
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
        $hashedRequestPayload = Encoder::hexEncode(Encoder::hash(Utils::toBytes(""), $signatureAlgorithm));
        if (!Utils::isUnset($request->stream)) {
            $tmp = Utils::readAsBytes($request->stream);
            $hashedRequestPayload = Encoder::hexEncode(Encoder::hash($tmp, $signatureAlgorithm));
            $request->stream = $tmp;
            $request->headers["content-type"] = "application/octet-stream";
        }
        else {
            if (!Utils::isUnset($request->body)) {
                if (Utils::equalString($request->reqBodyType, "json")) {
                    $jsonObj = Utils::toJSONString($request->body);
                    $hashedRequestPayload = Encoder::hexEncode(Encoder::hash(Utils::toBytes($jsonObj), $signatureAlgorithm));
                    $request->stream = $jsonObj;
                    $request->headers["content-type"] = "application/json; charset=utf-8";
                }
                else {
                    $m = Utils::assertAsMap($request->body);
                    $formObj = OpenApiUtilClient::toForm($m);
                    $hashedRequestPayload = Encoder::hexEncode(Encoder::hash(Utils::toBytes($formObj), $signatureAlgorithm));
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
            $request->headers["Authorization"] = $this->getAuthorizationForPop($request->pathname, $request->method, $request->query, $request->headers, $signatureAlgorithm, $hashedRequestPayload, $accessKeyId, $accessKeySecret);
        }
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
    public function getAuthorizationForFc($pathname, $method, $query, $headers, $ak, $secret){
        $sign = $this->getSignatureForFc($pathname, $method, $query, $headers, $secret);
        return "FC " . $ak . ":" . $sign . "";
    }

    /**
     * @param string $pathname
     * @param string $method
     * @param string[] $query
     * @param string[] $headers
     * @param string $secret
     * @return string
     */
    public function getSignatureForFc($pathname, $method, $query, $headers, $secret){
        $resource = $pathname;
        $contentMd5 = @$headers["content-md5"];
        if (Utils::isUnset($contentMd5)) {
            $contentMd5 = "";
        }
        $contentType = @$headers["content-type"];
        if (Utils::isUnset($contentType)) {
            $contentType = "";
        }
        $stringToSign = "";
        $canonicalizedResource = $this->buildCanonicalizedResourceForFc($resource, $query);
        $canonicalizedHeaders = $this->buildCanonicalizedHeadersForFc($headers);
        $stringToSign = "" . $method . "\n" . $contentMd5 . "\n" . $contentType . "\n" . @$headers["date"] . "\n" . $canonicalizedHeaders . "" . $canonicalizedResource . "";
        return Encoder::base64EncodeToString(Signer::HmacSHA256Sign($stringToSign, $secret));
    }

    /**
     * @param string $pathname
     * @param string[] $query
     * @return string
     */
    public function buildCanonicalizedResourceForFc($pathname, $query){
        $paths = StringUtil::split($pathname, "?", 2);
        $canonicalizedResource = @$paths[0];
        $resources = [];
        if (Utils::equalNumber(ArrayUtil::size($paths), 2)) {
            $resources = StringUtil::split(@$paths[1], "&", 0);
        }
        $subResources = [];
        $tmp = "";
        $separator = "";
        if (!Utils::isUnset($query)) {
            $queryList = MapUtil::keySet($query);
            foreach($queryList as $paramName){
                $tmp = "" . $tmp . "" . $separator . "" . $paramName . "";
                if (!Utils::isUnset(@$query[$paramName])) {
                    $tmp = "" . $tmp . "=" . @$query[$paramName] . "";
                }
                $separator = ";";
            }
            $subResources = StringUtil::split($tmp, ";", 0);
        }
        $result = ArrayUtil::concat($subResources, $resources);
        $sortedParams = ArrayUtil::ascSort($result);
        if (Utils::equalNumber(ArrayUtil::size($sortedParams), 0)) {
            return "" . $canonicalizedResource . "\n";
        }
        return "" . $canonicalizedResource . "\n" . ArrayUtil::join($sortedParams, "\n") . "";
    }

    /**
     * @param string[] $headers
     * @return string
     */
    public function buildCanonicalizedHeadersForFc($headers){
        $canonicalizedHeaders = "";
        $keys = MapUtil::keySet($headers);
        $sortedHeaders = ArrayUtil::ascSort($keys);
        foreach($sortedHeaders as $header){
            if (StringUtil::contains(StringUtil::toLower($header), "x-fc-")) {
                $canonicalizedHeaders = "" . $canonicalizedHeaders . "" . StringUtil::toLower($header) . ":" . @$headers[$header] . "\n";
            }
        }
        return $canonicalizedHeaders;
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
    public function getAuthorizationForPop($pathname, $method, $query, $headers, $signatureAlgorithm, $payload, $ak, $secret){
        $signature = $this->getSignatureForPop($pathname, $method, $query, $headers, $signatureAlgorithm, $payload, $secret);
        return "" . $signatureAlgorithm . "  Credential=" . $ak . ",SignedHeaders=" . ArrayUtil::join($this->getSignedHeaders($headers), ";") . ",Signature=" . $signature . "";
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
    public function getSignatureForPop($pathname, $method, $query, $headers, $signatureAlgorithm, $payload, $secret){
        $canonicalURI = "/";
        if (!Utils::empty_($pathname)) {
            $canonicalURI = $pathname;
        }
        $stringToSign = "";
        $canonicalizedResource = $this->buildCanonicalizedResourceForPop($query);
        $canonicalizedHeaders = $this->buildCanonicalizedHeadersForPop($headers);
        $signedHeaders = $this->getSignedHeaders($headers);
        $stringToSign = "" . $method . "\n" . $canonicalURI . "\n" . $canonicalizedResource . "\n" . $canonicalizedHeaders . "\n" . ArrayUtil::join($signedHeaders, ";") . "\n" . $payload . "";
        $hex = Encoder::hexEncode(Encoder::hash(Utils::toBytes($stringToSign), $signatureAlgorithm));
        $stringToSign = "" . $signatureAlgorithm . "\n" . $hex . "";
        $signature = Utils::toBytes("");
        if (StringUtil::equals($signatureAlgorithm, "ACS3-HMAC-SHA256")) {
            $signature = Signer::HmacSHA256Sign($stringToSign, $secret);
        }
        else if (StringUtil::equals($signatureAlgorithm, "ACS3-HMAC-SM3")) {
            $signature = Signer::HmacSM3Sign($stringToSign, $secret);
        }
        else if (StringUtil::equals($signatureAlgorithm, "ACS3-RSA-SHA256")) {
            $signature = Signer::SHA256withRSASign($stringToSign, $secret);
        }
        return Encoder::hexEncode($signature);
    }

    /**
     * @param string[] $query
     * @return string
     */
    public function buildCanonicalizedResourceForPop($query){
        $queryArray = MapUtil::keySet($query);
        $sortedQueryArray = ArrayUtil::ascSort($queryArray);
        $canonicalizedResource = "";
        foreach($sortedQueryArray as $key){
            $canonicalizedResource = "" . $canonicalizedResource . "&" . Encoder::percentEncode($key) . "";
            if (!Utils::empty_(@$query[$key])) {
                $canonicalizedResource = "" . $canonicalizedResource . "=" . Encoder::percentEncode(@$query[$key]) . "";
            }
        }
        return $canonicalizedResource;
    }

    /**
     * @param string[] $headers
     * @return string
     */
    public function buildCanonicalizedHeadersForPop($headers){
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
