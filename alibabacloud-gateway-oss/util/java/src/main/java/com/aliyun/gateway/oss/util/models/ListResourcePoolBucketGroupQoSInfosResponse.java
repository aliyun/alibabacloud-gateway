// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListResourcePoolBucketGroupQoSInfosResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListResourcePoolBucketGroupQoSInfosResponseBody body;

    public static ListResourcePoolBucketGroupQoSInfosResponse build(java.util.Map<String, ?> map) throws Exception {
        ListResourcePoolBucketGroupQoSInfosResponse self = new ListResourcePoolBucketGroupQoSInfosResponse();
        return TeaModel.build(map, self);
    }

    public ListResourcePoolBucketGroupQoSInfosResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListResourcePoolBucketGroupQoSInfosResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListResourcePoolBucketGroupQoSInfosResponse setBody(ListResourcePoolBucketGroupQoSInfosResponseBody body) {
        this.body = body;
        return this;
    }
    public ListResourcePoolBucketGroupQoSInfosResponseBody getBody() {
        return this.body;
    }

}
