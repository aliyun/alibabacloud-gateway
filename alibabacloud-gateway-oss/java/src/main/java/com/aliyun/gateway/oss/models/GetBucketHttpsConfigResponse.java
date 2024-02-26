// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class GetBucketHttpsConfigResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketHttpsConfigResponseBody body;

    public static GetBucketHttpsConfigResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketHttpsConfigResponse self = new GetBucketHttpsConfigResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketHttpsConfigResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketHttpsConfigResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketHttpsConfigResponse setBody(GetBucketHttpsConfigResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketHttpsConfigResponseBody getBody() {
        return this.body;
    }

}
