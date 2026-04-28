// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.oss20190517.models;

import com.aliyun.tea.*;

public class PutBucketObjectWormConfigurationResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    public static PutBucketObjectWormConfigurationResponse build(java.util.Map<String, ?> map) throws Exception {
        PutBucketObjectWormConfigurationResponse self = new PutBucketObjectWormConfigurationResponse();
        return TeaModel.build(map, self);
    }

    public PutBucketObjectWormConfigurationResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public PutBucketObjectWormConfigurationResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

}
