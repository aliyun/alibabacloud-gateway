// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListVectorsResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListVectorsResponseBody body;

    public static ListVectorsResponse build(java.util.Map<String, ?> map) throws Exception {
        ListVectorsResponse self = new ListVectorsResponse();
        return TeaModel.build(map, self);
    }

    public ListVectorsResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListVectorsResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListVectorsResponse setBody(ListVectorsResponseBody body) {
        this.body = body;
        return this;
    }
    public ListVectorsResponseBody getBody() {
        return this.body;
    }

}
