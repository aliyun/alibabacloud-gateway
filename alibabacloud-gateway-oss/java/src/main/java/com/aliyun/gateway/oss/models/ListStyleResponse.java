// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class ListStyleResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListStyleResponseBody body;

    public static ListStyleResponse build(java.util.Map<String, ?> map) throws Exception {
        ListStyleResponse self = new ListStyleResponse();
        return TeaModel.build(map, self);
    }

    public ListStyleResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListStyleResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListStyleResponse setBody(ListStyleResponseBody body) {
        this.body = body;
        return this;
    }
    public ListStyleResponseBody getBody() {
        return this.body;
    }

}
