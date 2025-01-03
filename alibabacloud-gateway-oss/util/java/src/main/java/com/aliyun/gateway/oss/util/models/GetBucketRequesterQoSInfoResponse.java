// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketRequesterQoSInfoResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketRequesterQoSInfoResponseBody body;

    public static GetBucketRequesterQoSInfoResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketRequesterQoSInfoResponse self = new GetBucketRequesterQoSInfoResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketRequesterQoSInfoResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketRequesterQoSInfoResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketRequesterQoSInfoResponse setBody(GetBucketRequesterQoSInfoResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketRequesterQoSInfoResponseBody getBody() {
        return this.body;
    }

}
