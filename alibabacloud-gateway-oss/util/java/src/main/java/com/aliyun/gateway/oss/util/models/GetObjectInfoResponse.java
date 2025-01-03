// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetObjectInfoResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetObjectInfoResponseBody body;

    public static GetObjectInfoResponse build(java.util.Map<String, ?> map) throws Exception {
        GetObjectInfoResponse self = new GetObjectInfoResponse();
        return TeaModel.build(map, self);
    }

    public GetObjectInfoResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetObjectInfoResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetObjectInfoResponse setBody(GetObjectInfoResponseBody body) {
        this.body = body;
        return this;
    }
    public GetObjectInfoResponseBody getBody() {
        return this.body;
    }

}
