// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class StartDataLakeStorageTransferJobResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public StartDataLakeStorageTransferJobResponseBody body;

    public static StartDataLakeStorageTransferJobResponse build(java.util.Map<String, ?> map) throws Exception {
        StartDataLakeStorageTransferJobResponse self = new StartDataLakeStorageTransferJobResponse();
        return TeaModel.build(map, self);
    }

    public StartDataLakeStorageTransferJobResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public StartDataLakeStorageTransferJobResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public StartDataLakeStorageTransferJobResponse setBody(StartDataLakeStorageTransferJobResponseBody body) {
        this.body = body;
        return this;
    }
    public StartDataLakeStorageTransferJobResponseBody getBody() {
        return this.body;
    }

}
