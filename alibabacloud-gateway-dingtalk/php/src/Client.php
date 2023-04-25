<?php

// This file is auto-generated, don't edit it. Thanks.
namespace Darabonba\GatewayDingTalk;

use Darabonba\GatewaySpi\Client as DarabonbaGatewaySpiClient;
use AlibabaCloud\Tea\Tea;
use AlibabaCloud\Tea\Utils\Utils;
use AlibabaCloud\Tea\Exception\TeaError;

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
            "host" => $config->endpoint,
            "user-agent" => $request->userAgent,
            "accept" => "application/json"
        ], $request->headers);
        if (!Utils::isUnset($request->body)) {
            $jsonObj = Utils::toJSONString($request->body);
            $request->stream = $jsonObj;
            $request->headers["content-type"] = "application/json; charset=utf-8";
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
                "message" => "code: " . (string) ($response->statusCode) . ", " . (string) ($this->defaultAny(@$err["Message"], @$err["message"])) . " request id: " . (string) ($this->defaultAny(@$err["RequestId"], @$err["requestid"])) . "",
                "data" => $err,
                "description" => "" . (string) ($this->defaultAny(@$err["Description"], @$err["description"])) . "",
                "accessDeniedDetail" => $this->defaultAny(@$err["AccessDeniedDetail"], @$err["accessdenieddetail"])
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
}
