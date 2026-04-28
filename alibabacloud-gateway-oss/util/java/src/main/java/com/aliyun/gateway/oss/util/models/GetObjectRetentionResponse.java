// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetObjectRetentionResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetObjectRetentionResponseBody body;

    public static GetObjectRetentionResponse build(java.util.Map<String, ?> map) throws Exception {
        GetObjectRetentionResponse self = new GetObjectRetentionResponse();
        return TeaModel.build(map, self);
    }

    public GetObjectRetentionResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetObjectRetentionResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetObjectRetentionResponse setBody(GetObjectRetentionResponseBody body) {
        this.body = body;
        return this;
    }
    public GetObjectRetentionResponseBody getBody() {
        return this.body;
    }

}
