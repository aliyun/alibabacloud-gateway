// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetResourcePoolRequesterQoSInfoResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetResourcePoolRequesterQoSInfoResponseBody body;

    public static GetResourcePoolRequesterQoSInfoResponse build(java.util.Map<String, ?> map) throws Exception {
        GetResourcePoolRequesterQoSInfoResponse self = new GetResourcePoolRequesterQoSInfoResponse();
        return TeaModel.build(map, self);
    }

    public GetResourcePoolRequesterQoSInfoResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetResourcePoolRequesterQoSInfoResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetResourcePoolRequesterQoSInfoResponse setBody(GetResourcePoolRequesterQoSInfoResponseBody body) {
        this.body = body;
        return this;
    }
    public GetResourcePoolRequesterQoSInfoResponseBody getBody() {
        return this.body;
    }

}
