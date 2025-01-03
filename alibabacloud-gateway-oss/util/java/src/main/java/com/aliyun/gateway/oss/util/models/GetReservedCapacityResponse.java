// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetReservedCapacityResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetReservedCapacityResponseBody body;

    public static GetReservedCapacityResponse build(java.util.Map<String, ?> map) throws Exception {
        GetReservedCapacityResponse self = new GetReservedCapacityResponse();
        return TeaModel.build(map, self);
    }

    public GetReservedCapacityResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetReservedCapacityResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetReservedCapacityResponse setBody(GetReservedCapacityResponseBody body) {
        this.body = body;
        return this;
    }
    public GetReservedCapacityResponseBody getBody() {
        return this.body;
    }

}
