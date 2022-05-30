<?php

// This file is auto-generated, don't edit it. Thanks.
namespace Darabonba\GatewayFc\Util;

use AlibabaCloud\Credentials\Credential;
use GuzzleHttp\Psr7\Response;
use \Exception;
use GuzzleHttp\Psr7\Request;

class Client {

    /**
     * @param Credential $credential
     * @param string $url
     * @param string $method
     * @param int[] $body
     * @param array $headers
     * @return Response
     */
    public static function InvokeHTTPTrigger($credential, $url, $method, $body, $headers){
        throw new Exception('Un-implemented');
    }

    /**
     * @param string $url
     * @param string $method
     * @param int[] $body
     * @param array $headers
     * @return Response
     */
    public static function InvokeAnonymousHTTPTrigger($url, $method, $body, $headers){
        throw new Exception('Un-implemented');
    }

    /**
     * @param Credential $credential
     * @param Request $req
     * @return Response
     */
    public static function SendHTTPRequestWithAuthorization($credential, $req){
        throw new Exception('Un-implemented');
    }

    /**
     * @param Request $req
     * @return Response
     */
    public static function SendHTTPRequest($req){
        throw new Exception('Un-implemented');
    }

    /**
     * @param Credential $credential
     * @param Request $req
     * @return Request
     */
    public static function SignRequest($credential, $req){
        throw new Exception('Un-implemented');
    }

    /**
     * @param Credential $credential
     * @param Request $req
     * @param string $contentMD5
     * @return Request
     */
    public static function SignRequestWithContentMD5($credential, $req, $contentMD5){
        throw new Exception('Un-implemented');
    }

    /**
     * @param string $url
     * @param string $method
     * @param int[] $body
     * @param array $headers
     * @return Request
     */
    public static function BuildHTTPRequest($url, $method, $body, $headers){
        throw new Exception('Un-implemented');
    }
}
