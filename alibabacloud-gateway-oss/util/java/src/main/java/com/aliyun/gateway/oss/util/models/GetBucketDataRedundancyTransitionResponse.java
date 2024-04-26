// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketDataRedundancyTransitionResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketDataRedundancyTransitionResponseBody body;

    public static GetBucketDataRedundancyTransitionResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketDataRedundancyTransitionResponse self = new GetBucketDataRedundancyTransitionResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketDataRedundancyTransitionResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketDataRedundancyTransitionResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketDataRedundancyTransitionResponse setBody(GetBucketDataRedundancyTransitionResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketDataRedundancyTransitionResponseBody getBody() {
        return this.body;
    }

}
