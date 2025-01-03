// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetObjectsResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public Object body;

    public static GetObjectsResponse build(java.util.Map<String, ?> map) throws Exception {
        GetObjectsResponse self = new GetObjectsResponse();
        return TeaModel.build(map, self);
    }

    public GetObjectsResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetObjectsResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetObjectsResponse setBody(Object body) {
        this.body = body;
        return this;
    }
    public Object getBody() {
        return this.body;
    }

}
