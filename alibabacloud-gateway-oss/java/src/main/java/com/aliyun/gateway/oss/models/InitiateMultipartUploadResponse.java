// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class InitiateMultipartUploadResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public InitiateMultipartUploadResponseBody body;

    public static InitiateMultipartUploadResponse build(java.util.Map<String, ?> map) throws Exception {
        InitiateMultipartUploadResponse self = new InitiateMultipartUploadResponse();
        return TeaModel.build(map, self);
    }

    public InitiateMultipartUploadResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public InitiateMultipartUploadResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public InitiateMultipartUploadResponse setBody(InitiateMultipartUploadResponseBody body) {
        this.body = body;
        return this;
    }
    public InitiateMultipartUploadResponseBody getBody() {
        return this.body;
    }

}
