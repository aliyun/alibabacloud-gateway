// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class GetAccessPointConfigForObjectProcessResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetAccessPointConfigForObjectProcessResponseBody body;

    public static GetAccessPointConfigForObjectProcessResponse build(java.util.Map<String, ?> map) throws Exception {
        GetAccessPointConfigForObjectProcessResponse self = new GetAccessPointConfigForObjectProcessResponse();
        return TeaModel.build(map, self);
    }

    public GetAccessPointConfigForObjectProcessResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetAccessPointConfigForObjectProcessResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetAccessPointConfigForObjectProcessResponse setBody(GetAccessPointConfigForObjectProcessResponseBody body) {
        this.body = body;
        return this;
    }
    public GetAccessPointConfigForObjectProcessResponseBody getBody() {
        return this.body;
    }

}
