// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetObjectLinkResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetObjectLinkResponseBody body;

    public static GetObjectLinkResponse build(java.util.Map<String, ?> map) throws Exception {
        GetObjectLinkResponse self = new GetObjectLinkResponse();
        return TeaModel.build(map, self);
    }

    public GetObjectLinkResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetObjectLinkResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetObjectLinkResponse setBody(GetObjectLinkResponseBody body) {
        this.body = body;
        return this;
    }
    public GetObjectLinkResponseBody getBody() {
        return this.body;
    }

}
