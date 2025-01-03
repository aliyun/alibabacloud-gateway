// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetProcessConfigurationResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetProcessConfigurationResponseBody body;

    public static GetProcessConfigurationResponse build(java.util.Map<String, ?> map) throws Exception {
        GetProcessConfigurationResponse self = new GetProcessConfigurationResponse();
        return TeaModel.build(map, self);
    }

    public GetProcessConfigurationResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetProcessConfigurationResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetProcessConfigurationResponse setBody(GetProcessConfigurationResponseBody body) {
        this.body = body;
        return this;
    }
    public GetProcessConfigurationResponseBody getBody() {
        return this.body;
    }

}
