// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetResourcePoolBucketGroupQoSInfoResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetResourcePoolBucketGroupQoSInfoResponseBody body;

    public static GetResourcePoolBucketGroupQoSInfoResponse build(java.util.Map<String, ?> map) throws Exception {
        GetResourcePoolBucketGroupQoSInfoResponse self = new GetResourcePoolBucketGroupQoSInfoResponse();
        return TeaModel.build(map, self);
    }

    public GetResourcePoolBucketGroupQoSInfoResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetResourcePoolBucketGroupQoSInfoResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetResourcePoolBucketGroupQoSInfoResponse setBody(GetResourcePoolBucketGroupQoSInfoResponseBody body) {
        this.body = body;
        return this;
    }
    public GetResourcePoolBucketGroupQoSInfoResponseBody getBody() {
        return this.body;
    }

}
