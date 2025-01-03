// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListCacheResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListCacheResponseBody body;

    public static ListCacheResponse build(java.util.Map<String, ?> map) throws Exception {
        ListCacheResponse self = new ListCacheResponse();
        return TeaModel.build(map, self);
    }

    public ListCacheResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListCacheResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListCacheResponse setBody(ListCacheResponseBody body) {
        this.body = body;
        return this;
    }
    public ListCacheResponseBody getBody() {
        return this.body;
    }

}
