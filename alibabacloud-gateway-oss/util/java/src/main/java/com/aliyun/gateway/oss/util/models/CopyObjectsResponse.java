// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CopyObjectsResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public CopyObjectsResponseBody body;

    public static CopyObjectsResponse build(java.util.Map<String, ?> map) throws Exception {
        CopyObjectsResponse self = new CopyObjectsResponse();
        return TeaModel.build(map, self);
    }

    public CopyObjectsResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public CopyObjectsResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public CopyObjectsResponse setBody(CopyObjectsResponseBody body) {
        this.body = body;
        return this;
    }
    public CopyObjectsResponseBody getBody() {
        return this.body;
    }

}
