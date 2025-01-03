// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketHashResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketHashResponseBody body;

    public static GetBucketHashResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketHashResponse self = new GetBucketHashResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketHashResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketHashResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketHashResponse setBody(GetBucketHashResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketHashResponseBody getBody() {
        return this.body;
    }

}
