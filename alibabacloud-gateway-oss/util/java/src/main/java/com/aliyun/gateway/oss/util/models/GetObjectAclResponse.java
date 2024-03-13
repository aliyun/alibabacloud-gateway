// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetObjectAclResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetObjectAclResponseBody body;

    public static GetObjectAclResponse build(java.util.Map<String, ?> map) throws Exception {
        GetObjectAclResponse self = new GetObjectAclResponse();
        return TeaModel.build(map, self);
    }

    public GetObjectAclResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetObjectAclResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetObjectAclResponse setBody(GetObjectAclResponseBody body) {
        this.body = body;
        return this;
    }
    public GetObjectAclResponseBody getBody() {
        return this.body;
    }

}
