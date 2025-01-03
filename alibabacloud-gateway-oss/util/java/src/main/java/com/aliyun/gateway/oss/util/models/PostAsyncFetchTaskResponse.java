// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PostAsyncFetchTaskResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public PostAsyncFetchTaskResponseBody body;

    public static PostAsyncFetchTaskResponse build(java.util.Map<String, ?> map) throws Exception {
        PostAsyncFetchTaskResponse self = new PostAsyncFetchTaskResponse();
        return TeaModel.build(map, self);
    }

    public PostAsyncFetchTaskResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public PostAsyncFetchTaskResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public PostAsyncFetchTaskResponse setBody(PostAsyncFetchTaskResponseBody body) {
        this.body = body;
        return this;
    }
    public PostAsyncFetchTaskResponseBody getBody() {
        return this.body;
    }

}
