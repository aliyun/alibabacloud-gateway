// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class GetBucketInventoryResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketInventoryResponseBody body;

    public static GetBucketInventoryResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketInventoryResponse self = new GetBucketInventoryResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketInventoryResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketInventoryResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketInventoryResponse setBody(GetBucketInventoryResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketInventoryResponseBody getBody() {
        return this.body;
    }

}
