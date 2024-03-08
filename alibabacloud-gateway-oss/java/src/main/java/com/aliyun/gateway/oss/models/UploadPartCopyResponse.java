// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class UploadPartCopyResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public UploadPartCopyResponseBody body;

    public static UploadPartCopyResponse build(java.util.Map<String, ?> map) throws Exception {
        UploadPartCopyResponse self = new UploadPartCopyResponse();
        return TeaModel.build(map, self);
    }

    public UploadPartCopyResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public UploadPartCopyResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public UploadPartCopyResponse setBody(UploadPartCopyResponseBody body) {
        this.body = body;
        return this;
    }
    public UploadPartCopyResponseBody getBody() {
        return this.body;
    }

}
