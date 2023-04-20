<?php

// This file is auto-generated, don't edit it. Thanks.
namespace Darabonba\GatewayDingTalk;

use Darabonba\GatewaySpi\Client as DarabonbaGatewaySpiClient;
use AlibabaCloud\Tea\Tea;
use AlibabaCloud\Tea\Utils\Utils;
use AlibabaCloud\Tea\Exception\TeaError;
use AlibabaCloud\Darabonba\SignatureUtil\SignatureUtil;
use AlibabaCloud\Darabonba\EncodeUtil\EncodeUtil;
use AlibabaCloud\Darabonba\MapUtil\MapUtil;
use AlibabaCloud\Darabonba\ArrayUtil\ArrayUtil;
use AlibabaCloud\Darabonba\String\StringUtil;

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
        if (!Utils::isUnset($request->body)) {
            $jsonObj = Utils::toJSONString($request->body);
            $request->stream = $jsonObj;
            $request->headers["content-type"] = "application/json; charset=utf-8";
        }
        if (!Utils::equalString($request->authType, "Anonymous")) {
            $credential = $request->credential;
            $accessKeyId = $credential->getAccessKeyId();
            $accessKeySecret = $credential->getAccessKeySecret();
            $securityToken = $credential->getSecurityToken();
            if (!Utils::empty_($securityToken)) {
                $request->headers["x-acs-accesskey-id"] = $accessKeyId;
                $request->headers["x-acs-security-token"] = $securityToken;
            }
            $request->headers["Authorization"] = $this->getAuthorization($request->pathname, $request->method, $request->query, $request->headers, $accessKeyId, $accessKeySecret);
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
            @$err["statusCode"] = $response->statusCode;
            throw new TeaError([
                "code" => "" . (string) ($this->defaultAny(@$err["Code"], @$err["code"])) . "",
                "message" => "code: " . (string) ($response->statusCode) . ", " . (string) ($this->defaultAny(@$err["Message"], @$err["message"])) . " request id: " . (string) ($this->defaultAny(@$err["requestId"], @$err["requestid"])) . "",
                "data" => $err,
                "description" => "" . (string) ($this->defaultAny(@$err["Description"], @$err["description"])) . "",
                "accessDeniedDetail" => $this->defaultAny(@$err["accessDeniedDetail"], @$err["accessdenieddetail"])
            ]);
        }
        if (Utils::equalNumber($response->statusCode, 204)) {
            Utils::readAsString($response->body);
        }
        else if (Utils::equalString($request->bodyType, "binary")) {
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
