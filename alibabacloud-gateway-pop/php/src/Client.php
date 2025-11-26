<?php

// This file is auto-generated, don't edit it. Thanks.
namespace Darabonba\GatewayPop;

use Darabonba\GatewaySpi\Client as DarabonbaGatewaySpiClient;
use AlibabaCloud\Tea\Utils\Utils;
use AlibabaCloud\Tea\Exception\TeaError;
use AlibabaCloud\OpenApiUtil\OpenApiUtilClient;
use AlibabaCloud\Tea\Tea;
use AlibabaCloud\Darabonba\EncodeUtil\EncodeUtil;
use AlibabaCloud\Darabonba\String\StringUtil;
use AlibabaCloud\Tea\XML\XML;
use AlibabaCloud\Endpoint\Endpoint;
use AlibabaCloud\Darabonba\ArrayUtil\ArrayUtil;
use AlibabaCloud\Darabonba\SignatureUtil\SignatureUtil;
use AlibabaCloud\Darabonba\MapUtil\MapUtil;

use Darabonba\GatewaySpi\Models\InterceptorContext;
use Darabonba\GatewaySpi\Models\AttributeMap;

class Client extends DarabonbaGatewaySpiClient
{
    protected $_endpointSuffix;

    protected $_signatureTypePrefix;

    protected $_signPrefix;

    protected $_sha256;

    protected $_sm3;

    public function __construct()
    {
        parent::__construct();
        // CLOUD4-
        $this->_signatureTypePrefix = "ACS4-";
        // cloud_v4
        $this->_signPrefix = "aliyun_v4";
        $this->_endpointSuffix = "aliyuncs.com";
        $this->_sha256 = "" . $this->_signatureTypePrefix . "HMAC-SHA256";
        $this->_sm3 = "" . $this->_signatureTypePrefix . "HMAC-SM3";
    }

    /**
     * @param InterceptorContext $context
     * @param AttributeMap $attributeMap
     * @return void
     */
    public function modifyConfiguration($context, $attributeMap)
    {
        $request = $context->request;
        $config = $context->configuration;
        $attributes = $attributeMap->key;
        if (!Utils::isUnset($attributes)) {
            $this->_signatureTypePrefix = @$attributes["signatureTypePrefix"];
            $this->_signPrefix = @$attributes["signPrefix"];
            $this->_endpointSuffix = @$attributes["endpointSuffix"];
            $this->_sha256 = "" . $this->_signatureTypePrefix . "HMAC-SHA256";
            $this->_sm3 = "" . $this->_signatureTypePrefix . "HMAC-SM3";
        }
        $config->endpoint = $this->getEndpoint($request->productId, $config->regionId, $config->endpointRule, $config->network, $config->suffix, $config->endpointMap, $config->endpoint);
    }

