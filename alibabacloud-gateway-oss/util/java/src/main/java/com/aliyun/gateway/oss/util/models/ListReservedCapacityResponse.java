// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListReservedCapacityResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListReservedCapacityResponseBody body;

    public static ListReservedCapacityResponse build(java.util.Map<String, ?> map) throws Exception {
        ListReservedCapacityResponse self = new ListReservedCapacityResponse();
        return TeaModel.build(map, self);
    }

    public ListReservedCapacityResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListReservedCapacityResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListReservedCapacityResponse setBody(ListReservedCapacityResponseBody body) {
        this.body = body;
        return this;
    }
    public ListReservedCapacityResponseBody getBody() {
        return this.body;
    }

}
