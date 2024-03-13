// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetMetaQueryStatusResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetMetaQueryStatusResponseBody body;

    public static GetMetaQueryStatusResponse build(java.util.Map<String, ?> map) throws Exception {
        GetMetaQueryStatusResponse self = new GetMetaQueryStatusResponse();
        return TeaModel.build(map, self);
    }

    public GetMetaQueryStatusResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetMetaQueryStatusResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetMetaQueryStatusResponse setBody(GetMetaQueryStatusResponseBody body) {
        this.body = body;
        return this;
    }
    public GetMetaQueryStatusResponseBody getBody() {
        return this.body;
    }

}
