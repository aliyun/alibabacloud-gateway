<?php

// This file is auto-generated, don't edit it. Thanks.
namespace Darabonba\GatewayPds;

use Darabonba\GatewaySpi\Client as DarabonbaGatewaySpiClient;
use AlibabaCloud\Tea\Utils\Utils;
use AlibabaCloud\Tea\Tea;
use AlibabaCloud\Darabonba\EncodeUtil\EncodeUtil;
use AlibabaCloud\OpenApiUtil\OpenApiUtilClient;
use AlibabaCloud\Darabonba\String\StringUtil;
use AlibabaCloud\Tea\Exception\TeaError;
use AlibabaCloud\Darabonba\SignatureUtil\SignatureUtil;
use AlibabaCloud\Darabonba\MapUtil\MapUtil;
use AlibabaCloud\Darabonba\ArrayUtil\ArrayUtil;

use Darabonba\GatewaySpi\Models\InterceptorContext;
use Darabonba\GatewaySpi\Models\AttributeMap;

class Client extends DarabonbaGatewaySpiClient
{
    public function __construct()
    {
        parent::__construct();
    }

    /**
     * @param InterceptorContext $context
     * @param AttributeMap $attributeMap
     * @return void
     */
    public function modifyConfiguration($context, $attributeMap)
    {
        $config = $context->configuration;
        $config->endpoint = $this->getEndpoint($config->network, $config->endpoint);
    }

