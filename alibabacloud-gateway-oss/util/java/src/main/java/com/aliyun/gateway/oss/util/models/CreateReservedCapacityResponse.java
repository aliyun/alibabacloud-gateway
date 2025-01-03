// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateReservedCapacityResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public CreateReservedCapacityResponseBody body;

    public static CreateReservedCapacityResponse build(java.util.Map<String, ?> map) throws Exception {
        CreateReservedCapacityResponse self = new CreateReservedCapacityResponse();
        return TeaModel.build(map, self);
    }

    public CreateReservedCapacityResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public CreateReservedCapacityResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public CreateReservedCapacityResponse setBody(CreateReservedCapacityResponseBody body) {
        this.body = body;
        return this;
    }
    public CreateReservedCapacityResponseBody getBody() {
        return this.body;
    }

}
