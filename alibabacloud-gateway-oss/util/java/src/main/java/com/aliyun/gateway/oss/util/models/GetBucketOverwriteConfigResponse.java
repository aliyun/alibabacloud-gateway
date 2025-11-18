// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketOverwriteConfigResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketOverwriteConfigResponseBody body;

    public static GetBucketOverwriteConfigResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketOverwriteConfigResponse self = new GetBucketOverwriteConfigResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketOverwriteConfigResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketOverwriteConfigResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketOverwriteConfigResponse setBody(GetBucketOverwriteConfigResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketOverwriteConfigResponseBody getBody() {
        return this.body;
    }

}
