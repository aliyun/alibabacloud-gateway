// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListVirtualBucketResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListVirtualBucketResponseBody body;

    public static ListVirtualBucketResponse build(java.util.Map<String, ?> map) throws Exception {
        ListVirtualBucketResponse self = new ListVirtualBucketResponse();
        return TeaModel.build(map, self);
    }

    public ListVirtualBucketResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListVirtualBucketResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListVirtualBucketResponse setBody(ListVirtualBucketResponseBody body) {
        this.body = body;
        return this;
    }
    public ListVirtualBucketResponseBody getBody() {
        return this.body;
    }

}
