// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class GetBucketWebsiteResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketWebsiteResponseBody body;

    public static GetBucketWebsiteResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketWebsiteResponse self = new GetBucketWebsiteResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketWebsiteResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketWebsiteResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketWebsiteResponse setBody(GetBucketWebsiteResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketWebsiteResponseBody getBody() {
        return this.body;
    }

}
