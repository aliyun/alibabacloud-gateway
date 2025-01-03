// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetVirtualBucketResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetVirtualBucketResponseBody body;

    public static GetVirtualBucketResponse build(java.util.Map<String, ?> map) throws Exception {
        GetVirtualBucketResponse self = new GetVirtualBucketResponse();
        return TeaModel.build(map, self);
    }

    public GetVirtualBucketResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetVirtualBucketResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetVirtualBucketResponse setBody(GetVirtualBucketResponseBody body) {
        this.body = body;
        return this;
    }
    public GetVirtualBucketResponseBody getBody() {
        return this.body;
    }

}
