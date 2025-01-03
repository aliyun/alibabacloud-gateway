// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketNotificationResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketNotificationResponseBody body;

    public static GetBucketNotificationResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketNotificationResponse self = new GetBucketNotificationResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketNotificationResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketNotificationResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketNotificationResponse setBody(GetBucketNotificationResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketNotificationResponseBody getBody() {
        return this.body;
    }

}
