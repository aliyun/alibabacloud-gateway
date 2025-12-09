// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetVectorsResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetVectorsResponseBody body;

    public static GetVectorsResponse build(java.util.Map<String, ?> map) throws Exception {
        GetVectorsResponse self = new GetVectorsResponse();
        return TeaModel.build(map, self);
    }

    public GetVectorsResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetVectorsResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetVectorsResponse setBody(GetVectorsResponseBody body) {
        this.body = body;
        return this;
    }
    public GetVectorsResponseBody getBody() {
        return this.body;
    }

}
