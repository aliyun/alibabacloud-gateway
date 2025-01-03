// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketQoSInfoResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketQoSInfoResponseBody body;

    public static GetBucketQoSInfoResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketQoSInfoResponse self = new GetBucketQoSInfoResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketQoSInfoResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketQoSInfoResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketQoSInfoResponse setBody(GetBucketQoSInfoResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketQoSInfoResponseBody getBody() {
        return this.body;
    }

}
