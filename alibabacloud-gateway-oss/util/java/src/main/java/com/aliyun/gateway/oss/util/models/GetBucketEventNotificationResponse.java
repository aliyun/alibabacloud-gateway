// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketEventNotificationResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketEventNotificationResponseBody body;

    public static GetBucketEventNotificationResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketEventNotificationResponse self = new GetBucketEventNotificationResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketEventNotificationResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketEventNotificationResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketEventNotificationResponse setBody(GetBucketEventNotificationResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketEventNotificationResponseBody getBody() {
        return this.body;
    }

}
