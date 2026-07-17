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
use AlibabaCloud\Darabonba\ArrayUtil\ArrayUtil;
use AlibabaCloud\Tea\Exception\TeaError;
use Darabonba\GatewayOss\Util\Client as DarabonbaGatewayOssUtilClient;
use \Exception;

use Darabonba\GatewaySpi\Models\InterceptorContext;
use Darabonba\GatewaySpi\Models\AttributeMap;

class Client extends DarabonbaGatewaySpiClient {
    protected $_default_signed_params;

    protected $_except_signed_params;

    protected $_default_additional_headers;

    public function __construct(){
        parent::__construct();
        $this->_default_signed_params = [
            "response-content-type",
            "response-content-language",
            "response-cache-control",
            "logging",
            "response-content-encoding",
            "acl",
            "uploadId",
            "uploads",
            "partNumber",
            "group",
            "link",
            "delete",
            "website",
            "location",
            "objectInfo",
            "objectMeta",
            "response-expires",
            "response-content-disposition",
            "cors",
            "lifecycle",
            "restore",
            "qos",
            "referer",
            "stat",
            "bucketInfo",
            "append",
            "position",
            "security-token",
            "live",
            "comp",
            "status",
            "vod",
            "startTime",
            "endTime",
            "x-oss-process",
            "symlink",
            "callback",
            "callback-var",
            "tagging",
            "encryption",
            "versions",
            "versioning",
            "versionId",
            "policy",
            "requestPayment",
            "x-oss-traffic-limit",
            "qosInfo",
            "asyncFetch",
            "x-oss-request-payer",
            "sequential",
            "inventory",
            "inventoryId",
            "continuation-token",
            "callback",
            "callback-var",
            "worm",
            "wormId",
            "wormExtend",
            "replication",
            "replicationLocation",
            "replicationProgress",
            "transferAcceleration",
            "cname",
            "metaQuery",
            "x-oss-ac-source-ip",
            "x-oss-ac-subnet-mask",
            "x-oss-ac-vpc-id",
            "x-oss-ac-forward-allow",
            "resourceGroup",
            "style",
            "styleName",
            "x-oss-async-process",
            "rtc",
            "accessPoint",
            "accessPointPolicy",
            "httpsConfig",
            "regionsV2",
            "publicAccessBlock",
            "policyStatus",
            "redundancyTransition",
            "redundancyType",
            "redundancyProgress",
            "dataAccelerator",
            "verbose",
            "accessPointForObjectProcess",
            "accessPointConfigForObjectProcess",
            "accessPointPolicyForObjectProcess",
            "bucketArchiveDirectRead",
            "responseHeader",
            "userDefinedLogFieldsConfig",
            "reservedcapacity",
            "requesterQosInfo",
            "qosRequester",
            "resourcePool",
            "resourcePoolInfo",
            "resourcePoolBuckets",
            "processConfiguration",
            "img",
            "asyncFetch",
            "virtualBucket",
            "copy",
            "userRegion",
            "partSize",
            "chunkSize",
            "partUploadId",
            "chunkNumber",
            "userRegion",
            "regionList",
            "eventnotification",
            "cacheConfiguration",
            "dfs",
            "dfsadmin",
            "dfssecurity"
        ];
        $this->_except_signed_params = [
            "list-type",
            "regions"
        ];
        $this->_default_additional_headers = [
            "range",
            "if-modified-since"
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
        if (Utils::isUnset($regionId) || Utils::empty_($regionId)) {
            $regionId = $this->getRegionIdFromEndpoint($config->endpoint);
        }
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
                // for python:
                // xml_str = OSS_UtilClient.to_xml(req_body_map)
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
        $host = $this->getHost($config->endpointType, $bucketName, $config->endpoint, $context);
        $request->headers = Tea::merge([
            "host" => $host,
            "date" => Utils::getDateUTCString(),
            "user-agent" => $request->userAgent
        ], $request->headers);
        $originPath = $request->pathname;
        $originQuery = $request->query;
        if (!Utils::empty_($originPath)) {
            $pathAndQueries = StringUtil::split($originPath, "?", 2);
            $request->pathname = @$pathAndQueries[0];
            if (Utils::equalNumber(ArrayUtil::size($pathAndQueries), 2)) {
                $pathQueries = StringUtil::split(@$pathAndQueries[1], "&", null);
                foreach($pathQueries as $sub){
                    $item = StringUtil::split($sub, "=", null);
                    $queryKey = @$item[0];
                    $queryValue = "";
                    if (Utils::equalNumber(ArrayUtil::size($item), 2)) {
                        $queryValue = @$item[1];
                    }
                    if (Utils::empty_(@$originQuery[$queryKey])) {
                        $request->query[$queryKey] = $queryValue;
                    }
                }
            }
        }
        $signatureVersion = Utils::defaultString($request->signatureVersion, "v4");
        $request->headers["authorization"] = $this->getAuthorization($signatureVersion, $bucketName, $request->pathname, $request->method, $request->query, $request->headers, $accessKeyId, $accessKeySecret, $regionId);
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
                        "hostId" => @$err["HostId"],
                        "AccessDeniedDetail" => @$err["AccessDeniedDetail"]
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
                    $result = DarabonbaGatewayOssUtilClient::parseXml($bodyStr, $request->action);
                    // for no util language
                    // var result : any = XML.parseXml(bodyStr, null);
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
     * @param string $endpoint
     * @return string
     */
    public function getRegionIdFromEndpoint($endpoint){
        if (!Utils::empty_($endpoint)) {
            $idx = -1;
            if (StringUtil::hasPrefix($endpoint, "oss-") && StringUtil::hasSuffix($endpoint, ".aliyuncs.com")) {
                $idx = StringUtil::index($endpoint, ".aliyuncs.com");
                return StringUtil::subString($endpoint, 4, $idx);
            }
            if (StringUtil::hasSuffix($endpoint, ".mgw.aliyuncs.com")) {
                $idx = StringUtil::index($endpoint, ".mgw.aliyuncs.com");
                return StringUtil::subString($endpoint, 0, $idx);
            }
            if (StringUtil::hasSuffix($endpoint, ".mgw-internal.aliyuncs.com")) {
                $idx = StringUtil::index($endpoint, ".mgw-internal.aliyuncs.com");
                return StringUtil::subString($endpoint, 0, $idx);
            }
            if (StringUtil::hasSuffix($endpoint, "-internal.oss-data-acc.aliyuncs.com")) {
                $idx = StringUtil::index($endpoint, "-internal.oss-data-acc.aliyuncs.com");
                return StringUtil::subString($endpoint, 0, $idx);
            }
            if (StringUtil::hasSuffix($endpoint, ".oss-dls.aliyuncs.com")) {
                $idx = StringUtil::index($endpoint, ".oss-dls.aliyuncs.com");
                return StringUtil::subString($endpoint, 0, $idx);
            }
        }
        return "cn-hangzhou";
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
                return "" . $regionId . ".oss.aliyuncs.com";
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
     * @param InterceptorContext $context
     * @return string
     */
    public function getHost($endpointType, $bucketName, $endpoint, $context){
        if (StringUtil::contains($endpoint, ".mgw.aliyuncs.com") && !Utils::isUnset(@$context->request->hostMap["userid"])) {
            return "" . @$context->request->hostMap["userid"] . "." . $endpoint . "";
        }
        if (StringUtil::contains($endpoint, ".mgw-internal.aliyuncs.com") && !Utils::isUnset(@$context->request->hostMap["userid"])) {
            return "" . @$context->request->hostMap["userid"] . "." . $endpoint . "";
        }
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
                $additionalHeaderNames = $this->getAdditionalHeaderNamesV2($headers);
                $sign = $this->getSignatureV2($bucketName, $pathname, $method, $query, $headers, $secret, $additionalHeaderNames);
                $additionalHeaders = $this->joinSemicolon($additionalHeaderNames);
                if (Utils::empty_($additionalHeaders)) {
                    return "OSS2 AccessKeyId:" . $ak . ",Signature:" . $sign . "";
                }
                return "OSS2 AccessKeyId:" . $ak . ",AdditionalHeaders:" . $additionalHeaders . ",Signature:" . $sign . "";
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
        $canonicalizedUri = $pathname;
        if (!Utils::empty_($pathname)) {
            if (!Utils::empty_($bucketName)) {
                $canonicalizedUri = "/" . $bucketName . "" . $canonicalizedUri . "";
            }
        }
        else {
            if (!Utils::empty_($bucketName)) {
                $canonicalizedUri = "/" . $bucketName . "/";
            }
            else {
                $canonicalizedUri = "/";
            }
        }
        // for java:
        // String suffix = (!canonicalizedUri.equals("/") && canonicalizedUri.endsWith("/"))? "/" : "";
        // canonicalizedUri = com.aliyun.openapiutil.Client.getEncodePath(canonicalizedUri) + suffix;
        $canonicalizedUri = OpenApiUtilClient::getEncodePath($canonicalizedUri);
        $queryMap = [];
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
        $canonicalizedResource = $pathname;
        $queryKeys = MapUtil::keySet($query);
        $sortedParams = ArrayUtil::ascSort($queryKeys);
        $separator = "?";
        foreach($sortedParams as $paramName){
            if (ArrayUtil::contains($this->_default_signed_params, $paramName) || StringUtil::hasPrefix($paramName, "x-oss-")) {
                $canonicalizedResource = "" . $canonicalizedResource . "" . $separator . "" . $paramName . "";
                if (!Utils::empty_(@$query[$paramName])) {
                    $canonicalizedResource = "" . $canonicalizedResource . "=" . @$query[$paramName] . "";
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
     * @param string $value
     * @return string
     */
    public static function v2UriEncode($value){
        if (Utils::empty_($value)) {
            return '';
        }
        $encoded = EncodeUtil::percentEncode($value);
        $encoded = StringUtil::replace($encoded, "+", "%20", null);
        $encoded = StringUtil::replace($encoded, "%7E", "~", null);
        $encoded = StringUtil::replace($encoded, "%2D", "-", null);
        $encoded = StringUtil::replace($encoded, "%5F", "_", null);
        $encoded = StringUtil::replace($encoded, "%2E", ".", null);
        return $encoded;
    }

    /**
     * @param string[] $headers
     * @param string $name
     * @return string
     */
    public function getHeaderValue($headers, $name){
        if (!Utils::isUnset(@$headers[$name])) {
            return @$headers[$name];
        }
        foreach(MapUtil::keySet($headers) as $header){
            if (StringUtil::equals(StringUtil::toLower($header), $name)) {
                return @$headers[$header];
            }
        }
        return '';
    }

    /**
     * @param string[] $headers
     * @return array
     */
    public function getAdditionalHeaderNamesV2($headers){
        $additionalHeaders = [];
        foreach(MapUtil::keySet($headers) as $header){
            $lowerHeader = StringUtil::toLower($header);
            if (ArrayUtil::contains($this->_default_additional_headers, $lowerHeader)) {
                $additionalHeaders[$lowerHeader] = $lowerHeader;
            }
        }
        return ArrayUtil::ascSort(MapUtil::keySet($additionalHeaders));
    }

    /**
     * @param string[] $items
     * @return string
     */
    public function joinSemicolon($items){
        $result = "";
        $separator = "";
        foreach($items as $item){
            $result = "" . $result . "" . $separator . "" . $item . "";
            $separator = ";";
        }
        return $result;
    }

    /**
     * @param string[] $headers
     * @param string[] $additionalHeaderNames
     * @return string
     */
    public function buildCanonicalizedOssHeadersV2($headers, $additionalHeaderNames){
        $canonHeaders = [];
        foreach(MapUtil::keySet($headers) as $header){
            $lowerHeader = StringUtil::toLower($header);
            if (StringUtil::hasPrefix($lowerHeader, "x-oss-")) {
                $canonHeaders[$lowerHeader] = @$headers[$header];
            }
        }
        foreach($additionalHeaderNames as $name){
            $canonHeaders[$name] = $this->getHeaderValue($headers, $name);
        }
        $canonicalizedHeaders = "";
        foreach(ArrayUtil::ascSort(MapUtil::keySet($canonHeaders)) as $header){
            $canonicalizedHeaders = "" . $canonicalizedHeaders . "" . $header . ":" . @$canonHeaders[$header] . "\n";
        }
        return $canonicalizedHeaders;
    }

    /**
     * @param string[] $query
     * @return string
     */
    public function buildCanonicalizedQueryStringV2($query){
        $queryMap = [];
        if (!Utils::isUnset($query)) {
            foreach(MapUtil::keySet($query) as $queryKey){
                $encodedKey = self::v2UriEncode($queryKey);
                $encodedValue = "";
                if (!Utils::empty_(@$query[$queryKey])) {
                    $encodedValue = self::v2UriEncode(@$query[$queryKey]);
                }
                $queryMap[$encodedKey] = $encodedValue;
            }
        }
        if (Utils::isUnset($queryMap) || Utils::equalNumber(ArrayUtil::size(MapUtil::keySet($queryMap)), 0)) {
            return '';
        }
        $canonicalizedQueryString = "?";
        $separator = "";
        foreach(ArrayUtil::ascSort(MapUtil::keySet($queryMap)) as $key){
            $canonicalizedQueryString = "" . $canonicalizedQueryString . "" . $separator . "" . $key . "";
            if (!Utils::empty_(@$queryMap[$key])) {
                $canonicalizedQueryString = "" . $canonicalizedQueryString . "=" . @$queryMap[$key] . "";
            }
            $separator = "&";
        }
        return $canonicalizedQueryString;
    }

    /**
     * @param string $bucketName
     * @param string $pathname
     * @param string[] $query
     * @return string
     */
    public function buildCanonicalizedResourceV2($bucketName, $pathname, $query){
        $resourcePath = "/";
        if (!Utils::empty_($bucketName)) {
            $resourcePath = "/" . $bucketName . "" . $pathname . "";
        }
        else if (!Utils::empty_($pathname)) {
            $resourcePath = $pathname;
        }
        $canonicalizedResource = self::v2UriEncode($resourcePath);
        $canonicalizedResource = "" . $canonicalizedResource . "" . $this->buildCanonicalizedQueryStringV2($query) . "";
        return $canonicalizedResource;
    }

    /**
     * @param string $bucketName
     * @param string $pathname
     * @param string $method
     * @param string[] $query
     * @param string[] $headers
     * @param string $secret
     * @param string[] $additionalHeaderNames
     * @return string
     */
    public function getSignatureV2($bucketName, $pathname, $method, $query, $headers, $secret, $additionalHeaderNames){
        $contentMd5 = Utils::defaultString(@$headers["content-md5"], "");
        $contentType = Utils::defaultString(@$headers["content-type"], "");
        $date = Utils::defaultString(@$headers["date"], "");
        $canonicalizedOssHeaders = $this->buildCanonicalizedOssHeadersV2($headers, $additionalHeaderNames);
        $additionalHeaders = $this->joinSemicolon($additionalHeaderNames);
        $canonicalizedResource = $this->buildCanonicalizedResourceV2($bucketName, $pathname, $query);
        $stringToSign = "" . $method . "\n" . $contentMd5 . "\n" . $contentType . "\n" . $date . "\n" . $canonicalizedOssHeaders . "" . $additionalHeaders . "\n" . $canonicalizedResource . "";
        return EncodeUtil::base64EncodeToString(SignatureUtil::HmacSHA256Sign($stringToSign, $secret));
    }
}
