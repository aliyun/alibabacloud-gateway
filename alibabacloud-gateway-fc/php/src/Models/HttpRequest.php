<?php

// This file is auto-generated, don't edit it. Thanks.
namespace Darabonba\GatewayFc\Models;

use AlibabaCloud\Tea\Model;

class HttpRequest extends Model
{
    public function validate()
    {
        Model::validateRequired('method', $this->method, true);
        Model::validateRequired('path', $this->path, true);
    }
    public function toMap()
    {
        $res = [];
        if (null !== $this->method) {
            $res['method'] = $this->method;
        }
        if (null !== $this->path) {
            $res['path'] = $this->path;
        }
        if (null !== $this->headers) {
            $res['headers'] = $this->headers;
        }
        if (null !== $this->body) {
            $res['body'] = $this->body;
        }
        if (null !== $this->reqBodyType) {
            $res['reqBodyType'] = $this->reqBodyType;
        }
        return $res;
    }
    /**
     * @param array $map
     * @return HttpRequest
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['method'])) {
            $model->method = $map['method'];
        }
        if (isset($map['path'])) {
            $model->path = $map['path'];
        }
        if (isset($map['headers'])) {
            $model->headers = $map['headers'];
        }
        if (isset($map['body'])) {
            $model->body = $map['body'];
        }
        if (isset($map['reqBodyType'])) {
            $model->reqBodyType = $map['reqBodyType'];
        }
        return $model;
    }
    /**
     * @var string
     */
    public $method;

    /**
     * @var string
     */
    public $path;

    public $headers;

    public $body;

    public $reqBodyType;
}
