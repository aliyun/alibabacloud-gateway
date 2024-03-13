// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListPartsResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListPartsResponseBody body;

    public static ListPartsResponse build(java.util.Map<String, ?> map) throws Exception {
        ListPartsResponse self = new ListPartsResponse();
        return TeaModel.build(map, self);
    }

    public ListPartsResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListPartsResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListPartsResponse setBody(ListPartsResponseBody body) {
        this.body = body;
        return this;
    }
    public ListPartsResponseBody getBody() {
        return this.body;
    }

}
