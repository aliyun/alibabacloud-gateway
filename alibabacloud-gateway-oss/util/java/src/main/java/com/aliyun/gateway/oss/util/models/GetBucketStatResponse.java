// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class GetBucketStatResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketStatResponseBody body;

    public static GetBucketStatResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketStatResponse self = new GetBucketStatResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketStatResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketStatResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketStatResponse setBody(GetBucketStatResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketStatResponseBody getBody() {
        return this.body;
    }

}
