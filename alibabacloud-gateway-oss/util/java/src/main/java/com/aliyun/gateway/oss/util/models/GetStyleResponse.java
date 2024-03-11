// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class GetStyleResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetStyleResponseBody body;

    public static GetStyleResponse build(java.util.Map<String, ?> map) throws Exception {
        GetStyleResponse self = new GetStyleResponse();
        return TeaModel.build(map, self);
    }

    public GetStyleResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetStyleResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetStyleResponse setBody(GetStyleResponseBody body) {
        this.body = body;
        return this;
    }
    public GetStyleResponseBody getBody() {
        return this.body;
    }

}
