<?php

// This file is auto-generated, don't edit it. Thanks.
namespace Darabonba\GatewayOss;

use Darabonba\GatewaySpi\Client as DarabonbaGatewaySpiClient;
use AlibabaCloud\Tea\Utils\Utils;
use AlibabaCloud\Darabonba\MapUtil\MapUtil;
use AlibabaCloud\Darabonba\String\StringUtil;
use AlibabaCloud\Tea\XML\XML;
use AlibabaCloud\Darabonba\SignatureUtil\SignatureUtil;
use AlibabaCloud\Darabonba\EncodeUtil\EncodeUtil;
use AlibabaCloud\OpenApiUtil\OpenApiUtilClient;
use AlibabaCloud\Tea\OSSUtils\OSSUtils;
use AlibabaCloud\Tea\Tea;
use AlibabaCloud\Tea\Exception\TeaError;
use AlibabaCloud\Darabonba\ArrayUtil\ArrayUtil;
use \Exception;

use Darabonba\GatewaySpi\Models\InterceptorContext;
use Darabonba\GatewaySpi\Models\AttributeMap;

class Client extends DarabonbaGatewaySpiClient {
    protected $_default_signed_params;

    protected $_except_signed_params;

    public function __construct(){
        parent::__construct();
        $this->_default_signed_params = [
            "location",
            "cors",
            "objectMeta",
            "uploadId",
            "partNumber",
            "security-token",
            "position",
            "img",
            "style",
            "styleName",
            "replication",
            "replicationProgress",
            "replicationLocation",
            "cname",
            "qos",
            "startTime",
            "endTime",
            "symlink",
            "x-oss-process",
            "response-content-type",
            "response-content-language",
            "response-expires",
            "response-cache-control",
            "response-content-disposition",
            "response-content-encoding",
            "udf",
            "udfName",
            "udfImage",
            "udfId",
            "udfImageDesc",
            "udfApplication",
            "udfApplicationLog",
            "restore",
            "callback",
            "callback-var",
            "policy",
            "encryption",
            "versions",
            "versioning",
            "versionId",
            "wormId"
        ];
        $this->_except_signed_params = [
            "list-type",
            "regions"
        ];
    }

    /**
     * @param InterceptorContext $context
     * @param AttributeMap $attributeMap
     * @return void
     */
    public function modifyConfiguration($context, $attributeMap){
        $config = $context->configuration;
        $config->endpoint = $this->getEndpoint($config->regionId, $config->network, $config->endpoint);
    }

