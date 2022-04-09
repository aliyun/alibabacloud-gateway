// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.fc.models;

import com.aliyun.tea.*;

public class HttpRequest extends TeaModel {
    @NameInMap("method")
    @Validation(required = true)
    public String method;

    @NameInMap("path")
    @Validation(required = true)
    public String path;

    @NameInMap("headers")
    public java.util.Map<String, Object> headers;

    @NameInMap("body")
    public byte[] body;

    @NameInMap("reqBodyType")
    public String reqBodyType;

    public static HttpRequest build(java.util.Map<String, ?> map) throws Exception {
        HttpRequest self = new HttpRequest();
        return TeaModel.build(map, self);
    }

    public HttpRequest setMethod(String method) {
        this.method = method;
        return this;
    }
    public String getMethod() {
        return this.method;
    }

    public HttpRequest setPath(String path) {
        this.path = path;
        return this;
    }
    public String getPath() {
        return this.path;
    }

    public HttpRequest setHeaders(java.util.Map<String, Object> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, Object> getHeaders() {
        return this.headers;
    }

    public HttpRequest setBody(byte[] body) {
        this.body = body;
        return this;
    }
    public byte[] getBody() {
        return this.body;
    }

    public HttpRequest setReqBodyType(String reqBodyType) {
        this.reqBodyType = reqBodyType;
        return this;
    }
    public String getReqBodyType() {
        return this.reqBodyType;
    }

}
