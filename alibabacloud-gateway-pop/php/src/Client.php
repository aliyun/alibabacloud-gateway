<?php

// This file is auto-generated, don't edit it. Thanks.
namespace Darabonba\GatewayPop;

use Darabonba\GatewaySpi\Client as DarabonbaGatewaySpiClient;
use AlibabaCloud\Tea\Exception\TeaError;
use AlibabaCloud\Tea\Tea;
use AlibabaCloud\Tea\XML\XML;

use Darabonba\GatewaySpi\Models\InterceptorContext;
use Darabonba\GatewaySpi\Models\AttributeMap;

class Client extends DarabonbaGatewaySpiClient
{
    use Runtime;

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
        if (!self::isUnset($attributes)) {
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
        $date = self::getTimestamp();
        $request->headers = Tea::merge([
            "host" => $config->endpoint,
            "x-acs-version" => $request->version,
            "user-agent" => $request->userAgent,
            "x-acs-date" => $date,
            "x-acs-signature-nonce" => self::getNonce(),
            "accept" => "application/json"
        ], $request->headers);
        if (!self::empty_($request->action)) {
            $request->headers["x-acs-action"] = $request->action;
        }
        $signatureAlgorithm = self::defaultString($request->signatureAlgorithm, $this->_sha256);
        $hashedRequestPayload = self::hexEncode(self::hash(self::toBytes(""), $signatureAlgorithm));
        if (!self::isUnset($request->stream)) {
            $tmp = self::readAsBytes($request->stream);
            $hashedRequestPayload = self::hexEncode(self::hash($tmp, $signatureAlgorithm));
            $request->stream = $tmp;
            $request->headers["content-type"] = "application/octet-stream";
        } else {
            if (!self::isUnset($request->body)) {
                if (self::equalString($request->reqBodyType, "byte")) {
                    $byteObj = self::assertAsBytes($request->body);
                    $hashedRequestPayload = self::hexEncode(self::hash($byteObj, $signatureAlgorithm));
                    $request->stream = $byteObj;
                } else if (self::equalString($request->reqBodyType, "json")) {
                    $jsonObj = self::toJSONString($request->body);
                    $hashedRequestPayload = self::hexEncode(self::hash(self::toBytes($jsonObj), $signatureAlgorithm));
                    $request->stream = $jsonObj;
                    $request->headers["content-type"] = "application/json; charset=utf-8";
                } else {
                    $m = self::assertAsMap($request->body);
                    $formObj = self::toForm($m);
                    $hashedRequestPayload = self::hexEncode(self::hash(self::toBytes($formObj), $signatureAlgorithm));
                    $request->stream = $formObj;
                    $request->headers["content-type"] = "application/x-www-form-urlencoded";
                }
            }
        }
        if (self::equalString($signatureAlgorithm, $this->_sm3)) {
            $request->headers["x-acs-content-sm3"] = $hashedRequestPayload;
        } else {
            $request->headers["x-acs-content-sha256"] = $hashedRequestPayload;
        }
        if (!self::equalString($request->authType, "Anonymous")) {
            $credential = $request->credential;
            if (self::isUnset($credential)) {
                throw new TeaError([
                    "code" => "ParameterMissing",
                    "message" => "'config.credential' can not be unset"
                ]);
            }
            $credentialModel = $credential->getCredential();
            if (!self::empty_($credentialModel->providerName)) {
                $request->headers["x-acs-credentials-provider"] = $credentialModel->providerName;
            }
            $authType = $credentialModel->type;
            if (self::equalString($authType, "bearer")) {
                $bearerToken = $credentialModel->bearerToken;
                $request->headers["x-acs-bearer-token"] = $bearerToken;
                $request->headers["x-acs-signature-type"] = "BEARERTOKEN";
                $request->headers["Authorization"] = "Bearer " . $bearerToken . "";
            } else if (self::equalString($authType, "id_token")) {
                $idToken = $credentialModel->securityToken;
                $request->headers["x-acs-zero-trust-idtoken"] = $idToken;
            } else {
                $accessKeyId = $credentialModel->accessKeyId;
                $accessKeySecret = $credentialModel->accessKeySecret;
                $securityToken = $credentialModel->securityToken;
                if (!self::empty_($securityToken)) {
                    $request->headers["x-acs-accesskey-id"] = $accessKeyId;
                    $request->headers["x-acs-security-token"] = $securityToken;
                }
                $dateNew = self::subString($date, 0, 10);
                $dateNew = self::replace($dateNew, "-", "", null);
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
        if (self::is4xx($response->statusCode) || self::is5xx($response->statusCode)) {
            $err = [];
            $contentType = isset($response->headers["content-type"]) ? $response->headers["content-type"] : null;
            if (!self::isUnset($contentType) && self::contains($contentType, "text/xml")) {
                $_str = self::readAsString($response->body);
                $respMap = XML::parseXml($_str, null);
                $err = self::assertAsMap(@$respMap["Error"]);
            } else {
                $_res = self::readAsJSON($response->body);
                $err = self::assertAsMap($_res);
            }
            $requestId = $this->defaultAny(@$err["RequestId"], @$err["requestId"]);
            $requestIdHeader = isset($response->headers["x-acs-request-id"]) ? $response->headers["x-acs-request-id"] : null;
            if (!self::isUnset($requestIdHeader)) {
                $requestId = $requestIdHeader;
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
        if (self::equalNumber($response->statusCode, 204)) {
            self::readAsString($response->body);
        } else if (self::equalString($request->bodyType, "binary")) {
            $response->deserializedBody = $response->body;
        } else if (self::equalString($request->bodyType, "byte")) {
            $byt = self::readAsBytes($response->body);
            $response->deserializedBody = $byt;
        } else if (self::equalString($request->bodyType, "string")) {
            $str = self::readAsString($response->body);
            $response->deserializedBody = $str;
        } else if (self::equalString($request->bodyType, "json")) {
            $obj = self::readAsJSON($response->body);
            $res = self::assertAsMap($obj);
            $response->deserializedBody = $res;
        } else if (self::equalString($request->bodyType, "array")) {
            $arr = self::readAsJSON($response->body);
            $response->deserializedBody = $arr;
        } else {
            $response->deserializedBody = self::readAsString($response->body);
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
        if (!self::empty_($endpoint)) {
            return $endpoint;
        }
        if (!self::isUnset($endpointMap) && !self::empty_(@$endpointMap[$regionId])) {
            return @$endpointMap[$regionId];
        }
        return self::getEndpointRules($productId, $regionId, $endpointRule, $network, $suffix);
    }

    /**
     * @param mixed $inputValue
     * @param mixed $defaultValue
     * @return any
     */
    public function defaultAny($inputValue, $defaultValue)
    {
        if (self::isUnset($inputValue)) {
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
        $signedHeadersStr = self::join($signedHeaders, ";");
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
        if (!self::empty_($pathname)) {
            $canonicalURI = $pathname;
        }
        $stringToSign = "";
        $canonicalizedResource = $this->buildCanonicalizedResource($query);
        $canonicalizedHeaders = $this->buildCanonicalizedHeaders($headers);
        $signedHeaders = $this->getSignedHeaders($headers);
        $this->validateSignedHeaders($signedHeaders);
        $signedHeadersStr = self::join($signedHeaders, ";");
        $stringToSign = "" . $method . "\n" . $canonicalURI . "\n" . $canonicalizedResource . "\n" . $canonicalizedHeaders . "\n" . $signedHeadersStr . "\n" . $payload . "";
        $hex = self::hexEncode(self::hash(self::toBytes($stringToSign), $signatureAlgorithm));
        $stringToSign = "" . $signatureAlgorithm . "\n" . $hex . "";
        $signature = self::toBytes("");
        if (self::equalString($signatureAlgorithm, $this->_sha256)) {
            $signature = self::HmacSHA256SignByBytes($stringToSign, $signingkey);
        } else if (self::equalString($signatureAlgorithm, $this->_sm3)) {
            $signature = self::HmacSM3SignByBytes($stringToSign, $signingkey);
        }
        return self::hexEncode($signature);
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
        $sc2 = self::toBytes("");
        if (self::equalString($signatureAlgorithm, $this->_sha256)) {
            $sc2 = self::HmacSHA256Sign($date, $sc1);
        } else if (self::equalString($signatureAlgorithm, $this->_sm3)) {
            $sc2 = self::HmacSM3Sign($date, $sc1);
        }
        $sc3 = self::toBytes("");
        if (self::equalString($signatureAlgorithm, $this->_sha256)) {
            $sc3 = self::HmacSHA256SignByBytes($region, $sc2);
        } else if (self::equalString($signatureAlgorithm, $this->_sm3)) {
            $sc3 = self::HmacSM3SignByBytes($region, $sc2);
        }
        $sc4 = self::toBytes("");
        if (self::equalString($signatureAlgorithm, $this->_sha256)) {
            $sc4 = self::HmacSHA256SignByBytes($product, $sc3);
        } else if (self::equalString($signatureAlgorithm, $this->_sm3)) {
            $sc4 = self::HmacSM3SignByBytes($product, $sc3);
        }
        $hmac = self::toBytes("");
        if (self::equalString($signatureAlgorithm, $this->_sha256)) {
            $hmac = self::HmacSHA256SignByBytes("" . $this->_signPrefix . "_request", $sc4);
        } else if (self::equalString($signatureAlgorithm, $this->_sm3)) {
            $hmac = self::HmacSM3SignByBytes("" . $this->_signPrefix . "_request", $sc4);
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
        if (!self::empty_($regionId)) {
            return $regionId;
        }
        $region = "center";
        if (self::empty_($product) || self::empty_($endpoint)) {
            return $region;
        }
        $strs = self::split($endpoint, ":", null);
        $withoutPort = @$strs[0];
        $preRegion = self::replace($withoutPort, "." . $this->_endpointSuffix . "", "", null);
        $nodes = self::split($preRegion, ".", null);
        if (self::equalNumber(self::size($nodes), 2)) {
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
        if (!self::isUnset($query)) {
            $queryArray = self::keySet($query);
            $sortedQueryArray = self::ascSort($queryArray);
            $separator = "";
            foreach ($sortedQueryArray as $key) {
                $canonicalizedResource = "" . $canonicalizedResource . "" . $separator . "" . self::percentEncode($key) . "=";
                if (!self::empty_(@$query[$key])) {
                    $canonicalizedResource = "" . $canonicalizedResource . "" . self::percentEncode(@$query[$key]) . "";
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
        $headersArray = self::keySet($headers);
        $newHeaders = [];
        foreach ($headersArray as $key) {
            $lowerKey = self::toLower($key);
            $value = @$headers[$key];
            if (!self::isUnset($value)) {
                $existingHeader = array_key_exists($lowerKey, $newHeaders) ? $newHeaders[$lowerKey] : null;
                if (self::isUnset($existingHeader) || self::equals($lowerKey, "host")) {
                    $newHeaders[$lowerKey] = self::trim($value);
                } else {
                    $newHeaders[$lowerKey] = "" . @$newHeaders[$lowerKey] . "," . self::trim($value) . "";
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
    /**
     * @param string[] $signedHeaders
     * @return void
     */
    public function validateSignedHeaders($signedHeaders)
    {
        if (!self::contains($signedHeaders, "host") || !self::contains($signedHeaders, "x-acs-date")) {
            throw new TeaError([
                "code" => "InvalidSignedHeaders",
                "message" => "signed headers must include host and x-acs-date"
            ]);
        }
    }

    /**
     * @param string[] $headers
     * @return array
     */
    public function getSignedHeaders($headers)
    {
        $headersArray = self::keySet($headers);
        $newHeadersArray = [];
        foreach ($headersArray as $key) {
            $lowerKey = self::toLower($key);
            $value = @$headers[$key];
            if (!self::isUnset($value)) {
                self::append($newHeadersArray, $lowerKey);
            }
        }
        $sortedHeadersArray = self::ascSort($newHeadersArray);
        $result = [];
        foreach ($sortedHeadersArray as $key) {
            if (self::hasPrefix($key, "x-acs-") || self::equals($key, "host") || self::equals($key, "content-type")) {
                if (!self::contains($result, $key)) {
                    self::append($result, $key);
                }
            }
        }
        return $result;
    }
}
