// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PostObjectGroupResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public PostObjectGroupResponseBody body;

    public static PostObjectGroupResponse build(java.util.Map<String, ?> map) throws Exception {
        PostObjectGroupResponse self = new PostObjectGroupResponse();
        return TeaModel.build(map, self);
    }

    public PostObjectGroupResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public PostObjectGroupResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public PostObjectGroupResponse setBody(PostObjectGroupResponseBody body) {
        this.body = body;
        return this;
    }
    public PostObjectGroupResponseBody getBody() {
        return this.body;
    }

}
