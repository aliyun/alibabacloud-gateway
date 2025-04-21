// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutResourcePoolBucketGroupQoSInfoResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    public static PutResourcePoolBucketGroupQoSInfoResponse build(java.util.Map<String, ?> map) throws Exception {
        PutResourcePoolBucketGroupQoSInfoResponse self = new PutResourcePoolBucketGroupQoSInfoResponse();
        return TeaModel.build(map, self);
    }

    public PutResourcePoolBucketGroupQoSInfoResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public PutResourcePoolBucketGroupQoSInfoResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

}
