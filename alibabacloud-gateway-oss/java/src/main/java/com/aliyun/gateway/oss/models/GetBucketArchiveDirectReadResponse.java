// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class GetBucketArchiveDirectReadResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketArchiveDirectReadResponseBody body;

    public static GetBucketArchiveDirectReadResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketArchiveDirectReadResponse self = new GetBucketArchiveDirectReadResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketArchiveDirectReadResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketArchiveDirectReadResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketArchiveDirectReadResponse setBody(GetBucketArchiveDirectReadResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketArchiveDirectReadResponseBody getBody() {
        return this.body;
    }

}
