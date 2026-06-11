<?php

// This file is auto-generated, don't edit it. Thanks.
namespace Darabonba\GatewayMns;

use Darabonba\GatewaySpi\Client as DarabonbaGatewaySpiClient;
use AlibabaCloud\Tea\Exception\TeaError;
use AlibabaCloud\Tea\Utils\Utils;
use AlibabaCloud\Darabonba\String\StringUtil;
use AlibabaCloud\Tea\XML\XML;
use AlibabaCloud\Darabonba\SignatureUtil\SignatureUtil;
use AlibabaCloud\Darabonba\EncodeUtil\EncodeUtil;
use AlibabaCloud\Tea\Tea;
use AlibabaCloud\Darabonba\ArrayUtil\ArrayUtil;
use AlibabaCloud\Darabonba\MapUtil\MapUtil;
use AlibabaCloud\OpenApiUtil\OpenApiUtilClient;

use Darabonba\GatewaySpi\Models\InterceptorContext;
use Darabonba\GatewaySpi\Models\AttributeMap;

class Client extends DarabonbaGatewaySpiClient {
    protected $_signPrefix;

    protected $_signSuffix;

    protected $_authPrefix;

    public function __construct(){
        parent::__construct();
        $this->_signPrefix = "aliyun_v4";
        $this->_signSuffix = "aliyun_v4_request";
        $this->_authPrefix = "MNS4-HMAC-SHA256";
    }

    /**
     * @param InterceptorContext $context
     * @param AttributeMap $attributeMap
     * @return void
     */
    public function modifyConfiguration($context, $attributeMap){
        $config = $context->configuration;
        $config->endpoint = $this->getEndpoint($config->regionId, $config->endpoint);
    }

