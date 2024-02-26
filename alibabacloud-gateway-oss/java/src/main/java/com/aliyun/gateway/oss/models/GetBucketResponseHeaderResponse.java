// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class GetBucketResponseHeaderResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketResponseHeaderResponseBody body;

    public static GetBucketResponseHeaderResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketResponseHeaderResponse self = new GetBucketResponseHeaderResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketResponseHeaderResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketResponseHeaderResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketResponseHeaderResponse setBody(GetBucketResponseHeaderResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketResponseHeaderResponseBody getBody() {
        return this.body;
    }

}
