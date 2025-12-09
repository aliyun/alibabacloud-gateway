// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class UpdateJobPriorityResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public UpdateJobPriorityResponseBody body;

    public static UpdateJobPriorityResponse build(java.util.Map<String, ?> map) throws Exception {
        UpdateJobPriorityResponse self = new UpdateJobPriorityResponse();
        return TeaModel.build(map, self);
    }

    public UpdateJobPriorityResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public UpdateJobPriorityResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public UpdateJobPriorityResponse setBody(UpdateJobPriorityResponseBody body) {
        this.body = body;
        return this;
    }
    public UpdateJobPriorityResponseBody getBody() {
        return this.body;
    }

}
