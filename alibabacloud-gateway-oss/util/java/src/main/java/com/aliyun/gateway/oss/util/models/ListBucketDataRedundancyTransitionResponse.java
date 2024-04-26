// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListBucketDataRedundancyTransitionResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListBucketDataRedundancyTransitionResponseBody body;

    public static ListBucketDataRedundancyTransitionResponse build(java.util.Map<String, ?> map) throws Exception {
        ListBucketDataRedundancyTransitionResponse self = new ListBucketDataRedundancyTransitionResponse();
        return TeaModel.build(map, self);
    }

    public ListBucketDataRedundancyTransitionResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListBucketDataRedundancyTransitionResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListBucketDataRedundancyTransitionResponse setBody(ListBucketDataRedundancyTransitionResponseBody body) {
        this.body = body;
        return this;
    }
    public ListBucketDataRedundancyTransitionResponseBody getBody() {
        return this.body;
    }

}
