<?php

// This file is auto-generated, don't edit it. Thanks.
namespace Darabonba\GatewaySls;

use Darabonba\GatewaySpi\Client as DarabonbaGatewaySpiClient;
use AlibabaCloud\Tea\Utils\Utils;
use AlibabaCloud\Darabonba\String\StringUtil;
use AlibabaCloud\Tea\Tea;
use AlibabaCloud\Darabonba\EncodeUtil\EncodeUtil;
use AlibabaCloud\Darabonba\SignatureUtil\SignatureUtil;
use AlibabaCloud\Tea\Exception\TeaError;
use Darabonba\GatewaySls\Util\Client as DarabonbaGatewaySlsUtilClient;
use AlibabaCloud\Darabonba\MapUtil\MapUtil;
use AlibabaCloud\Darabonba\ArrayUtil\ArrayUtil;
use AlibabaCloud\OpenApiUtil\OpenApiUtilClient;

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
        $project = @$hostMap["project"];
        $config = $context->configuration;
        $credential = $request->credential;
        $credentialModel = $credential->getCredential();
        $accessKeyId = $credentialModel->accessKeyId;
        $accessKeySecret = $credentialModel->accessKeySecret;
        $securityToken = $credentialModel->securityToken;
        if (!Utils::empty_($securityToken)) {
            $request->headers["x-acs-security-token"] = $securityToken;
        }
        $signatureVersion = Utils::defaultString($request->signatureVersion, "v1");
        $contentHash = "";
        if (!Utils::isUnset($request->body)) {
            if (StringUtil::equals($request->reqBodyType, "json")) {
                $bodyStr = Utils::toJSONString($request->body);
                $contentHash = $this->MakeContentHash(Utils::toBytes($bodyStr), $signatureVersion);
                $request->stream = $bodyStr;
                $request->headers["content-type"] = "application/json";
            }
            else if (StringUtil::equals($request->reqBodyType, "formData")) {
                $str = Utils::toJSONString($request->body);
                $contentHash = $this->MakeContentHash(Utils::toBytes($str), $signatureVersion);
                $request->stream = $str;
                $request->headers["content-type"] = "application/json";
            }
            else if (StringUtil::equals($request->reqBodyType, "binary")) {
                // content-type: application/octet-stream
                $bodyBytes = Utils::assertAsBytes($request->body);
                $contentHash = $this->MakeContentHash($bodyBytes, $signatureVersion);
                $request->stream = $bodyBytes;
            }
        }
        $host = $this->getHost($config->network, $project, $config->endpoint);
        $request->headers = Tea::merge([
            "accept" => "application/json",
            "host" => $host,
            "user-agent" => $request->userAgent,
            "x-log-apiversion" => "0.6.0",
            "x-log-bodyrawsize" => "0"
        ], $request->headers);
        $this->buildRequest($context);
        // move param in path to query
        if (StringUtil::equals($signatureVersion, "v4")) {
            if (Utils::empty_($contentHash)) {
                $contentHash = "e3b0c44298fc1c149afbf4c8996fb9242a7e41e4649b934ca495991b7852b855";
            }
            $date = $this->getDateISO8601();
            $request->headers["x-log-date"] = $date;
            $request->headers["x-log-content-sha256"] = $contentHash;
            $request->headers["authorization"] = $this->getAuthorizationV4($context, $date, $contentHash, $accessKeyId, $accessKeySecret);
            return null;
        }
        if (!Utils::empty_($accessKeyId)) {
            $request->headers["x-log-signaturemethod"] = "hmac-sha256";
        }
        $request->headers["date"] = Utils::getDateUTCString();
        $request->headers["content-md5"] = $contentHash;
        $request->headers["authorization"] = $this->getAuthorization($request->pathname, $request->method, $request->query, $request->headers, $accessKeyId, $accessKeySecret);
    }

    /**
     * @param int[] $content
     * @param string $signatureVersion
     * @return string
     */
    public function MakeContentHash($content, $signatureVersion){
        if (StringUtil::equals($signatureVersion, "v4")) {
            // TODO: 这里应当检查 length == 0，但是还不支持。通常情况下也不会出现 body 设置了但是长度为 0
            if (Utils::isUnset($content)) {
                return '';
            }
            return StringUtil::toLower(EncodeUtil::hexEncode(EncodeUtil::hash($content, "SLS4-HMAC-SHA256")));
        }
        return StringUtil::toUpper(EncodeUtil::hexEncode(SignatureUtil::MD5SignForBytes($content)));
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
            $error = Utils::readAsJSON($response->body);
            $resMap = Utils::assertAsMap($error);
            throw new TeaError([
                "code" => @$resMap["errorCode"],
                "message" => @$resMap["errorMessage"],
                "accessDeniedDetail" => @$resMap["accessDeniedDetail"],
                "data" => [
                    "httpCode" => $response->statusCode,
                    "requestId" => @$response->headers["x-log-requestid"],
                    "statusCode" => $response->statusCode
                ]
            ]);
        }
        if (!Utils::isUnset($response->body)) {
            $bodyrawSize = @$response->headers["x-log-bodyrawsize"];
            $compressType = @$response->headers["x-log-compresstype"];
            $uncompressedData = $response->body;
            if (!Utils::isUnset($bodyrawSize) && !Utils::isUnset($compressType)) {
                $uncompressedData = DarabonbaGatewaySlsUtilClient::readAndUncompressBlock($response->body, $compressType, $bodyrawSize);
            }
            if (Utils::equalString($request->bodyType, "binary")) {
                $response->deserializedBody = $uncompressedData;
            }
            else if (Utils::equalString($request->bodyType, "byte")) {
                $byt = Utils::readAsBytes($uncompressedData);
                $response->deserializedBody = $byt;
            }
            else if (Utils::equalString($request->bodyType, "string")) {
                $response->deserializedBody = Utils::readAsString($uncompressedData);
            }
            else if (Utils::equalString($request->bodyType, "json")) {
                $obj = Utils::readAsJSON($uncompressedData);
                // var res = Util.assertAsMap(obj);
                $response->deserializedBody = $obj;
            }
            else if (Utils::equalString($request->bodyType, "array")) {
                $response->deserializedBody = Utils::readAsJSON($uncompressedData);
            }
            else {
                $response->deserializedBody = Utils::readAsString($uncompressedData);
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
            if (StringUtil::equals($network, "intranet")) {
                return "" . $regionId . "-intranet.log.aliyuncs.com";
            }
            else if (StringUtil::equals($network, "accelerate")) {
                return "log-global.aliyuncs.com";
            }
            else if (StringUtil::equals($network, "share")) {
                if (StringUtil::equals($regionId, "cn-hangzhou-corp") || StringUtil::equals($regionId, "cn-shanghai-corp")) {
                    return "" . $regionId . ".sls.aliyuncs.com";
                }
                else if (StringUtil::equals($regionId, "cn-zhangbei-corp")) {
                    return "zhangbei-corp-share.log.aliyuncs.com";
                }
                return "" . $regionId . "-share.log.aliyuncs.com";
            }
        }
        return "" . $regionId . ".log.aliyuncs.com";
    }

    /**
     * @param string $network
     * @param string $project
     * @param string $endpoint
     * @return string
     */
    public function getHost($network, $project, $endpoint){
        if (Utils::isUnset($project)) {
            return $endpoint;
        }
        return "" . $project . "." . $endpoint . "";
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
        $sign = $this->getSignature($pathname, $method, $query, $headers, $secret);
        return "LOG " . $ak . ":" . $sign . "";
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
        $resource = $pathname;
        $stringToSign = "";
        $canonicalizedResource = $this->buildCanonicalizedResource($resource, $query);
        $canonicalizedHeaders = $this->buildCanonicalizedHeaders($headers);
        $stringToSign = "" . $method . "\n" . $canonicalizedHeaders . "" . $canonicalizedResource . "";
        return EncodeUtil::base64EncodeToString(SignatureUtil::HmacSHA256Sign($stringToSign, $secret));
    }

    /**
     * @param string $pathname
     * @param string[] $query
     * @return string
     */
    public function buildCanonicalizedResource($pathname, $query){
        $canonicalizedResource = $pathname;
        if (!Utils::isUnset($query)) {
            $queryList = MapUtil::keySet($query);
            $sortedParams = ArrayUtil::ascSort($queryList);
            $separator = "?";
            foreach($sortedParams as $paramName){
                $canonicalizedResource = "" . $canonicalizedResource . "" . $separator . "" . $paramName . "";
                $paramValue = @$query[$paramName];
                if (!Utils::isUnset($paramValue)) {
                    $canonicalizedResource = "" . $canonicalizedResource . "=" . $paramValue . "";
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
            if (StringUtil::contains(StringUtil::toLower($header), "x-log-") || StringUtil::contains(StringUtil::toLower($header), "x-acs-")) {
                $canonicalizedHeaders = "" . $canonicalizedHeaders . "" . $header . ":" . @$headers[$header] . "\n";
            }
        }
        return $canonicalizedHeaders;
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
     * @param InterceptorContext $context
     * @param string $date
     * @param string $contentHash
     * @param string $accessKeyId
     * @param string $accessKeySecret
     * @return string
     */
    public function getAuthorizationV4($context, $date, $contentHash, $accessKeyId, $accessKeySecret){
        $region = $this->getRegion($context);
        $headerStr = $this->getSignedHeaderStrV4($context->request->headers);
        $dateShort = StringUtil::subString($date, 0, 8);
        $dateShort = StringUtil::replace($dateShort, "T", "", null);
        // for fix php sdk bug
        $scope = "" . $dateShort . "/" . $region . "/sls/aliyun_v4_request";
        $signingkey = $this->getSigningkeyV4("SLS4-HMAC-SHA256", $accessKeySecret, $region, $dateShort);
        $signature = $this->getSignatureV4($context, "SLS4-HMAC-SHA256", $headerStr, $date, $scope, $contentHash, $signingkey);
        return "" . "SLS4-HMAC-SHA256" . " Credential=" . $accessKeyId . "/" . $scope . ",Signature=" . $signature . "";
    }

    /**
     * @param string $signatureAlgorithm
     * @param string $secret
     * @param string $region
     * @param string $date
     * @return array
     */
    public function getSigningkeyV4($signatureAlgorithm, $secret, $region, $date){
        $sc1 = "aliyun_v4" . $secret . "";
        $sc2 = SignatureUtil::HmacSHA256Sign($date, $sc1);
        $sc3 = SignatureUtil::HmacSHA256SignByBytes($region, $sc2);
        $sc4 = SignatureUtil::HmacSHA256SignByBytes("sls", $sc3);
        return SignatureUtil::HmacSHA256SignByBytes("aliyun_v4_request", $sc4);
    }

    /**
     * @param InterceptorContext $context
     * @param string $signatureAlgorithm
     * @param string $signedHeaderStr
     * @param string $date
     * @param string $scope
     * @param string $contentSha256
     * @param int[] $signingkey
     * @return string
     */
    public function getSignatureV4($context, $signatureAlgorithm, $signedHeaderStr, $date, $scope, $contentSha256, $signingkey){
        $request = $context->request;
        $canonicalURI = "/";
        if (!Utils::empty_($request->pathname)) {
            $canonicalURI = $request->pathname;
        }
        $resources = $this->buildCanonicalizedResourceV4($request->query);
        $headers = $this->buildCanonicalizedHeadersV4($request->headers, $signedHeaderStr);
        $stringToHash = "" . $request->method . "\n" . $canonicalURI . "\n" . $resources . "\n" . $headers . "\n" . $signedHeaderStr . "\n" . $contentSha256 . "";
        $hex = EncodeUtil::hexEncode(EncodeUtil::hash(Utils::toBytes($stringToHash), $signatureAlgorithm));
        $stringToSign = "" . $signatureAlgorithm . "\n" . $date . "\n" . $scope . "\n" . $hex . "";
        $signature = SignatureUtil::HmacSHA256SignByBytes($stringToSign, $signingkey);
        return EncodeUtil::hexEncode($signature);
    }

    /**
     * @param string[] $query
     * @return string
     */
    public function buildCanonicalizedResourceV4($query){
        $canonicalizedResource = "";
        if (!Utils::isUnset($query)) {
            $queryArray = MapUtil::keySet($query);
            $sortedQueryArray = ArrayUtil::ascSort($queryArray);
            $separator = "";
            foreach($sortedQueryArray as $key){
                $canonicalizedResource = "" . $canonicalizedResource . "" . $separator . "" . $key . "";
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
     * @param string $signedHeaderStr
     * @return string
     */
    public function buildCanonicalizedHeadersV4($headers, $signedHeaderStr){
        $canonicalizedHeaders = "";
        $signedHeaders = StringUtil::split($signedHeaderStr, ";", null);
        foreach($signedHeaders as $header){
            $canonicalizedHeaders = "" . $canonicalizedHeaders . "" . $header . ":" . StringUtil::trim(@$headers[$header]) . "\n";
        }
        return $canonicalizedHeaders;
    }

    /**
     * @param string[] $headers
     * @return string
     */
    public function getSignedHeaderStrV4($headers){
        $headersArray = MapUtil::keySet($headers);
        $sortedHeadersArray = ArrayUtil::ascSort($headersArray);
        $tmp = "";
        $separator = "";
        foreach($sortedHeadersArray as $key){
            $lowerKey = StringUtil::toLower($key);
            if (StringUtil::hasPrefix($lowerKey, "x-log-") || StringUtil::hasPrefix($lowerKey, "x-acs-") || StringUtil::equals($lowerKey, "host") || StringUtil::equals($lowerKey, "content-type")) {
                $tmp = "" . $tmp . "" . $separator . "" . $lowerKey . "";
                $separator = ";";
            }
        }
        return $tmp;
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
        // try parse region from endpoint
        // do not use String.subString, subString has bug in php implementation
        $region = StringUtil::replace($config->endpoint, ".log.aliyuncs.com", "", null);
        $region = StringUtil::replace($region, ".sls.aliyuncs.com", "", null);
        if (StringUtil::equals($region, $config->endpoint)) {
            throw new TeaError([
                "code" => "ClientConfigError",
                "message" => "The regionId configuration of sls client is missing."
            ]);
        }
        $region = StringUtil::replace($region, "-share", "", null);
        $region = StringUtil::replace($region, "-intranet", "", null);
        $region = StringUtil::replace($region, "-vpc", "", null);
        return $region;
    }

    /**
     * format: YYYYMMDDTHHMMSSZ
     * @return string
     */
    public function getDateISO8601(){
        $date = OpenApiUtilClient::getTimestamp();
        // 2024-02-04T11:31:58Z
        $date = StringUtil::replace($date, "-", "", null);
        return StringUtil::replace($date, ":", "", null);
    }
}
