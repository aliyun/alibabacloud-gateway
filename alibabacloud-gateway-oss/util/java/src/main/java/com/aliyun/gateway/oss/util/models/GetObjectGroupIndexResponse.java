// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetObjectGroupIndexResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetObjectGroupIndexResponseBody body;

    public static GetObjectGroupIndexResponse build(java.util.Map<String, ?> map) throws Exception {
        GetObjectGroupIndexResponse self = new GetObjectGroupIndexResponse();
        return TeaModel.build(map, self);
    }

    public GetObjectGroupIndexResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetObjectGroupIndexResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetObjectGroupIndexResponse setBody(GetObjectGroupIndexResponseBody body) {
        this.body = body;
        return this;
    }
    public GetObjectGroupIndexResponseBody getBody() {
        return this.body;
    }

}
