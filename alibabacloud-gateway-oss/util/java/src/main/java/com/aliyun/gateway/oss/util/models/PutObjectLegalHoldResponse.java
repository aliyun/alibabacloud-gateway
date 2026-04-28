// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.oss20190517.models;

import com.aliyun.tea.*;

public class PutObjectLegalHoldResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    public static PutObjectLegalHoldResponse build(java.util.Map<String, ?> map) throws Exception {
        PutObjectLegalHoldResponse self = new PutObjectLegalHoldResponse();
        return TeaModel.build(map, self);
    }

    public PutObjectLegalHoldResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public PutObjectLegalHoldResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

}
