// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class UpdateJobStatusResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public UpdateJobStatusResponseBody body;

    public static UpdateJobStatusResponse build(java.util.Map<String, ?> map) throws Exception {
        UpdateJobStatusResponse self = new UpdateJobStatusResponse();
        return TeaModel.build(map, self);
    }

    public UpdateJobStatusResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public UpdateJobStatusResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public UpdateJobStatusResponse setBody(UpdateJobStatusResponseBody body) {
        this.body = body;
        return this;
    }
    public UpdateJobStatusResponseBody getBody() {
        return this.body;
    }

}
