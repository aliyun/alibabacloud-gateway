// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetCacheResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetCacheResponseBody body;

    public static GetCacheResponse build(java.util.Map<String, ?> map) throws Exception {
        GetCacheResponse self = new GetCacheResponse();
        return TeaModel.build(map, self);
    }

    public GetCacheResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetCacheResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetCacheResponse setBody(GetCacheResponseBody body) {
        this.body = body;
        return this;
    }
    public GetCacheResponseBody getBody() {
        return this.body;
    }

}
