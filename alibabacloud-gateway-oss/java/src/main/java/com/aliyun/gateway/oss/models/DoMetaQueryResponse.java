// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class DoMetaQueryResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public DoMetaQueryResponseBody body;

    public static DoMetaQueryResponse build(java.util.Map<String, ?> map) throws Exception {
        DoMetaQueryResponse self = new DoMetaQueryResponse();
        return TeaModel.build(map, self);
    }

    public DoMetaQueryResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public DoMetaQueryResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public DoMetaQueryResponse setBody(DoMetaQueryResponseBody body) {
        this.body = body;
        return this;
    }
    public DoMetaQueryResponseBody getBody() {
        return this.body;
    }

}
