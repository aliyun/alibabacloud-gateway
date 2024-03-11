// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class UpdateBucketAntiDDosInfoResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    public static UpdateBucketAntiDDosInfoResponse build(java.util.Map<String, ?> map) throws Exception {
        UpdateBucketAntiDDosInfoResponse self = new UpdateBucketAntiDDosInfoResponse();
        return TeaModel.build(map, self);
    }

    public UpdateBucketAntiDDosInfoResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public UpdateBucketAntiDDosInfoResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

}
