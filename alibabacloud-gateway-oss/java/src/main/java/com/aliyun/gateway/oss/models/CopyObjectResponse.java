// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class CopyObjectResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public CopyObjectResponseBody body;

    public static CopyObjectResponse build(java.util.Map<String, ?> map) throws Exception {
        CopyObjectResponse self = new CopyObjectResponse();
        return TeaModel.build(map, self);
    }

    public CopyObjectResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public CopyObjectResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public CopyObjectResponse setBody(CopyObjectResponseBody body) {
        this.body = body;
        return this;
    }
    public CopyObjectResponseBody getBody() {
        return this.body;
    }

}
