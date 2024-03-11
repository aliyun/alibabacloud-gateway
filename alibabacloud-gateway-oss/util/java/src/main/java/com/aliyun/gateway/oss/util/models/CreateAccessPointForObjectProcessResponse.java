// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class CreateAccessPointForObjectProcessResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public CreateAccessPointForObjectProcessResponseBody body;

    public static CreateAccessPointForObjectProcessResponse build(java.util.Map<String, ?> map) throws Exception {
        CreateAccessPointForObjectProcessResponse self = new CreateAccessPointForObjectProcessResponse();
        return TeaModel.build(map, self);
    }

    public CreateAccessPointForObjectProcessResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public CreateAccessPointForObjectProcessResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public CreateAccessPointForObjectProcessResponse setBody(CreateAccessPointForObjectProcessResponseBody body) {
        this.body = body;
        return this;
    }
    public CreateAccessPointForObjectProcessResponseBody getBody() {
        return this.body;
    }

}
