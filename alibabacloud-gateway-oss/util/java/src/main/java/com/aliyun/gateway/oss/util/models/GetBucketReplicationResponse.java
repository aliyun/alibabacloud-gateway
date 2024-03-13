// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketReplicationResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketReplicationResponseBody body;

    public static GetBucketReplicationResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketReplicationResponse self = new GetBucketReplicationResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketReplicationResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketReplicationResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketReplicationResponse setBody(GetBucketReplicationResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketReplicationResponseBody getBody() {
        return this.body;
    }

}
