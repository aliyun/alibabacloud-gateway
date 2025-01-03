// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutDataLakeCachePrefetchJobResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public PutDataLakeCachePrefetchJobResponseBody body;

    public static PutDataLakeCachePrefetchJobResponse build(java.util.Map<String, ?> map) throws Exception {
        PutDataLakeCachePrefetchJobResponse self = new PutDataLakeCachePrefetchJobResponse();
        return TeaModel.build(map, self);
    }

    public PutDataLakeCachePrefetchJobResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public PutDataLakeCachePrefetchJobResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public PutDataLakeCachePrefetchJobResponse setBody(PutDataLakeCachePrefetchJobResponseBody body) {
        this.body = body;
        return this;
    }
    public PutDataLakeCachePrefetchJobResponseBody getBody() {
        return this.body;
    }

}
