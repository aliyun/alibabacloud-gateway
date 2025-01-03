// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PostDataLakeStorageSecurityOperationResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public String body;

    public static PostDataLakeStorageSecurityOperationResponse build(java.util.Map<String, ?> map) throws Exception {
        PostDataLakeStorageSecurityOperationResponse self = new PostDataLakeStorageSecurityOperationResponse();
        return TeaModel.build(map, self);
    }

    public PostDataLakeStorageSecurityOperationResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public PostDataLakeStorageSecurityOperationResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public PostDataLakeStorageSecurityOperationResponse setBody(String body) {
        this.body = body;
        return this;
    }
    public String getBody() {
        return this.body;
    }

}
