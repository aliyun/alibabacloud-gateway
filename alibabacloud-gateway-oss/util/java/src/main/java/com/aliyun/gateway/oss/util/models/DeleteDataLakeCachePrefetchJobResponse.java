// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DeleteDataLakeCachePrefetchJobResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    public static DeleteDataLakeCachePrefetchJobResponse build(java.util.Map<String, ?> map) throws Exception {
        DeleteDataLakeCachePrefetchJobResponse self = new DeleteDataLakeCachePrefetchJobResponse();
        return TeaModel.build(map, self);
    }

    public DeleteDataLakeCachePrefetchJobResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public DeleteDataLakeCachePrefetchJobResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

}
