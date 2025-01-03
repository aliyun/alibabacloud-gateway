// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutObjectLinkResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public PutObjectLinkResponseBody body;

    public static PutObjectLinkResponse build(java.util.Map<String, ?> map) throws Exception {
        PutObjectLinkResponse self = new PutObjectLinkResponse();
        return TeaModel.build(map, self);
    }

    public PutObjectLinkResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public PutObjectLinkResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public PutObjectLinkResponse setBody(PutObjectLinkResponseBody body) {
        this.body = body;
        return this;
    }
    public PutObjectLinkResponseBody getBody() {
        return this.body;
    }

}
