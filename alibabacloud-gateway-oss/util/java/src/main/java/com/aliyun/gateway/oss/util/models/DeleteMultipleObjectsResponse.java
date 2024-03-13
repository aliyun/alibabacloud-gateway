// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DeleteMultipleObjectsResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public DeleteMultipleObjectsResponseBody body;

    public static DeleteMultipleObjectsResponse build(java.util.Map<String, ?> map) throws Exception {
        DeleteMultipleObjectsResponse self = new DeleteMultipleObjectsResponse();
        return TeaModel.build(map, self);
    }

    public DeleteMultipleObjectsResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public DeleteMultipleObjectsResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public DeleteMultipleObjectsResponse setBody(DeleteMultipleObjectsResponseBody body) {
        this.body = body;
        return this;
    }
    public DeleteMultipleObjectsResponseBody getBody() {
        return this.body;
    }

}
