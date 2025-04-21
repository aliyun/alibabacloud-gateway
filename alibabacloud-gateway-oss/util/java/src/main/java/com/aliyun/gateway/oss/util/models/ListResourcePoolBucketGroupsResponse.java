// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListResourcePoolBucketGroupsResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListResourcePoolBucketGroupsResponseBody body;

    public static ListResourcePoolBucketGroupsResponse build(java.util.Map<String, ?> map) throws Exception {
        ListResourcePoolBucketGroupsResponse self = new ListResourcePoolBucketGroupsResponse();
        return TeaModel.build(map, self);
    }

    public ListResourcePoolBucketGroupsResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListResourcePoolBucketGroupsResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListResourcePoolBucketGroupsResponse setBody(ListResourcePoolBucketGroupsResponseBody body) {
        this.body = body;
        return this;
    }
    public ListResourcePoolBucketGroupsResponseBody getBody() {
        return this.body;
    }

}
