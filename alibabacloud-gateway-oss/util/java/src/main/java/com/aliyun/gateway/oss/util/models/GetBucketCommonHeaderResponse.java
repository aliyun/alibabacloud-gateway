// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketCommonHeaderResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketCommonHeaderResponseBody body;

    public static GetBucketCommonHeaderResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketCommonHeaderResponse self = new GetBucketCommonHeaderResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketCommonHeaderResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketCommonHeaderResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketCommonHeaderResponse setBody(GetBucketCommonHeaderResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketCommonHeaderResponseBody getBody() {
        return this.body;
    }

}
