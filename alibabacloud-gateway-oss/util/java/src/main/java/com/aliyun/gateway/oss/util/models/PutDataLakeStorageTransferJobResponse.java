// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutDataLakeStorageTransferJobResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public PutDataLakeStorageTransferJobResponseBody body;

    public static PutDataLakeStorageTransferJobResponse build(java.util.Map<String, ?> map) throws Exception {
        PutDataLakeStorageTransferJobResponse self = new PutDataLakeStorageTransferJobResponse();
        return TeaModel.build(map, self);
    }

    public PutDataLakeStorageTransferJobResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public PutDataLakeStorageTransferJobResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public PutDataLakeStorageTransferJobResponse setBody(PutDataLakeStorageTransferJobResponseBody body) {
        this.body = body;
        return this;
    }
    public PutDataLakeStorageTransferJobResponseBody getBody() {
        return this.body;
    }

}
