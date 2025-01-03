// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListBucketRequesterQoSInfosResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListBucketRequesterQoSInfosResponseBody body;

    public static ListBucketRequesterQoSInfosResponse build(java.util.Map<String, ?> map) throws Exception {
        ListBucketRequesterQoSInfosResponse self = new ListBucketRequesterQoSInfosResponse();
        return TeaModel.build(map, self);
    }

    public ListBucketRequesterQoSInfosResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListBucketRequesterQoSInfosResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListBucketRequesterQoSInfosResponse setBody(ListBucketRequesterQoSInfosResponseBody body) {
        this.body = body;
        return this;
    }
    public ListBucketRequesterQoSInfosResponseBody getBody() {
        return this.body;
    }

}
