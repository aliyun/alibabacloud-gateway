// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class ListAccessPointsResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListAccessPointsResponseBody body;

    public static ListAccessPointsResponse build(java.util.Map<String, ?> map) throws Exception {
        ListAccessPointsResponse self = new ListAccessPointsResponse();
        return TeaModel.build(map, self);
    }

    public ListAccessPointsResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListAccessPointsResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListAccessPointsResponse setBody(ListAccessPointsResponseBody body) {
        this.body = body;
        return this;
    }
    public ListAccessPointsResponseBody getBody() {
        return this.body;
    }

}
