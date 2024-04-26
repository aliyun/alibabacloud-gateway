// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetUserDefinedLogFieldsConfigResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetUserDefinedLogFieldsConfigResponseBody body;

    public static GetUserDefinedLogFieldsConfigResponse build(java.util.Map<String, ?> map) throws Exception {
        GetUserDefinedLogFieldsConfigResponse self = new GetUserDefinedLogFieldsConfigResponse();
        return TeaModel.build(map, self);
    }

    public GetUserDefinedLogFieldsConfigResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetUserDefinedLogFieldsConfigResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetUserDefinedLogFieldsConfigResponse setBody(GetUserDefinedLogFieldsConfigResponseBody body) {
        this.body = body;
        return this;
    }
    public GetUserDefinedLogFieldsConfigResponseBody getBody() {
        return this.body;
    }

}
