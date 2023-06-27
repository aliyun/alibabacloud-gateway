<?php

// This file is auto-generated, don't edit it. Thanks.
namespace Darabonba\GatewayOss;

use Darabonba\GatewaySpi\Client as DarabonbaGatewaySpiClient;
use AlibabaCloud\Tea\Utils\Utils;
use AlibabaCloud\Darabonba\MapUtil\MapUtil;
use AlibabaCloud\Darabonba\String\StringUtil;
use AlibabaCloud\Tea\XML\XML;
use AlibabaCloud\OpenApiUtil\OpenApiUtilClient;
use AlibabaCloud\Tea\OSSUtils\OSSUtils;
use AlibabaCloud\Tea\Tea;
use AlibabaCloud\Tea\Exception\TeaError;
use AlibabaCloud\Darabonba\ArrayUtil\ArrayUtil;
use \Exception;
use AlibabaCloud\Darabonba\SignatureUtil\SignatureUtil;
use AlibabaCloud\Darabonba\EncodeUtil\EncodeUtil;

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
                $request->stream = XML::toXML($reqBodyMap);
                $request->headers["content-type"] = "application/xml";
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
        $request->headers["authorization"] = $this->getAuthorization($request->signatureVersion, $bucketName, $request->pathname, $request->method, $request->query, $request->headers, $accessKeyId, $accessKeySecret);
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
                        "hostId" => @$err["HostId"]
                    ]
                ]);
            }
            else {
                $headers = $response->headers;
                $requestId = @$headers["x-oss-request-id"];
                throw new TeaError([
                    "code" => $response->statusCode,
                    "message" => null,
                    "data" => [
                        "statusCode" => $response->statusCode,
                        "requestId" => "" . $requestId . ""
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
                $result = XML::parseXml($bodyStr, null);
                $list = MapUtil::keySet($result);
                if (Utils::equalNumber(ArrayUtil::size($list), 1)) {
                    $tmp = @$list_[0];
                    try {
                        $response->deserializedBody = Utils::assertAsMap(@$result[$tmp]);
                    }
                    catch (Exception $error) {
                        if (!($error instanceof TeaError)) {
                            $error = new TeaError([], $error->getMessage(), $error->getCode(), $error);
                        }
                        $response->deserializedBody = $result;
                    }
                }
                else {
                    $response->deserializedBody = $result;
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
     * @return string
     */
    public function getAuthorization($signatureVersion, $bucketName, $pathname, $method, $query, $headers, $ak, $secret){
        $sign = "";
        if (Utils::isUnset($signatureVersion) || StringUtil::equals($signatureVersion, "v1")) {
            $sign = $this->getSignatureV1($bucketName, $pathname, $method, $query, $headers, $secret);
            return "OSS " . $ak . ":" . $sign . "";
        }
        else {
            $sign = $this->getSignatureV2($bucketName, $pathname, $method, $query, $headers, $secret);
            return "OSS2 AccessKeyId:" . $ak . ",Signature:" . $sign . "";
        }
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
}
