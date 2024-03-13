// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class InitBucketAntiDDosInfoResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    public static InitBucketAntiDDosInfoResponse build(java.util.Map<String, ?> map) throws Exception {
        InitBucketAntiDDosInfoResponse self = new InitBucketAntiDDosInfoResponse();
        return TeaModel.build(map, self);
    }

    public InitBucketAntiDDosInfoResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public InitBucketAntiDDosInfoResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

}
