// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetResourcePoolInfoResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetResourcePoolInfoResponseBody body;

    public static GetResourcePoolInfoResponse build(java.util.Map<String, ?> map) throws Exception {
        GetResourcePoolInfoResponse self = new GetResourcePoolInfoResponse();
        return TeaModel.build(map, self);
    }

    public GetResourcePoolInfoResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetResourcePoolInfoResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetResourcePoolInfoResponse setBody(GetResourcePoolInfoResponseBody body) {
        this.body = body;
        return this;
    }
    public GetResourcePoolInfoResponseBody getBody() {
        return this.body;
    }

}
