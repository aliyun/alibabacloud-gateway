// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DeleteResourcePoolBucketGroupQoSInfoResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    public static DeleteResourcePoolBucketGroupQoSInfoResponse build(java.util.Map<String, ?> map) throws Exception {
        DeleteResourcePoolBucketGroupQoSInfoResponse self = new DeleteResourcePoolBucketGroupQoSInfoResponse();
        return TeaModel.build(map, self);
    }

    public DeleteResourcePoolBucketGroupQoSInfoResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public DeleteResourcePoolBucketGroupQoSInfoResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

}
