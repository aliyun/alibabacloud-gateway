// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class QueryVectorsResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public QueryVectorsResponseBody body;

    public static QueryVectorsResponse build(java.util.Map<String, ?> map) throws Exception {
        QueryVectorsResponse self = new QueryVectorsResponse();
        return TeaModel.build(map, self);
    }

    public QueryVectorsResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public QueryVectorsResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public QueryVectorsResponse setBody(QueryVectorsResponseBody body) {
        this.body = body;
        return this;
    }
    public QueryVectorsResponseBody getBody() {
        return this.body;
    }

}
