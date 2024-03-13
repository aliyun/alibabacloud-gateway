// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetAccessPointForObjectProcessResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetAccessPointForObjectProcessResponseBody body;

    public static GetAccessPointForObjectProcessResponse build(java.util.Map<String, ?> map) throws Exception {
        GetAccessPointForObjectProcessResponse self = new GetAccessPointForObjectProcessResponse();
        return TeaModel.build(map, self);
    }

    public GetAccessPointForObjectProcessResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetAccessPointForObjectProcessResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetAccessPointForObjectProcessResponse setBody(GetAccessPointForObjectProcessResponseBody body) {
        this.body = body;
        return this;
    }
    public GetAccessPointForObjectProcessResponseBody getBody() {
        return this.body;
    }

}
