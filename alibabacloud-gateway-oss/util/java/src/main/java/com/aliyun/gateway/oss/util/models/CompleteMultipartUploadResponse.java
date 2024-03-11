// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class CompleteMultipartUploadResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public CompleteMultipartUploadResponseBody body;

    public static CompleteMultipartUploadResponse build(java.util.Map<String, ?> map) throws Exception {
        CompleteMultipartUploadResponse self = new CompleteMultipartUploadResponse();
        return TeaModel.build(map, self);
    }

    public CompleteMultipartUploadResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public CompleteMultipartUploadResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public CompleteMultipartUploadResponse setBody(CompleteMultipartUploadResponseBody body) {
        this.body = body;
        return this;
    }
    public CompleteMultipartUploadResponseBody getBody() {
        return this.body;
    }

}
