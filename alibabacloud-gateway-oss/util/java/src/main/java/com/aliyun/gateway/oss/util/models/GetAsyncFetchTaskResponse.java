// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetAsyncFetchTaskResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetAsyncFetchTaskResponseBody body;

    public static GetAsyncFetchTaskResponse build(java.util.Map<String, ?> map) throws Exception {
        GetAsyncFetchTaskResponse self = new GetAsyncFetchTaskResponse();
        return TeaModel.build(map, self);
    }

    public GetAsyncFetchTaskResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetAsyncFetchTaskResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetAsyncFetchTaskResponse setBody(GetAsyncFetchTaskResponseBody body) {
        this.body = body;
        return this;
    }
    public GetAsyncFetchTaskResponseBody getBody() {
        return this.body;
    }

}
