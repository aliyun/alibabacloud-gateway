// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketRequestPaymentResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketRequestPaymentResponseBody body;

    public static GetBucketRequestPaymentResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketRequestPaymentResponse self = new GetBucketRequestPaymentResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketRequestPaymentResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketRequestPaymentResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketRequestPaymentResponse setBody(GetBucketRequestPaymentResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketRequestPaymentResponseBody getBody() {
        return this.body;
    }

}
