// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.oss20190517.models;

import com.aliyun.tea.*;

public class GetObjectLegalHoldResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetObjectLegalHoldResponseBody body;

    public static GetObjectLegalHoldResponse build(java.util.Map<String, ?> map) throws Exception {
        GetObjectLegalHoldResponse self = new GetObjectLegalHoldResponse();
        return TeaModel.build(map, self);
    }

    public GetObjectLegalHoldResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetObjectLegalHoldResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetObjectLegalHoldResponse setBody(GetObjectLegalHoldResponseBody body) {
        this.body = body;
        return this;
    }
    public GetObjectLegalHoldResponseBody getBody() {
        return this.body;
    }

}
