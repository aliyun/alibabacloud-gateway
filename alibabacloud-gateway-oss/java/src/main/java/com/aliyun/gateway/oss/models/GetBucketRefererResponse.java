// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class GetBucketRefererResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketRefererResponseBody body;

    public static GetBucketRefererResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketRefererResponse self = new GetBucketRefererResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketRefererResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketRefererResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketRefererResponse setBody(GetBucketRefererResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketRefererResponseBody getBody() {
        return this.body;
    }

}
