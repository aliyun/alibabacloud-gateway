// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketCommonHeaderResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    public static PutBucketCommonHeaderResponse build(java.util.Map<String, ?> map) throws Exception {
        PutBucketCommonHeaderResponse self = new PutBucketCommonHeaderResponse();
        return TeaModel.build(map, self);
    }

    public PutBucketCommonHeaderResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public PutBucketCommonHeaderResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

}