    /**
     * @param InterceptorContext $context
     * @param AttributeMap $attributeMap
     * @return void
     */
    public function modifyRequest($context, $attributeMap)
    {
        $request = $context->request;
        $config = $context->configuration;
        $date = Utils::getDateUTCString();
        $request->headers = Tea::merge([
            "date" => $date,
            "host" => $config->endpoint,
            "x-acs-version" => $request->version,
            "x-acs-action" => $request->action,
            "user-agent" => $request->userAgent,
            "x-acs-signature-nonce" => Utils::getNonce(),
            "accept" => "application/json"
        ], $request->headers);
        $signatureAlgorithm = Utils::defaultString($request->signatureAlgorithm, "ACS4-HMAC-SHA256");
        $signatureVersion = Utils::defaultString($request->signatureVersion, "v1");
        $hashedRequestPayload = EncodeUtil::hexEncode(EncodeUtil::hash(Utils::toBytes(""), $signatureAlgorithm));
        if (!Utils::isUnset($request->stream)) {
            $tmp = Utils::readAsBytes($request->stream);
            $hashedRequestPayload = EncodeUtil::hexEncode(EncodeUtil::hash($tmp, $signatureAlgorithm));
            $request->stream = $tmp;
            $request->headers["content-type"] = "application/octet-stream";
        } else {
            if (!Utils::isUnset($request->body)) {
                if (Utils::equalString($request->reqBodyType, "json")) {
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
        if (StringUtil::equals($signatureVersion, "v4")) {
            if (Utils::equalString($signatureAlgorithm, "ACS4-HMAC-SM3")) {
                $request->headers["x-acs-content-sm3"] = $hashedRequestPayload;
            } else {
                $request->headers["x-acs-content-sha256"] = $hashedRequestPayload;
            }
        } else {
            $request->headers["x-acs-signature-method"] = "HMAC-SHA1";
            $request->headers["x-acs-signature-version"] = "1.0";
        }
        if (!Utils::equalString($request->authType, "Anonymous") && !Utils::isUnset($request->credential)) {
            $credential = $request->credential;
            $credentialModel = $credential->getCredential();
            $authType = $credentialModel->type;
            if (Utils::equalString($authType, "bearer")) {
                $bearerToken = $credentialModel->bearerToken;
                $request->headers["x-acs-bearer-token"] = $bearerToken;
                $request->headers["Authorization"] = "Bearer " . $bearerToken . "";
            } else {
                $accessKeyId = $credentialModel->accessKeyId;
                $accessKeySecret = $credentialModel->accessKeySecret;
                $securityToken = $credentialModel->securityToken;
                if (!Utils::empty_($securityToken)) {
                    $request->headers["x-acs-security-token"] = $securityToken;
                }
                if (StringUtil::equals($signatureVersion, "v4")) {
                    $dateNew = StringUtil::subString($date, 0, 10);
                    $region = $this->getRegion($config->endpoint);
                    $signingkey = $this->getSigningkey($signatureAlgorithm, $accessKeySecret, $region, $dateNew);
                    $request->headers["Authorization"] = $this->getAuthorizationV4($request->pathname, $request->method, $request->query, $request->headers, $signatureAlgorithm, $hashedRequestPayload, $accessKeyId, $signingkey, $request->productId, $region, $dateNew);
                } else {
                    $request->headers["Authorization"] = $this->getAuthorization($request->pathname, $request->method, $request->query, $request->headers, $accessKeyId, $accessKeySecret);
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
    public function modifyResponse($context, $attributeMap)
    {
        $request = $context->request;
        $response = $context->response;
        if (Utils::is4xx($response->statusCode) || Utils::is5xx($response->statusCode)) {
            $_res = Utils::readAsJSON($response->body);
            $err = Utils::assertAsMap($_res);
            $headers = $response->headers;
            $requestId = @$headers["x-ca-request-id"];
            @$err["statusCode"] = $response->statusCode;
            @$err["requestId"] = $requestId;
            throw new TeaError([
                "code" => "" . (string) ($this->defaultAny(@$err["Code"], @$err["code"])) . "",
                "message" => "" . (string) ($this->defaultAny(@$err["Message"], @$err["message"])) . "",
                "data" => $err
            ]);
        }
        if (!Utils::isUnset($response->body)) {
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
                $response->deserializedBody = Utils::readAsJSON($response->body);
            } else if (Utils::equalString($request->bodyType, "array")) {
                $response->deserializedBody = Utils::readAsJSON($response->body);
            } else {
                $response->deserializedBody = Utils::readAsString($response->body);
            }
        }
    }

    /**
     * @param string $network
     * @param string $endpoint
     * @return string
     */
    public function getEndpoint($network, $endpoint)
    {
        $realEndpoint = "api.aliyunpds.com";
        if (!Utils::empty_($endpoint)) {
            $realEndpoint = $endpoint;
        }
        if (!Utils::empty_($network) && StringUtil::equals($network, "vpc")) {
            $realEndpoint = StringUtil::replace($realEndpoint, "api.aliyunpds.com", "api-vpc.aliyunpds.com", null);
            $realEndpoint = StringUtil::replace($realEndpoint, "admin.aliyunpds.com", "admin-vpc.aliyunpds.com", null);
        }
        return $realEndpoint;
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
     * @param string $ak
     * @param string $secret
     * @return string
     */
    public function getAuthorization($pathname, $method, $query, $headers, $ak, $secret)
    {
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
    public function getSignature($pathname, $method, $query, $headers, $secret)
    {
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
    public function buildCanonicalizedResource($pathname, $query)
    {
        $canonicalizedResource = $pathname;
        if (!Utils::isUnset($query)) {
            $queryArray = MapUtil::keySet($query);
            $sortedQueryArray = ArrayUtil::ascSort($queryArray);
            $separator = "?";
            foreach ($sortedQueryArray as $key) {
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
    public function buildCanonicalizedHeaders($headers)
    {
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
        foreach ($sortedHeaders as $header) {
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
    public function getSignedHeaders($headers)
    {
        $headersArray = MapUtil::keySet($headers);
        $sortedHeadersArray = ArrayUtil::ascSort($headersArray);
        $tmp = "";
        $separator = "";
        foreach ($sortedHeadersArray as $key) {
            $lowerKey = StringUtil::toLower($key);
            if (StringUtil::hasPrefix($lowerKey, "x-acs-")) {
                if (!StringUtil::contains($tmp, $lowerKey)) {
                    $tmp = "" . $tmp . "" . $separator . "" . $lowerKey . "";
                    $separator = ";";
                }
            }
        }
        return StringUtil::split($tmp, ";", null);
    }

    /**
     * @param string $endpoint
     * @return string
     */
    public function getRegion($endpoint)
    {
        $region = "center";
        if (Utils::empty_($endpoint)) {
            return $region;
        }
        if (StringUtil::contains($endpoint, ".admin.aliyunpds.com")) {
            $region = StringUtil::replace($endpoint, ".admin.aliyunpds.com", "", null);
        }
        return $region;
    }

    /**
     * @param string $signatureAlgorithm
     * @param string $secret
     * @param string $region
     * @param string $date
     * @return array
     */
    public function getSigningkey($signatureAlgorithm, $secret, $region, $date)
    {
        $sc1 = "aliyun_v4" . $secret . "";
        $sc2 = Utils::toBytes("");
        if (Utils::equalString($signatureAlgorithm, "ACS4-HMAC-SHA256")) {
            $sc2 = SignatureUtil::HmacSHA256Sign($date, $sc1);
        } else if (Utils::equalString($signatureAlgorithm, "ACS4-HMAC-SM3")) {
            $sc2 = SignatureUtil::HmacSM3Sign($date, $sc1);
        }
        $sc3 = Utils::toBytes("");
        if (Utils::equalString($signatureAlgorithm, "ACS4-HMAC-SHA256")) {
            $sc3 = SignatureUtil::HmacSHA256SignByBytes($region, $sc2);
        } else if (Utils::equalString($signatureAlgorithm, "ACS4-HMAC-SM3")) {
            $sc3 = SignatureUtil::HmacSM3SignByBytes($region, $sc2);
        }
        $sc4 = Utils::toBytes("");
        if (Utils::equalString($signatureAlgorithm, "ACS4-HMAC-SHA256")) {
            $sc4 = SignatureUtil::HmacSHA256SignByBytes("pds", $sc3);
        } else if (Utils::equalString($signatureAlgorithm, "ACS4-HMAC-SM3")) {
            $sc4 = SignatureUtil::HmacSM3SignByBytes("pds", $sc3);
        }
        $hmac = Utils::toBytes("");
        if (Utils::equalString($signatureAlgorithm, "ACS4-HMAC-SHA256")) {
            $hmac = SignatureUtil::HmacSHA256SignByBytes("aliyun_v4_request", $sc4);
        } else if (Utils::equalString($signatureAlgorithm, "ACS4-HMAC-SM3")) {
            $hmac = SignatureUtil::HmacSM3SignByBytes("aliyun_v4_request", $sc4);
        }
        return $hmac;
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
    public function getAuthorizationV4($pathname, $method, $query, $headers, $signatureAlgorithm, $payload, $ak, $signingkey, $product, $region, $date)
    {
        $signature = $this->getSignatureV4($pathname, $method, $query, $headers, $signatureAlgorithm, $payload, $signingkey);
        $signedHeaders = $this->getSignedHeaders($headers);
        $signedHeadersStr = ArrayUtil::join($signedHeaders, ";");
        return "" . $signatureAlgorithm . " Credential=" . $ak . "/" . $date . "/" . $region . "/" . $product . "/aliyun_v4_request,SignedHeaders=" . $signedHeadersStr . ",Signature=" . $signature . "";
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
    public function getSignatureV4($pathname, $method, $query, $headers, $signatureAlgorithm, $payload, $signingkey)
    {
        $stringToSign = "";
        $canonicalizedResource = $this->buildCanonicalizedResource($pathname, $query);
        $canonicalizedHeaders = $this->buildCanonicalizedHeaders($headers);
        $signedHeaders = $this->getSignedHeaders($headers);
        $signedHeadersStr = ArrayUtil::join($signedHeaders, ";");
        $stringToSign = "" . $method . "\n" . $canonicalizedResource . "\n" . $canonicalizedHeaders . "\n" . $signedHeadersStr . "\n" . $payload . "";
        $hex = EncodeUtil::hexEncode(EncodeUtil::hash(Utils::toBytes($stringToSign), $signatureAlgorithm));
        $stringToSign = "" . $signatureAlgorithm . "\n" . $hex . "";
        $signature = Utils::toBytes("");
        if (Utils::equalString($signatureAlgorithm, "ACS4-HMAC-SHA256")) {
            $signature = SignatureUtil::HmacSHA256SignByBytes($stringToSign, $signingkey);
        } else if (Utils::equalString($signatureAlgorithm, "ACS4-HMAC-SM3")) {
            $signature = SignatureUtil::HmacSM3SignByBytes($stringToSign, $signingkey);
        }
        return EncodeUtil::hexEncode($signature);
    }
}
