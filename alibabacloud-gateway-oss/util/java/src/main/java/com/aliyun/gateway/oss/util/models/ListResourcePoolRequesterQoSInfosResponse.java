// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListResourcePoolRequesterQoSInfosResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListResourcePoolRequesterQoSInfosResponseBody body;

    public static ListResourcePoolRequesterQoSInfosResponse build(java.util.Map<String, ?> map) throws Exception {
        ListResourcePoolRequesterQoSInfosResponse self = new ListResourcePoolRequesterQoSInfosResponse();
        return TeaModel.build(map, self);
    }

    public ListResourcePoolRequesterQoSInfosResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListResourcePoolRequesterQoSInfosResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListResourcePoolRequesterQoSInfosResponse setBody(ListResourcePoolRequesterQoSInfosResponseBody body) {
        this.body = body;
        return this;
    }
    public ListResourcePoolRequesterQoSInfosResponseBody getBody() {
        return this.body;
    }

}
