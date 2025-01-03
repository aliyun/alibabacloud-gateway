// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetDataLakeCachePrefetchJobResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetDataLakeCachePrefetchJobResponseBody body;

    public static GetDataLakeCachePrefetchJobResponse build(java.util.Map<String, ?> map) throws Exception {
        GetDataLakeCachePrefetchJobResponse self = new GetDataLakeCachePrefetchJobResponse();
        return TeaModel.build(map, self);
    }

    public GetDataLakeCachePrefetchJobResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetDataLakeCachePrefetchJobResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetDataLakeCachePrefetchJobResponse setBody(GetDataLakeCachePrefetchJobResponseBody body) {
        this.body = body;
        return this;
    }
    public GetDataLakeCachePrefetchJobResponseBody getBody() {
        return this.body;
    }

}
