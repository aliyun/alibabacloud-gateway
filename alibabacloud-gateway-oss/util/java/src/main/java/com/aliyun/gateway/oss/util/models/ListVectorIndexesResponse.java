// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListVectorIndexesResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListVectorIndexesResponseBody body;

    public static ListVectorIndexesResponse build(java.util.Map<String, ?> map) throws Exception {
        ListVectorIndexesResponse self = new ListVectorIndexesResponse();
        return TeaModel.build(map, self);
    }

    public ListVectorIndexesResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListVectorIndexesResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListVectorIndexesResponse setBody(ListVectorIndexesResponseBody body) {
        this.body = body;
        return this;
    }
    public ListVectorIndexesResponseBody getBody() {
        return this.body;
    }

}