    /**
     * @param InterceptorContext $context
     * @param AttributeMap $attributeMap
     * @return void
     */
    public function modifyRequest($context, $attributeMap){
        $request = $context->request;
        $hostMap = [];
        if (!Utils::isUnset($request->hostMap)) {
            $hostMap = $request->hostMap;
        }
        $bucketName = @$hostMap["bucket"];
        if (Utils::isUnset($bucketName)) {
            $bucketName = "";
        }
        if (!Utils::isUnset(@$request->headers["x-oss-meta-*"])) {
            $tmp = Utils::parseJSON(@$request->headers["x-oss-meta-*"]);
            $mapData = Utils::assertAsMap($tmp);
            $metaData = Utils::stringifyMapValue($mapData);
            $metaKeySet = MapUtil::keySet($metaData);
            $request->headers["x-oss-meta-*"] = null;
            foreach($metaKeySet as $key){
                $newKey = "x-oss-meta-" . $key . "";
                $request->headers[$newKey] = @$metaData[$key];
            }
        }
        $config = $context->configuration;
        $regionId = $config->regionId;
        $credential = $request->credential;
        $accessKeyId = $credential->getAccessKeyId();
        $accessKeySecret = $credential->getAccessKeySecret();
        $securityToken = $credential->getSecurityToken();
        if (!Utils::empty_($securityToken)) {
            $request->headers["x-oss-security-token"] = $securityToken;
        }
        if (!Utils::isUnset($request->body)) {
            if (StringUtil::equals($request->reqBodyType, "xml")) {
                $reqBodyMap = Utils::assertAsMap($request->body);
                $xmlStr = XML::toXML($reqBodyMap);
                $request->stream = $xmlStr;
                $request->headers["content-type"] = "application/xml";
                $request->headers["content-md5"] = EncodeUtil::base64EncodeToString(SignatureUtil::MD5Sign($xmlStr));
            }
            else if (StringUtil::equals($request->reqBodyType, "json")) {
                $reqBodyStr = Utils::toJSONString($request->body);
                $request->stream = $reqBodyStr;
                $request->headers["content-type"] = "application/json; charset=utf-8";
            }
            else if (StringUtil::equals($request->reqBodyType, "formData")) {
                $reqBodyForm = Utils::assertAsMap($request->body);
                $request->stream = OpenApiUtilClient::toForm($reqBodyForm);
                $request->headers["content-type"] = "application/x-www-form-urlencoded";
            }
            else if (StringUtil::equals($request->reqBodyType, "binary")) {
                $attributeMap->key = [
                    "crc" => "",
                    "md5" => ""
                ];
                $request->stream = OSSUtils::inject($request->stream, $attributeMap->key);
                $request->headers["content-type"] = "application/octet-stream";
            }
        }
        $host = $this->getHost($config->endpointType, $bucketName, $config->endpoint);
        $request->headers = Tea::merge([
            "host" => $host,
            "date" => Utils::getDateUTCString(),
            "user-agent" => $request->userAgent
        ], $request->headers);
        $request->headers["authorization"] = $this->getAuthorization($request->signatureVersion, $bucketName, $request->pathname, $request->method, $request->query, $request->headers, $accessKeyId, $accessKeySecret, $regionId);
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
        $bodyStr = null;
        if (Utils::is4xx($response->statusCode) || Utils::is5xx($response->statusCode)) {
            $bodyStr = Utils::readAsString($response->body);
            if (!Utils::empty_($bodyStr)) {
                $respMap = XML::parseXml($bodyStr, null);
                $err = Utils::assertAsMap(@$respMap["Error"]);
                throw new TeaError([
                    "code" => @$err["Code"],
                    "message" => @$err["Message"],
                    "data" => [
                        "statusCode" => $response->statusCode,
                        "requestId" => @$err["RequestId"],
                        "ecCode" => @$err["EC"],
                        "Recommend" => @$err["RecommendDoc"],
                        "hostId" => @$err["HostId"]
                    ]
                ]);
            }
            else {
                $headers = $response->headers;
                $requestId = @$headers["x-oss-request-id"];
                $ecCode = @$headers["x-oss-ec-code"];
                throw new TeaError([
                    "code" => $response->statusCode,
                    "message" => null,
                    "data" => [
                        "statusCode" => $response->statusCode,
                        "requestId" => "" . $requestId . "",
                        "ecCode" => $ecCode
                    ]
                ]);
            }
        }
        $ctx = $attributeMap->key;
        if (!Utils::isUnset($ctx)) {
            if (!Utils::isUnset(@$ctx["crc"]) && !Utils::isUnset(@$response->headers["x-oss-hash-crc64ecma"]) && !StringUtil::equals(@$ctx["crc"], @$response->headers["x-oss-hash-crc64ecma"])) {
                throw new TeaError([
                    "code" => "CrcNotMatched",
                    "data" => [
                        "clientCrc" => @$ctx["crc"],
                        "serverCrc" => @$response->headers["x-oss-hash-crc64ecma"]
                    ]
                ]);
            }
            if (!Utils::isUnset(@$ctx["md5"]) && !Utils::isUnset(@$response->headers["content-md5"]) && !StringUtil::equals(@$ctx["md5"], @$response->headers["content-md5"])) {
                throw new TeaError([
                    "code" => "MD5NotMatched",
                    "data" => [
                        "clientMD5" => @$ctx["md5"],
                        "serverMD5" => @$response->headers["content-md5"]
                    ]
                ]);
            }
        }
        if (!Utils::isUnset($response->body)) {
            if (Utils::equalNumber($response->statusCode, 204)) {
                Utils::readAsString($response->body);
            }
            else if (StringUtil::equals($request->bodyType, "xml")) {
                $bodyStr = Utils::readAsString($response->body);
                $response->deserializedBody = $bodyStr;
                if (!Utils::empty_($bodyStr)) {
                    $respStruct = self::getResponseBodySchema($request->action);
                    $result = XML::parseXml($bodyStr, $respStruct);
                    try {
                        $response->deserializedBody = Utils::assertAsMap($result);
                    }
                    catch (Exception $error) {
                        if (!($error instanceof TeaError)) {
                            $error = new TeaError([], $error->getMessage(), $error->getCode(), $error);
                        }
                        $response->deserializedBody = $result;
                    }
                }
            }
            else if (Utils::equalString($request->bodyType, "binary")) {
                $response->deserializedBody = $response->body;
            }
            else if (Utils::equalString($request->bodyType, "byte")) {
                $byt = Utils::readAsBytes($response->body);
                $response->deserializedBody = $byt;
            }
            else if (Utils::equalString($request->bodyType, "string")) {
                $response->deserializedBody = Utils::readAsString($response->body);
            }
            else if (Utils::equalString($request->bodyType, "json")) {
                $obj = Utils::readAsJSON($response->body);
                $res = Utils::assertAsMap($obj);
                $response->deserializedBody = $res;
            }
            else if (Utils::equalString($request->bodyType, "array")) {
                $response->deserializedBody = Utils::readAsJSON($response->body);
            }
            else {
                $response->deserializedBody = Utils::readAsString($response->body);
            }
        }
    }

