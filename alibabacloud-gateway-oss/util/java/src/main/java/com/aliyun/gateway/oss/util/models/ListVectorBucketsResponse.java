// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListVectorBucketsResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListVectorBucketsResponseBody body;

    public static ListVectorBucketsResponse build(java.util.Map<String, ?> map) throws Exception {
        ListVectorBucketsResponse self = new ListVectorBucketsResponse();
        return TeaModel.build(map, self);
    }

    public ListVectorBucketsResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListVectorBucketsResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListVectorBucketsResponse setBody(ListVectorBucketsResponseBody body) {
        this.body = body;
        return this;
    }
    public ListVectorBucketsResponseBody getBody() {
        return this.body;
    }

}
