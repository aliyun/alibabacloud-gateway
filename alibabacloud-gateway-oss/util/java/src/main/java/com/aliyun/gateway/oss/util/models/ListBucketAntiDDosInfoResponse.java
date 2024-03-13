// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListBucketAntiDDosInfoResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListBucketAntiDDosInfoResponseBody body;

    public static ListBucketAntiDDosInfoResponse build(java.util.Map<String, ?> map) throws Exception {
        ListBucketAntiDDosInfoResponse self = new ListBucketAntiDDosInfoResponse();
        return TeaModel.build(map, self);
    }

    public ListBucketAntiDDosInfoResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListBucketAntiDDosInfoResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListBucketAntiDDosInfoResponse setBody(ListBucketAntiDDosInfoResponseBody body) {
        this.body = body;
        return this;
    }
    public ListBucketAntiDDosInfoResponseBody getBody() {
        return this.body;
    }

}