    /**
     * @param InterceptorContext $context
     * @param AttributeMap $attributeMap
     * @return void
     * @throws TeaError
     */
    public function modifyRequest($context, $attributeMap)
    {
        $request = $context->request;
        $config = $context->configuration;
        $date = OpenApiUtilClient::getTimestamp();
        $request->headers = Tea::merge([
            "host" => $config->endpoint,
            "x-acs-version" => $request->version,
            "user-agent" => $request->userAgent,
            "x-acs-date" => $date,
            "x-acs-signature-nonce" => Utils::getNonce(),
            "accept" => "application/json"
        ], $request->headers);
        if (!Utils::empty_($request->action)) {
            $request->headers["x-acs-action"] = $request->action;
        }
        $signatureAlgorithm = Utils::defaultString($request->signatureAlgorithm, $this->_sha256);
        $hashedRequestPayload = EncodeUtil::hexEncode(EncodeUtil::hash(Utils::toBytes(""), $signatureAlgorithm));
        if (!Utils::isUnset($request->stream)) {
            $tmp = Utils::readAsBytes($request->stream);
            $hashedRequestPayload = EncodeUtil::hexEncode(EncodeUtil::hash($tmp, $signatureAlgorithm));
            $request->stream = $tmp;
            $request->headers["content-type"] = "application/octet-stream";
        } else {
            if (!Utils::isUnset($request->body)) {
                if (Utils::equalString($request->reqBodyType, "byte")) {
                    $byteObj = Utils::assertAsBytes($request->body);
                    $hashedRequestPayload = EncodeUtil::hexEncode(EncodeUtil::hash($byteObj, $signatureAlgorithm));
                    $request->stream = $byteObj;
                } else if (Utils::equalString($request->reqBodyType, "json")) {
                    $jsonObj = Utils::toJSONString($request->body);
                    $hashedRequestPayload = EncodeUtil::hexEncode(EncodeUtil::hash(Utils::toBytes($jsonObj), $signatureAlgorithm));
                    $request->stream = $jsonObj;
                    $request->headers["content-type"] = "application/json; charset=utf-8";
                } else {
                    $m = Utils::assertAsMap($request->body);
                    $formObj = OpenApiUtilClient::toForm($m);
                    $hashedRequestPayload = EncodeUtil::hexEncode(EncodeUtil::hash(Utils::toBytes($formObj), $signatureAlgorithm));
                    $request->stream = $formObj;
                    $request->headers["content-type"] = "application/x-www-form-urlencoded";
                }
            }
        }
        if (Utils::equalString($signatureAlgorithm, $this->_sm3)) {
            $request->headers["x-acs-content-sm3"] = $hashedRequestPayload;
        } else {
            $request->headers["x-acs-content-sha256"] = $hashedRequestPayload;
        }
        if (!Utils::equalString($request->authType, "Anonymous")) {
            $credential = $request->credential;
            if (Utils::isUnset($credential)) {
                throw new TeaError([
                    "code" => "ParameterMissing",
                    "message" => "'config.credential' can not be unset"
                ]);
            }
            $credentialModel = $credential->getCredential();
            if (!Utils::empty_($credentialModel->providerName)) {
                $request->headers["x-acs-credentials-provider"] = $credentialModel->providerName;
            }
            $authType = $credentialModel->type;
            if (Utils::equalString($authType, "bearer")) {
                $bearerToken = $credentialModel->bearerToken;
                $request->headers["x-acs-bearer-token"] = $bearerToken;
                $request->headers["x-acs-signature-type"] = "BEARERTOKEN";
                $request->headers["Authorization"] = "Bearer " . $bearerToken . "";
            } else if (Utils::equalString($authType, "id_token")) {
                $idToken = $credentialModel->securityToken;
                $request->headers["x-acs-zero-trust-idtoken"] = $idToken;
            } else {
                $accessKeyId = $credentialModel->accessKeyId;
                $accessKeySecret = $credentialModel->accessKeySecret;
                $securityToken = $credentialModel->securityToken;
                if (!Utils::empty_($securityToken)) {
                    $request->headers["x-acs-accesskey-id"] = $accessKeyId;
                    $request->headers["x-acs-security-token"] = $securityToken;
                }
                $dateNew = StringUtil::subString($date, 0, 10);
                $dateNew = StringUtil::replace($dateNew, "-", "", null);
                $region = $this->getRegion($request->productId, $config->endpoint, $config->regionId);
                $signingkey = $this->getSigningkey($signatureAlgorithm, $accessKeySecret, $request->productId, $region, $dateNew);
                $request->headers["Authorization"] = $this->getAuthorization($request->pathname, $request->method, $request->query, $request->headers, $signatureAlgorithm, $hashedRequestPayload, $accessKeyId, $signingkey, $request->productId, $region, $dateNew);
            }
        }
    }

