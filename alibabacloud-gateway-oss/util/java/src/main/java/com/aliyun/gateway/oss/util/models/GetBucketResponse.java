// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class GetBucketResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketResponseBody body;

    public static GetBucketResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketResponse self = new GetBucketResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketResponse setBody(GetBucketResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketResponseBody getBody() {
        return this.body;
    }

}
