// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class GetBucketResourceGroupResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketResourceGroupResponseBody body;

    public static GetBucketResourceGroupResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketResourceGroupResponse self = new GetBucketResourceGroupResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketResourceGroupResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketResourceGroupResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketResourceGroupResponse setBody(GetBucketResourceGroupResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketResourceGroupResponseBody getBody() {
        return this.body;
    }

}
