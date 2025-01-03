// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListDataLakeStorageTransferJobResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListDataLakeStorageTransferJobResponseBody body;

    public static ListDataLakeStorageTransferJobResponse build(java.util.Map<String, ?> map) throws Exception {
        ListDataLakeStorageTransferJobResponse self = new ListDataLakeStorageTransferJobResponse();
        return TeaModel.build(map, self);
    }

    public ListDataLakeStorageTransferJobResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListDataLakeStorageTransferJobResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListDataLakeStorageTransferJobResponse setBody(ListDataLakeStorageTransferJobResponseBody body) {
        this.body = body;
        return this;
    }
    public ListDataLakeStorageTransferJobResponseBody getBody() {
        return this.body;
    }

}
