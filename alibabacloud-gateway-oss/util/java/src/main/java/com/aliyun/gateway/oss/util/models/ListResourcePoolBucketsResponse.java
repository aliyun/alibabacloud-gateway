// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListResourcePoolBucketsResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListResourcePoolBucketsResponseBody body;

    public static ListResourcePoolBucketsResponse build(java.util.Map<String, ?> map) throws Exception {
        ListResourcePoolBucketsResponse self = new ListResourcePoolBucketsResponse();
        return TeaModel.build(map, self);
    }

    public ListResourcePoolBucketsResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListResourcePoolBucketsResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListResourcePoolBucketsResponse setBody(ListResourcePoolBucketsResponseBody body) {
        this.body = body;
        return this;
    }
    public ListResourcePoolBucketsResponseBody getBody() {
        return this.body;
    }

}
