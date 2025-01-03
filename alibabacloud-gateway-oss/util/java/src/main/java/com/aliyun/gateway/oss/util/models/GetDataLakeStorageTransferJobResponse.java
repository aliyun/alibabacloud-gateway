// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetDataLakeStorageTransferJobResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetDataLakeStorageTransferJobResponseBody body;

    public static GetDataLakeStorageTransferJobResponse build(java.util.Map<String, ?> map) throws Exception {
        GetDataLakeStorageTransferJobResponse self = new GetDataLakeStorageTransferJobResponse();
        return TeaModel.build(map, self);
    }

    public GetDataLakeStorageTransferJobResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetDataLakeStorageTransferJobResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetDataLakeStorageTransferJobResponse setBody(GetDataLakeStorageTransferJobResponseBody body) {
        this.body = body;
        return this;
    }
    public GetDataLakeStorageTransferJobResponseBody getBody() {
        return this.body;
    }

}