    /**
     * @param string $regionId
     * @param string $network
     * @param string $endpoint
     * @return string
     */
    public function getEndpoint($regionId, $network, $endpoint){
        if (!Utils::empty_($endpoint)) {
            return $endpoint;
        }
        if (Utils::empty_($regionId)) {
            $regionId = "cn-hangzhou";
        }
        if (!Utils::empty_($network)) {
            if (StringUtil::contains($network, "internal")) {
                return "oss-" . $regionId . "-internal.aliyuncs.com";
            }
            else if (StringUtil::contains($network, "ipv6")) {
                return "" . $regionId . "oss.aliyuncs.com";
            }
            else if (StringUtil::contains($network, "accelerate")) {
                return "oss-" . $network . ".aliyuncs.com";
            }
        }
        return "oss-" . $regionId . ".aliyuncs.com";
    }

    /**
     * @param string $endpointType
     * @param string $bucketName
     * @param string $endpoint
     * @return string
     */
    public function getHost($endpointType, $bucketName, $endpoint){
        if (Utils::empty_($bucketName)) {
            return $endpoint;
        }
        $host = "" . $bucketName . "." . $endpoint . "";
        if (!Utils::empty_($endpointType)) {
            if (StringUtil::equals($endpointType, "ip")) {
                $host = "" . $endpoint . "/" . $bucketName . "";
            }
            else if (StringUtil::equals($endpointType, "cname")) {
                $host = $endpoint;
            }
        }
        return $host;
    }

