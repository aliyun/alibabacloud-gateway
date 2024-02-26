// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class GetUserAntiDDosInfoResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetUserAntiDDosInfoResponseBody body;

    public static GetUserAntiDDosInfoResponse build(java.util.Map<String, ?> map) throws Exception {
        GetUserAntiDDosInfoResponse self = new GetUserAntiDDosInfoResponse();
        return TeaModel.build(map, self);
    }

    public GetUserAntiDDosInfoResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetUserAntiDDosInfoResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetUserAntiDDosInfoResponse setBody(GetUserAntiDDosInfoResponseBody body) {
        this.body = body;
        return this;
    }
    public GetUserAntiDDosInfoResponseBody getBody() {
        return this.body;
    }

}
