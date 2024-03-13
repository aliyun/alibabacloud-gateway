// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateCnameTokenResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public CreateCnameTokenResponseBody body;

    public static CreateCnameTokenResponse build(java.util.Map<String, ?> map) throws Exception {
        CreateCnameTokenResponse self = new CreateCnameTokenResponse();
        return TeaModel.build(map, self);
    }

    public CreateCnameTokenResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public CreateCnameTokenResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public CreateCnameTokenResponse setBody(CreateCnameTokenResponseBody body) {
        this.body = body;
        return this;
    }
    public CreateCnameTokenResponseBody getBody() {
        return this.body;
    }

}