    /**
     * @param string $signatureVersion
     * @param string $bucketName
     * @param string $pathname
     * @param string $method
     * @param string[] $query
     * @param string[] $headers
     * @param string $ak
     * @param string $secret
     * @param string $regionId
     * @return string
     */
    public function getAuthorization($signatureVersion, $bucketName, $pathname, $method, $query, $headers, $ak, $secret, $regionId){
        $sign = "";
        if (!Utils::isUnset($signatureVersion)) {
            if (StringUtil::equals($signatureVersion, "v1")) {
                $sign = $this->getSignatureV1($bucketName, $pathname, $method, $query, $headers, $secret);
                return "OSS " . $ak . ":" . $sign . "";
            }
            if (StringUtil::equals($signatureVersion, "v2")) {
                $sign = $this->getSignatureV2($bucketName, $pathname, $method, $query, $headers, $secret);
                return "OSS2 AccessKeyId:" . $ak . ",Signature:" . $sign . "";
            }
        }
        $dateTime = OpenApiUtilClient::getTimestamp();
        $dateTime = StringUtil::replace($dateTime, "-", "", null);
        $dateTime = StringUtil::replace($dateTime, ":", "", null);
        @$headers["x-oss-date"] = $dateTime;
        @$headers["x-oss-content-sha256"] = "UNSIGNED-PAYLOAD";
        $onlyDate = StringUtil::subString($dateTime, 0, 8);
        $cred = "" . $ak . "/" . $onlyDate . "/" . $regionId . "/oss/aliyun_v4_request";
        $sign = $this->getSignatureV4($bucketName, $pathname, $method, $query, $headers, $onlyDate, $regionId, $secret);
        return "OSS4-HMAC-SHA256 Credential=" . $cred . ", Signature=" . $sign . "";
    }

    /**
     * @param string $secret
     * @param string $onlyDate
     * @param string $regionId
     * @return array
     */
    public function getSignKey($secret, $onlyDate, $regionId){
        $temp = "aliyun_v4" . $secret . "";
        $res = SignatureUtil::HmacSHA256Sign($onlyDate, $temp);
        $res = SignatureUtil::HmacSHA256SignByBytes($regionId, $res);
        $res = SignatureUtil::HmacSHA256SignByBytes("oss", $res);
        $res = SignatureUtil::HmacSHA256SignByBytes("aliyun_v4_request", $res);
        return $res;
    }

    /**
     * @param string $bucketName
     * @param string $pathname
     * @param string $method
     * @param string[] $query
     * @param string[] $headers
     * @param string $onlyDate
     * @param string $regionId
     * @param string $secret
     * @return string
     */
    public function getSignatureV4($bucketName, $pathname, $method, $query, $headers, $onlyDate, $regionId, $secret){
        $signingkey = $this->getSignKey($secret, $onlyDate, $regionId);
        $objectName = "/";
        $queryMap = [];
        if (!Utils::empty_($pathname)) {
            $paths = StringUtil::split($pathname, "?", 2);
            $objectName = @$paths[0];
            if (Utils::equalNumber(ArrayUtil::size($paths), 2)) {
                $subResources = StringUtil::split(@$paths[1], "&", null);
                foreach($subResources as $sub){
                    $item = StringUtil::split($sub, "=", null);
                    $key = @$item[0];
                    $key = EncodeUtil::percentEncode($key);
                    $key = StringUtil::replace($key, "+", "%20", null);
                    $value = null;
                    if (Utils::equalNumber(ArrayUtil::size($item), 2)) {
                        $value = EncodeUtil::percentEncode(@$item[1]);
                        $value = StringUtil::replace($value, "+", "%20", null);
                    }
                    // for go : queryMap[tea.StringValue(key)] = value
                    $queryMap[$key] = $value;
                }
            }
        }
        $canonicalizedUri = "/";
        if (!Utils::empty_($bucketName)) {
            $canonicalizedUri = "/" . $bucketName . "" . $objectName . "";
        }
        // for java:
        // String suffix = (!canonicalizedUri.equals("/") && canonicalizedUri.endsWith("/"))? "/" : "";
        // canonicalizedUri = com.aliyun.openapiutil.Client.getEncodePath(canonicalizedUri) + suffix;
        $canonicalizedUri = OpenApiUtilClient::getEncodePath($canonicalizedUri);
        foreach(MapUtil::keySet($query) as $queryKey){
            $queryValue = null;
            if (!Utils::empty_(@$query[$queryKey])) {
                $queryValue = EncodeUtil::percentEncode(@$query[$queryKey]);
                $queryValue = StringUtil::replace($queryValue, "+", "%20", null);
            }
            $queryKey = EncodeUtil::percentEncode($queryKey);
            $queryKey = StringUtil::replace($queryKey, "+", "%20", null);
            // for go : queryMap[tea.StringValue(queryKey)] = queryValue
            $queryMap[$queryKey] = $queryValue;
        }
        $canonicalizedQueryString = $this->buildCanonicalizedQueryStringV4($queryMap);
        $canonicalizedHeaders = $this->buildCanonicalizedHeadersV4($headers);
        $payload = "UNSIGNED-PAYLOAD";
        $canonicalRequest = "" . $method . "\n" . $canonicalizedUri . "\n" . $canonicalizedQueryString . "\n" . $canonicalizedHeaders . "\n\n" . $payload . "";
        $hex = EncodeUtil::hexEncode(EncodeUtil::hash(Utils::toBytes($canonicalRequest), "ACS4-HMAC-SHA256"));
        $scope = "" . $onlyDate . "/" . $regionId . "/oss/aliyun_v4_request";
        $stringToSign = "OSS4-HMAC-SHA256\n" . @$headers["x-oss-date"] . "\n" . $scope . "\n" . $hex . "";
        $signature = SignatureUtil::HmacSHA256SignByBytes($stringToSign, $signingkey);
        return EncodeUtil::hexEncode($signature);
    }

