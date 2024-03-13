// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketLocationResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketLocationResponseBody body;

    public static GetBucketLocationResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketLocationResponse self = new GetBucketLocationResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketLocationResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketLocationResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketLocationResponse setBody(GetBucketLocationResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketLocationResponseBody getBody() {
        return this.body;
    }

}
