// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListDataLakeCachePrefetchJobHistoryResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListDataLakeCachePrefetchJobHistoryResponseBody body;

    public static ListDataLakeCachePrefetchJobHistoryResponse build(java.util.Map<String, ?> map) throws Exception {
        ListDataLakeCachePrefetchJobHistoryResponse self = new ListDataLakeCachePrefetchJobHistoryResponse();
        return TeaModel.build(map, self);
    }

    public ListDataLakeCachePrefetchJobHistoryResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListDataLakeCachePrefetchJobHistoryResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListDataLakeCachePrefetchJobHistoryResponse setBody(ListDataLakeCachePrefetchJobHistoryResponseBody body) {
        this.body = body;
        return this;
    }
    public ListDataLakeCachePrefetchJobHistoryResponseBody getBody() {
        return this.body;
    }

}