    /**
     * @param string[] $queryMap
     * @return string
     */
    public function buildCanonicalizedQueryStringV4($queryMap){
        $canonicalizedQueryString = "";
        if (!Utils::isUnset($queryMap)) {
            $queryArray = MapUtil::keySet($queryMap);
            $sortedQueryArray = ArrayUtil::ascSort($queryArray);
            $separator = "";
            foreach($sortedQueryArray as $key){
                $canonicalizedQueryString = "" . $canonicalizedQueryString . "" . $separator . "" . $key . "";
                if (!Utils::empty_(@$queryMap[$key])) {
                    $canonicalizedQueryString = "" . $canonicalizedQueryString . "=" . @$queryMap[$key] . "";
                }
                $separator = "&";
            }
        }
        return $canonicalizedQueryString;
    }

    /**
     * @param string[] $headers
     * @return string
     */
    public function buildCanonicalizedHeadersV4($headers){
        $canonicalizedHeaders = "";
        $headersArray = MapUtil::keySet($headers);
        $sortedHeadersArray = ArrayUtil::ascSort($headersArray);
        foreach($sortedHeadersArray as $key){
            $lowerKey = StringUtil::toLower($key);
            if (StringUtil::hasPrefix($lowerKey, "x-oss-") || StringUtil::equals($lowerKey, "content-type") || StringUtil::equals($lowerKey, "content-md5")) {
                $canonicalizedHeaders = "" . $canonicalizedHeaders . "" . $key . ":" . StringUtil::trim(@$headers[$key]) . "\n";
            }
        }
        return $canonicalizedHeaders;
    }

    /**
     * @param string $bucketName
     * @param string $pathname
     * @param string $method
     * @param string[] $query
     * @param string[] $headers
     * @param string $secret
     * @return string
     */
    public function getSignatureV1($bucketName, $pathname, $method, $query, $headers, $secret){
        $resource = "";
        $stringToSign = "";
        if (!Utils::empty_($bucketName)) {
            $resource = "/" . $bucketName . "";
        }
        $resource = "" . $resource . "" . $pathname . "";
        $canonicalizedResource = $this->buildCanonicalizedResource($resource, $query);
        $canonicalizedHeaders = $this->buildCanonicalizedHeaders($headers);
        $stringToSign = "" . $method . "\n" . $canonicalizedHeaders . "" . $canonicalizedResource . "";
        return EncodeUtil::base64EncodeToString(SignatureUtil::HmacSHA1Sign($stringToSign, $secret));
    }

