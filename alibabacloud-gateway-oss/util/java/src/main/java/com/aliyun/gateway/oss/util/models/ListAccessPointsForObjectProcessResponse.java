// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListAccessPointsForObjectProcessResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListAccessPointsForObjectProcessResponseBody body;

    public static ListAccessPointsForObjectProcessResponse build(java.util.Map<String, ?> map) throws Exception {
        ListAccessPointsForObjectProcessResponse self = new ListAccessPointsForObjectProcessResponse();
        return TeaModel.build(map, self);
    }

    public ListAccessPointsForObjectProcessResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListAccessPointsForObjectProcessResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListAccessPointsForObjectProcessResponse setBody(ListAccessPointsForObjectProcessResponseBody body) {
        this.body = body;
        return this;
    }
    public ListAccessPointsForObjectProcessResponseBody getBody() {
        return this.body;
    }

}
