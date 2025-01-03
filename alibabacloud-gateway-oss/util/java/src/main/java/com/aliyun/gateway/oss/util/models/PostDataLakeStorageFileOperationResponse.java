// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PostDataLakeStorageFileOperationResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public String body;

    public static PostDataLakeStorageFileOperationResponse build(java.util.Map<String, ?> map) throws Exception {
        PostDataLakeStorageFileOperationResponse self = new PostDataLakeStorageFileOperationResponse();
        return TeaModel.build(map, self);
    }

    public PostDataLakeStorageFileOperationResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public PostDataLakeStorageFileOperationResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public PostDataLakeStorageFileOperationResponse setBody(String body) {
        this.body = body;
        return this;
    }
    public String getBody() {
        return this.body;
    }

}
