// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class StartPartUploadResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public StartPartUploadResponseBody body;

    public static StartPartUploadResponse build(java.util.Map<String, ?> map) throws Exception {
        StartPartUploadResponse self = new StartPartUploadResponse();
        return TeaModel.build(map, self);
    }

    public StartPartUploadResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public StartPartUploadResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public StartPartUploadResponse setBody(StartPartUploadResponseBody body) {
        this.body = body;
        return this;
    }
    public StartPartUploadResponseBody getBody() {
        return this.body;
    }

}