    /**
     * @param InterceptorContext $context
     * @param AttributeMap $attributeMap
     * @return void
     * @throws TeaError
     */
    public function modifyResponse($context, $attributeMap)
    {
        $request = $context->request;
        $response = $context->response;
        if (Utils::is4xx($response->statusCode) || Utils::is5xx($response->statusCode)) {
            $err = [];
            if (!Utils::isUnset(@$response->headers["content-type"]) && StringUtil::contains(@$response->headers["content-type"], "text/xml")) {
                $_str = Utils::readAsString($response->body);
                $respMap = XML::parseXml($_str, null);
                $err = Utils::assertAsMap(@$respMap["Error"]);
            } else {
                $_res = Utils::readAsJSON($response->body);
                $err = Utils::assertAsMap($_res);
            }
            $requestId = $this->defaultAny(@$err["RequestId"], @$err["requestId"]);
            if (!Utils::isUnset(@$response->headers["x-acs-request-id"])) {
                $requestId = @$response->headers["x-acs-request-id"];
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
        } else if (Utils::equalString($request->bodyType, "binary")) {
            $response->deserializedBody = $response->body;
        } else if (Utils::equalString($request->bodyType, "byte")) {
            $byt = Utils::readAsBytes($response->body);
            $response->deserializedBody = $byt;
        } else if (Utils::equalString($request->bodyType, "string")) {
            $str = Utils::readAsString($response->body);
            $response->deserializedBody = $str;
        } else if (Utils::equalString($request->bodyType, "json")) {
            $obj = Utils::readAsJSON($response->body);
            $res = Utils::assertAsMap($obj);
            $response->deserializedBody = $res;
        } else if (Utils::equalString($request->bodyType, "array")) {
            $arr = Utils::readAsJSON($response->body);
            $response->deserializedBody = $arr;
        } else {
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
    public function getEndpoint($productId, $regionId, $endpointRule, $network, $suffix, $endpointMap, $endpoint)
    {
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
    public function defaultAny($inputValue, $defaultValue)
    {
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
     * @param int[] $signingkey
     * @param string $product
     * @param string $region
     * @param string $date
     * @return string
     */
    public function getAuthorization($pathname, $method, $query, $headers, $signatureAlgorithm, $payload, $ak, $signingkey, $product, $region, $date)
    {
        $signature = $this->getSignature($pathname, $method, $query, $headers, $signatureAlgorithm, $payload, $signingkey);
        $signedHeaders = $this->getSignedHeaders($headers);
        $signedHeadersStr = ArrayUtil::join($signedHeaders, ";");
        return "" . $signatureAlgorithm . " Credential=" . $ak . "/" . $date . "/" . $region . "/" . $product . "/" . $this->_signPrefix . "_request,SignedHeaders=" . $signedHeadersStr . ",Signature=" . $signature . "";
    }

    /**
     * @param string $pathname
     * @param string $method
     * @param string[] $query
     * @param string[] $headers
     * @param string $signatureAlgorithm
     * @param string $payload
     * @param int[] $signingkey
     * @return string
     */
    public function getSignature($pathname, $method, $query, $headers, $signatureAlgorithm, $payload, $signingkey)
    {
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
        if (Utils::equalString($signatureAlgorithm, $this->_sha256)) {
            $signature = SignatureUtil::HmacSHA256SignByBytes($stringToSign, $signingkey);
        } else if (Utils::equalString($signatureAlgorithm, $this->_sm3)) {
            $signature = SignatureUtil::HmacSM3SignByBytes($stringToSign, $signingkey);
        }
        return EncodeUtil::hexEncode($signature);
    }

    /**
     * @param string $signatureAlgorithm
     * @param string $secret
     * @param string $product
     * @param string $region
     * @param string $date
     * @return array
     */
    public function getSigningkey($signatureAlgorithm, $secret, $product, $region, $date)
    {
        $sc1 = "" . $this->_signPrefix . "" . $secret . "";
        $sc2 = Utils::toBytes("");
        if (Utils::equalString($signatureAlgorithm, $this->_sha256)) {
            $sc2 = SignatureUtil::HmacSHA256Sign($date, $sc1);
        } else if (Utils::equalString($signatureAlgorithm, $this->_sm3)) {
            $sc2 = SignatureUtil::HmacSM3Sign($date, $sc1);
        }
        $sc3 = Utils::toBytes("");
        if (Utils::equalString($signatureAlgorithm, $this->_sha256)) {
            $sc3 = SignatureUtil::HmacSHA256SignByBytes($region, $sc2);
        } else if (Utils::equalString($signatureAlgorithm, $this->_sm3)) {
            $sc3 = SignatureUtil::HmacSM3SignByBytes($region, $sc2);
        }
        $sc4 = Utils::toBytes("");
        if (Utils::equalString($signatureAlgorithm, $this->_sha256)) {
            $sc4 = SignatureUtil::HmacSHA256SignByBytes($product, $sc3);
        } else if (Utils::equalString($signatureAlgorithm, $this->_sm3)) {
            $sc4 = SignatureUtil::HmacSM3SignByBytes($product, $sc3);
        }
        $hmac = Utils::toBytes("");
        if (Utils::equalString($signatureAlgorithm, $this->_sha256)) {
            $hmac = SignatureUtil::HmacSHA256SignByBytes("" . $this->_signPrefix . "_request", $sc4);
        } else if (Utils::equalString($signatureAlgorithm, $this->_sm3)) {
            $hmac = SignatureUtil::HmacSM3SignByBytes("" . $this->_signPrefix . "_request", $sc4);
        }
        return $hmac;
    }

    /**
     * @param string $product
     * @param string $endpoint
     * @param string $regionId
     * @return string
     */
    public function getRegion($product, $endpoint, $regionId)
    {
        if (!Utils::empty_($regionId)) {
            return $regionId;
        }
        $region = "center";
        if (Utils::empty_($product) || Utils::empty_($endpoint)) {
            return $region;
        }
        $strs = StringUtil::split($endpoint, ":", null);
        $withoutPort = @$strs[0];
        $preRegion = StringUtil::replace($withoutPort, "." . $this->_endpointSuffix . "", "", null);
        $nodes = StringUtil::split($preRegion, ".", null);
        if (Utils::equalNumber(ArrayUtil::size($nodes), 2)) {
            $region = @$nodes[1];
        }
        return $region;
    }

    /**
     * @param string[] $query
     * @return string
     */
    public function buildCanonicalizedResource($query)
    {
        $canonicalizedResource = "";
        if (!Utils::isUnset($query)) {
            $queryArray = MapUtil::keySet($query);
            $sortedQueryArray = ArrayUtil::ascSort($queryArray);
            $separator = "";
            foreach ($sortedQueryArray as $key) {
                $canonicalizedResource = "" . $canonicalizedResource . "" . $separator . "" . EncodeUtil::percentEncode($key) . "=";
                if (!Utils::empty_(@$query[$key])) {
                    $canonicalizedResource = "" . $canonicalizedResource . "" . EncodeUtil::percentEncode(@$query[$key]) . "";
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
    public function buildCanonicalizedHeaders($headers)
    {
        // lower header key
        $headersArray = MapUtil::keySet($headers);
        $newHeaders = [];
        $tmp = "";
        foreach ($headersArray as $key) {
            $lowerKey = StringUtil::toLower($key);
            $value = @$headers[$key];
            if (!Utils::isUnset($value)) {
                if (!StringUtil::contains($tmp, $lowerKey) || StringUtil::equals($lowerKey, "host")) {
                    $tmp = "" . $tmp . "," . $lowerKey . "";
                    $newHeaders[$lowerKey] = StringUtil::trim($value);
                } else {
                    $newHeaders[$lowerKey] = "" . @$newHeaders[$lowerKey] . "," . StringUtil::trim($value) . "";
                }
            }
        }
        $canonicalizedHeaders = "";
        $sortedHeaders = $this->getSignedHeaders($headers);
        foreach ($sortedHeaders as $header) {
            $canonicalizedHeaders = "" . $canonicalizedHeaders . "" . $header . ":" . @$newHeaders[$header] . "\n";
        }
        return $canonicalizedHeaders;
    }

    /**
     * @param string[] $headers
     * @return array
     */
    public function getSignedHeaders($headers)
    {
        $headersArray = MapUtil::keySet($headers);
        $newHeadersArray = [];
        foreach ($headersArray as $key) {
            $lowerKey = StringUtil::toLower($key);
            $value = @$headers[$key];
            if (!Utils::isUnset($value)) {
                ArrayUtil::append($newHeadersArray, $lowerKey);
            }
        }
        $sortedHeadersArray = ArrayUtil::ascSort($newHeadersArray);
        $tmp = "";
        $separator = "";
        foreach ($sortedHeadersArray as $key) {
            if (StringUtil::hasPrefix($key, "x-acs-") || StringUtil::equals($key, "host") || StringUtil::equals($key, "content-type")) {
                if (!StringUtil::contains($tmp, $key)) {
                    $tmp = "" . $tmp . "" . $separator . "" . $key . "";
                    $separator = ";";
                }
            }
        }
        return StringUtil::split($tmp, ";", null);
    }
}
