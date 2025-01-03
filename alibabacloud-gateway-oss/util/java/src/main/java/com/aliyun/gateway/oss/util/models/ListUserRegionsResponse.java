// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListUserRegionsResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListUserRegionsResponseBody body;

    public static ListUserRegionsResponse build(java.util.Map<String, ?> map) throws Exception {
        ListUserRegionsResponse self = new ListUserRegionsResponse();
        return TeaModel.build(map, self);
    }

    public ListUserRegionsResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListUserRegionsResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListUserRegionsResponse setBody(ListUserRegionsResponseBody body) {
        this.body = body;
        return this;
    }
    public ListUserRegionsResponseBody getBody() {
        return this.body;
    }

}
