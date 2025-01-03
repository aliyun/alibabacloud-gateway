// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DeleteBucketRequesterQoSInfoResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    public static DeleteBucketRequesterQoSInfoResponse build(java.util.Map<String, ?> map) throws Exception {
        DeleteBucketRequesterQoSInfoResponse self = new DeleteBucketRequesterQoSInfoResponse();
        return TeaModel.build(map, self);
    }

    public DeleteBucketRequesterQoSInfoResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public DeleteBucketRequesterQoSInfoResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

}
