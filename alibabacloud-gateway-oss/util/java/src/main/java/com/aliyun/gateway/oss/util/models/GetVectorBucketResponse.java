// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetVectorBucketResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetVectorBucketResponseBody body;

    public static GetVectorBucketResponse build(java.util.Map<String, ?> map) throws Exception {
        GetVectorBucketResponse self = new GetVectorBucketResponse();
        return TeaModel.build(map, self);
    }

    public GetVectorBucketResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetVectorBucketResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetVectorBucketResponse setBody(GetVectorBucketResponseBody body) {
        this.body = body;
        return this;
    }
    public GetVectorBucketResponseBody getBody() {
        return this.body;
    }

}
