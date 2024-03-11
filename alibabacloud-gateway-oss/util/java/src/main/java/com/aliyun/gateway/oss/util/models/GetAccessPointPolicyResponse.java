// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class GetAccessPointPolicyResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public String body;

    public static GetAccessPointPolicyResponse build(java.util.Map<String, ?> map) throws Exception {
        GetAccessPointPolicyResponse self = new GetAccessPointPolicyResponse();
        return TeaModel.build(map, self);
    }

    public GetAccessPointPolicyResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetAccessPointPolicyResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetAccessPointPolicyResponse setBody(String body) {
        this.body = body;
        return this;
    }
    public String getBody() {
        return this.body;
    }

}
