// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListDataLakeCachePrefetchJobResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListDataLakeCachePrefetchJobResponseBody body;

    public static ListDataLakeCachePrefetchJobResponse build(java.util.Map<String, ?> map) throws Exception {
        ListDataLakeCachePrefetchJobResponse self = new ListDataLakeCachePrefetchJobResponse();
        return TeaModel.build(map, self);
    }

    public ListDataLakeCachePrefetchJobResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListDataLakeCachePrefetchJobResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListDataLakeCachePrefetchJobResponse setBody(ListDataLakeCachePrefetchJobResponseBody body) {
        this.body = body;
        return this;
    }
    public ListDataLakeCachePrefetchJobResponseBody getBody() {
        return this.body;
    }

}
