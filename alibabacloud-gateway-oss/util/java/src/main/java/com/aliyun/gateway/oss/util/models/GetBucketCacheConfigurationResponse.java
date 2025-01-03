// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketCacheConfigurationResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketCacheConfigurationResponseBody body;

    public static GetBucketCacheConfigurationResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketCacheConfigurationResponse self = new GetBucketCacheConfigurationResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketCacheConfigurationResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketCacheConfigurationResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketCacheConfigurationResponse setBody(GetBucketCacheConfigurationResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketCacheConfigurationResponseBody getBody() {
        return this.body;
    }

}