    /**
     * @param InterceptorContext $context
     * @param AttributeMap $attributeMap
     * @return void
     * @throws TeaError
     */
    public function modifyRequest($context, $attributeMap){
        $request = $context->request;
        $config = $context->configuration;
        $signatureVersion = Utils::defaultString($request->signatureVersion, "v2");
        if (!Utils::isUnset($request->body)) {
            if (StringUtil::equals($request->reqBodyType, "xml")) {
                $reqBodyMap = Utils::assertAsMap($request->body);
                $xmlStr = XML::toXML($reqBodyMap);
                $request->stream = $xmlStr;
                $request->headers["content-type"] = "text/xml;charset=UTF-8";
                $request->headers["content-md5"] = EncodeUtil::base64EncodeToString(SignatureUtil::MD5Sign($xmlStr));
            }
            else if (StringUtil::equals($request->reqBodyType, "json") || StringUtil::equals($request->reqBodyType, "formData")) {
                $bodyStr = Utils::toJSONString($request->body);
                $request->stream = $bodyStr;
                $request->headers["content-type"] = "application/json";
                $request->headers["content-md5"] = EncodeUtil::base64EncodeToString(SignatureUtil::MD5Sign($bodyStr));
            }
            else if (StringUtil::equals($request->reqBodyType, "byte") || StringUtil::equals($request->reqBodyType, "binary")) {
                $bodyBytes = Utils::assertAsBytes($request->body);
                $request->stream = $bodyBytes;
                $request->headers["content-md5"] = EncodeUtil::base64EncodeToString(SignatureUtil::MD5SignForBytes($bodyBytes));
            }
        }
        $this->buildRequest($context);
        $request->headers = Tea::merge([
            "host" => $config->endpoint,
            "x-mns-version" => $request->version,
            "user-agent" => $request->userAgent,
            "accept" => "application/json"
        ], $request->headers);
        if (!Utils::equalString($request->authType, "Anonymous")) {
            $credential = $request->credential;
            if (Utils::isUnset($credential)) {
                throw new TeaError([
                    "code" => "ParameterMissing",
                    "message" => "'config.credential' can not be unset"
                ]);
            }
            $credentialModel = $credential->getCredential();
            $authType = $credentialModel->type;
            if (Utils::equalString($authType, "bearer")) {
                $bearerToken = $credentialModel->bearerToken;
                $request->headers["x-acs-bearer-token"] = $bearerToken;
                $request->headers["x-acs-signature-type"] = "BEARERTOKEN";
                $request->headers["Authorization"] = "Bearer " . $bearerToken . "";
            }
            else {
                $accessKeyId = $credentialModel->accessKeyId;
                $accessKeySecret = $credentialModel->accessKeySecret;
                $securityToken = $credentialModel->securityToken;
                if (!Utils::empty_($securityToken)) {
                    $request->headers["security-token"] = $securityToken;
                }
                $request->headers["date"] = Utils::getDateUTCString();
                if (StringUtil::equals($signatureVersion, "v4")) {
                    $date = $this->getDateISO8601();
                    $request->headers["authorization"] = $this->getAuthorizationV4($context, $date, $accessKeyId, $accessKeySecret);
                }
                else {
                    $request->headers["authorization"] = $this->getAuthorizationV2($request->pathname, $request->method, $request->headers, $accessKeyId, $accessKeySecret);
                }
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
            $err = [];
            if (!Utils::isUnset(@$response->headers["content-type"]) && StringUtil::contains(@$response->headers["content-type"], "text/xml")) {
                $_str = Utils::readAsString($response->body);
                $respMap = XML::parseXml($_str, null);
                $err = Utils::assertAsMap(@$respMap["Error"]);
            }
            else {
                $_res = Utils::readAsJSON($response->body);
                $err = Utils::assertAsMap($_res);
            }
            $requestId = $this->defaultAny(@$err["RequestId"], @$err["requestId"]);
            if (!Utils::isUnset(@$response->headers["x-mns-request-id"])) {
                $requestId = @$response->headers["x-mns-request-id"];
            }
            @$err["statusCode"] = $response->statusCode;
            throw new TeaError([
                "code" => "" . (string) ($this->defaultAny(@$err["Code"], @$err["code"])) . "",
                "message" => "code: " . (string) ($response->statusCode) . ", " . (string) ($this->defaultAny(@$err["Message"], @$err["message"])) . " request id: " . (string) ($requestId) . "",
                "data" => $err,
                "description" => "" . (string) ($this->defaultAny(@$err["Description"], @$err["description"])) . "",
                "accessDeniedDetail" => $this->defaultAny(@$err["AccessDeniedDetail"], @$err["accessDeniedDetail"])
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
     * @param string $regionId
     * @param string $endpoint
     * @return string
     */
    public function getEndpoint($regionId, $endpoint){
        if (!Utils::empty_($endpoint)) {
            return $endpoint;
        }
        if (Utils::empty_($regionId)) {
            $regionId = "cn-hangzhou";
        }
        return "" . $regionId . ".mns.aliyuncs.com";
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
     * @param string $inputValue
     * @param string $defaultValue
     * @return string
     */
    public function defaultString($inputValue, $defaultValue){
        if (Utils::isUnset($inputValue) || Utils::empty_($inputValue)) {
            return $defaultValue;
        }
        return $inputValue;
    }

    /**
     * @param string $input
     * @return string
     */
    public function base64Encode($input){
        if (Utils::isUnset($input)) {
            return '';
        }
        return EncodeUtil::base64EncodeToString(Utils::toBytes($input));
    }

    /**
     * @param string $input
     * @return string
     */
    public function base64Decode($input){
        if (Utils::isUnset($input)) {
            return '';
        }
        return Utils::toString(EncodeUtil::base64Decode($input));
    }

    /**
     * @param InterceptorContext $context
     * @return void
     */
    public function buildRequest($context){
        $request = $context->request;
        $resource = $request->pathname;
        if (!Utils::empty_($resource)) {
            $paths = StringUtil::split($resource, "?", 2);
            $resource = @$paths[0];
            if (Utils::equalNumber(ArrayUtil::size($paths), 2)) {
                $params = StringUtil::split(@$paths[1], "&", null);
                foreach($params as $sub){
                    $item = StringUtil::split($sub, "=", null);
                    $key = @$item[0];
                    $value = null;
                    if (Utils::equalNumber(ArrayUtil::size($item), 2)) {
                        $value = @$item[1];
                    }
                    $request->query[$key] = $value;
                }
            }
        }
        $request->pathname = $resource;
    }

    /**
     * @param string $pathname
     * @param string $method
     * @param string[] $headers
     * @param string $ak
     * @param string $secret
     * @return string
     */
    public function getAuthorizationV2($pathname, $method, $headers, $ak, $secret){
        $sign = $this->getSignatureV2($pathname, $method, $headers, $secret);
        return "MNS " . $ak . ":" . $sign . "";
    }

    /**
     * @param string $pathname
     * @param string $method
     * @param string[] $headers
     * @param string $secret
     * @return string
     */
    public function getSignatureV2($pathname, $method, $headers, $secret){
        $canonicalizedResource = $this->defaultString($pathname, "/");
        $canonicalizedHeaders = $this->buildCanonicalizedHeadersV2($headers);
        $stringToSign = "" . $method . "\n" . $canonicalizedHeaders . "" . $canonicalizedResource . "";
        return EncodeUtil::base64EncodeToString(SignatureUtil::HmacSHA1Sign($stringToSign, $secret));
    }

    /**
     * @param string[] $headers
     * @return string
     */
    public function buildCanonicalizedHeadersV2($headers){
        $contentMd5 = $this->defaultString(@$headers["content-md5"], "");
        $contentType = $this->defaultString(@$headers["content-type"], "");
        $date = $this->defaultString(@$headers["date"], "");
        $canonicalizedHeaders = "" . $contentMd5 . "\n" . $contentType . "\n" . $date . "\n";
        $mnsHeaders = [];
        foreach(MapUtil::keySet($headers) as $header){
            $lowerHeader = StringUtil::toLower($header);
            if (StringUtil::hasPrefix($lowerHeader, "x-mns-")) {
                $mnsHeaders[$lowerHeader] = @$headers[$header];
            }
        }
        foreach(ArrayUtil::ascSort(MapUtil::keySet($mnsHeaders)) as $header){
            $canonicalizedHeaders = "" . $canonicalizedHeaders . "" . $header . ":" . @$mnsHeaders[$header] . "\n";
        }
        return $canonicalizedHeaders;
    }

    /**
     * @param InterceptorContext $context
     * @param string $date
     * @param string $accessKeyId
     * @param string $accessKeySecret
     * @return string
     */
    public function getAuthorizationV4($context, $date, $accessKeyId, $accessKeySecret){
        $region = $this->getRegion($context);
        $dateShort = StringUtil::subString($date, 0, 8);
        $dateShort = StringUtil::replace($dateShort, "T", "", null);
        $scope = "" . $dateShort . "/" . $region . "/mns/" . $this->_signSuffix . "";
        $signingkey = $this->getSigningkeyV4($accessKeySecret, $region, $dateShort);
        $signature = $this->getSignatureV4($context, $date, $scope, $signingkey);
        return "" . $this->_authPrefix . " Credential=" . $accessKeyId . "/" . $scope . ",Signature=" . $signature . "";
    }

    /**
     * @param string $secret
     * @param string $region
     * @param string $date
     * @return array
     */
    public function getSigningkeyV4($secret, $region, $date){
        $sc1 = "" . $this->_signPrefix . "" . $secret . "";
        $sc2 = SignatureUtil::HmacSHA256Sign($date, $sc1);
        $sc3 = SignatureUtil::HmacSHA256SignByBytes($region, $sc2);
        $sc4 = SignatureUtil::HmacSHA256SignByBytes("mns", $sc3);
        return SignatureUtil::HmacSHA256SignByBytes($this->_signSuffix, $sc4);
    }

    /**
     * @param InterceptorContext $context
     * @param string $date
     * @param string $scope
     * @param int[] $signingkey
     * @return string
     */
    public function getSignatureV4($context, $date, $scope, $signingkey){
        $request = $context->request;
        $canonicalString = $this->buildCanonicalRequestV4($request->pathname, $request->method, $request->query, $request->headers);
        $stringToSign = "" . $this->_authPrefix . "\n" . $date . "\n" . $scope . "\n" . $canonicalString . "";
        $signature = SignatureUtil::HmacSHA256SignByBytes($stringToSign, $signingkey);
        return EncodeUtil::hexEncode($signature);
    }

    /**
     * @param string $pathname
     * @param string $method
     * @param string[] $query
     * @param string[] $headers
     * @return string
     */
    public function buildCanonicalRequestV4($pathname, $method, $query, $headers){
        $canonicalURI = "/";
        if (!Utils::empty_($pathname)) {
            $canonicalURI = $pathname;
            if (!StringUtil::hasPrefix($canonicalURI, "/")) {
                $canonicalURI = "/" . $canonicalURI . "";
            }
        }
        $suffix = "";
        if (!StringUtil::equals($canonicalURI, "/") && StringUtil::hasSuffix($canonicalURI, "/")) {
            $suffix = "/";
        }
        $canonicalURI = "" . OpenApiUtilClient::getEncodePath($canonicalURI) . "" . $suffix . "";
        $resources = $this->buildCanonicalizedResourceV4($query);
        $canonicalHeaders = $this->buildCanonicalizedHeadersV4($headers);
        return "" . $method . "\n" . $canonicalURI . "\n" . $resources . "\n" . $canonicalHeaders . "\n";
    }

    /**
     * @param string[] $query
     * @return string
     */
    public function buildCanonicalizedResourceV4($query){
        $canonicalizedResource = "";
        if (!Utils::isUnset($query)) {
            $queryMap = [];
            foreach(MapUtil::keySet($query) as $key){
                $encodedKey = $this->percentEncodeMns(StringUtil::toLower(StringUtil::trim($key)));
                $encodedValue = "";
                if (!Utils::isUnset(@$query[$key]) && !Utils::empty_(@$query[$key])) {
                    $encodedValue = $this->percentEncodeMns(StringUtil::trim(@$query[$key]));
                }
                $queryMap[$encodedKey] = $encodedValue;
            }
            $queryArray = MapUtil::keySet($queryMap);
            $sortedQueryArray = ArrayUtil::ascSort($queryArray);
            $separator = "";
            foreach($sortedQueryArray as $encodedKey){
                $canonicalizedResource = "" . $canonicalizedResource . "" . $separator . "" . $encodedKey . "";
                $encodedValue = @$queryMap[$encodedKey];
                if (!Utils::empty_($encodedValue)) {
                    $canonicalizedResource = "" . $canonicalizedResource . "=" . $encodedValue . "";
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
    public function buildCanonicalizedHeadersV4($headers){
        $signedHeaders = [];
        foreach(MapUtil::keySet($headers) as $key){
            $lowerKey = StringUtil::toLower($key);
            if ($this->hasSignedHeaderV4($lowerKey)) {
                $signedHeaders[$lowerKey] = StringUtil::trim(@$headers[$key]);
            }
        }
        $canonicalizedHeaders = "";
        foreach(ArrayUtil::ascSort(MapUtil::keySet($signedHeaders)) as $lowerKey){
            $canonicalizedHeaders = "" . $canonicalizedHeaders . "" . $lowerKey . ":" . @$signedHeaders[$lowerKey] . "\n";
        }
        return $canonicalizedHeaders;
    }

    /**
     * @param string $header
     * @return bool
     */
    public function hasSignedHeaderV4($header){
        if (StringUtil::equals($header, "content-type") || StringUtil::equals($header, "content-md5")) {
            return true;
        }
        return StringUtil::hasPrefix($header, "x-mns-");
    }

    /**
     * @param string $value
     * @return string
     */
    public function percentEncodeMns($value){
        $encoded = EncodeUtil::percentEncode($value);
        return StringUtil::replace($encoded, "+", "%20", null);
    }

    /**
     * @param InterceptorContext $context
     * @return string
     * @throws TeaError
     */
    public function getRegion($context){
        $config = $context->configuration;
        if (!Utils::isUnset($config->regionId) && !Utils::empty_($config->regionId)) {
            return $config->regionId;
        }
        $region = StringUtil::replace($config->endpoint, ".mns.aliyuncs.com", "", null);
        if (StringUtil::equals($region, $config->endpoint)) {
            throw new TeaError([
                "code" => "ClientConfigError",
                "message" => "The regionId configuration of mns client is missing."
            ]);
        }
        return $region;
    }

    /**
     * @return string
     */
    public function getDateISO8601(){
        $date = OpenApiUtilClient::getTimestamp();
        $date = StringUtil::replace($date, "-", "", null);
        return StringUtil::replace($date, ":", "", null);
    }
}
