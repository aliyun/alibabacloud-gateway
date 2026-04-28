// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketObjectWormConfigurationResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketObjectWormConfigurationResponseBody body;

    public static GetBucketObjectWormConfigurationResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketObjectWormConfigurationResponse self = new GetBucketObjectWormConfigurationResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketObjectWormConfigurationResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketObjectWormConfigurationResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketObjectWormConfigurationResponse setBody(GetBucketObjectWormConfigurationResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketObjectWormConfigurationResponseBody getBody() {
        return this.body;
    }

}