    /**
     * @param string $pathname
     * @param string[] $query
     * @return string
     */
    public function buildCanonicalizedResource($pathname, $query){
        $subResourcesMap = [];
        $canonicalizedResource = $pathname;
        if (!Utils::empty_($pathname)) {
            $paths = StringUtil::split($pathname, "?", 2);
            $canonicalizedResource = @$paths[0];
            if (Utils::equalNumber(ArrayUtil::size($paths), 2)) {
                $subResources = StringUtil::split(@$paths[1], "&", null);
                foreach($subResources as $sub){
                    $hasExcepts = false;
                    foreach($this->_except_signed_params as $excepts){
                        if (StringUtil::contains($sub, $excepts)) {
                            $hasExcepts = true;
                        }
                    }
                    if (!$hasExcepts) {
                        $item = StringUtil::split($sub, "=", null);
                        $key = @$item[0];
                        $value = null;
                        if (Utils::equalNumber(ArrayUtil::size($item), 2)) {
                            $value = @$item[1];
                        }
                        // for go : subResourcesMap[tea.StringValue(key)] = value
                        $subResourcesMap[$key] = $value;
                    }
                }
            }
        }
        $subResourcesArray = MapUtil::keySet($subResourcesMap);
        $newQueryList = $subResourcesArray;
        if (!Utils::isUnset($query)) {
            $queryList = MapUtil::keySet($query);
            $newQueryList = ArrayUtil::concat($queryList, $subResourcesArray);
        }
        $sortedParams = ArrayUtil::ascSort($newQueryList);
        $separator = "?";
        foreach($sortedParams as $paramName){
            if (ArrayUtil::contains($this->_default_signed_params, $paramName)) {
                $canonicalizedResource = "" . $canonicalizedResource . "" . $separator . "" . $paramName . "";
                if (!Utils::isUnset($query) && !Utils::isUnset(@$query[$paramName])) {
                    $canonicalizedResource = "" . $canonicalizedResource . "=" . @$query[$paramName] . "";
                }
                else if (!Utils::isUnset(@$subResourcesMap[$paramName])) {
                    $canonicalizedResource = "" . $canonicalizedResource . "=" . @$subResourcesMap[$paramName] . "";
                }
            }
            else if (ArrayUtil::contains($subResourcesArray, $paramName)) {
                $canonicalizedResource = "" . $canonicalizedResource . "" . $separator . "" . $paramName . "";
                if (!Utils::isUnset(@$subResourcesMap[$paramName])) {
                    $canonicalizedResource = "" . $canonicalizedResource . "=" . @$subResourcesMap[$paramName] . "";
                }
            }
            $separator = "&";
        }
        return $canonicalizedResource;
    }

    /**
     * @param string[] $headers
     * @return string
     */
    public function buildCanonicalizedHeaders($headers){
        $canonicalizedHeaders = "";
        $contentType = @$headers["content-type"];
        if (Utils::isUnset($contentType)) {
            $contentType = "";
        }
        $contentMd5 = @$headers["content-md5"];
        if (Utils::isUnset($contentMd5)) {
            $contentMd5 = "";
        }
        $canonicalizedHeaders = "" . $canonicalizedHeaders . "" . $contentMd5 . "\n" . $contentType . "\n" . @$headers["date"] . "\n";
        $keys = MapUtil::keySet($headers);
        $sortedHeaders = ArrayUtil::ascSort($keys);
        foreach($sortedHeaders as $header){
            if (StringUtil::contains(StringUtil::toLower($header), "x-oss-") && !Utils::isUnset(@$headers[$header])) {
                $canonicalizedHeaders = "" . $canonicalizedHeaders . "" . $header . ":" . @$headers[$header] . "\n";
            }
        }
        return $canonicalizedHeaders;
    }

    /**
     * @param string $bucketName
     * @param string $pathname
     * @param string $method
     * @param string[] $query
     * @param string[] $headers
     * @param string $secret
     * @return string
     */
    public function getSignatureV2($bucketName, $pathname, $method, $query, $headers, $secret){
        return '';
    }

    /**
     * @param string $apiName
     * @return null
     */
    public static function getResponseBodySchema($apiName){
        return null;
    }
}
