// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketDataAcceleratorResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketDataAcceleratorResponseBody body;

    public static GetBucketDataAcceleratorResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketDataAcceleratorResponse self = new GetBucketDataAcceleratorResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketDataAcceleratorResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketDataAcceleratorResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketDataAcceleratorResponse setBody(GetBucketDataAcceleratorResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketDataAcceleratorResponseBody getBody() {
        return this.body;
    }

}
