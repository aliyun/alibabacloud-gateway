// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListResourcePoolsResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListResourcePoolsResponseBody body;

    public static ListResourcePoolsResponse build(java.util.Map<String, ?> map) throws Exception {
        ListResourcePoolsResponse self = new ListResourcePoolsResponse();
        return TeaModel.build(map, self);
    }

    public ListResourcePoolsResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListResourcePoolsResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListResourcePoolsResponse setBody(ListResourcePoolsResponseBody body) {
        this.body = body;
        return this;
    }
    public ListResourcePoolsResponseBody getBody() {
        return this.body;
    }

}
