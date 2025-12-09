// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetVectorIndexResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetVectorIndexResponseBody body;

    public static GetVectorIndexResponse build(java.util.Map<String, ?> map) throws Exception {
        GetVectorIndexResponse self = new GetVectorIndexResponse();
        return TeaModel.build(map, self);
    }

    public GetVectorIndexResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetVectorIndexResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetVectorIndexResponse setBody(GetVectorIndexResponseBody body) {
        this.body = body;
        return this;
    }
    public GetVectorIndexResponseBody getBody() {
        return this.body;
    }

}
