// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetUserQoSInfoResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetUserQoSInfoResponseBody body;

    public static GetUserQoSInfoResponse build(java.util.Map<String, ?> map) throws Exception {
        GetUserQoSInfoResponse self = new GetUserQoSInfoResponse();
        return TeaModel.build(map, self);
    }

    public GetUserQoSInfoResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetUserQoSInfoResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetUserQoSInfoResponse setBody(GetUserQoSInfoResponseBody body) {
        this.body = body;
        return this;
    }
    public GetUserQoSInfoResponseBody getBody() {
        return this.body;
    }

}
