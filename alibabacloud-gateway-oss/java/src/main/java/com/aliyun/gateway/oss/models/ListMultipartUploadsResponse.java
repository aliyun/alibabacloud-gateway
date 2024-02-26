// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class ListMultipartUploadsResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListMultipartUploadsResponseBody body;

    public static ListMultipartUploadsResponse build(java.util.Map<String, ?> map) throws Exception {
        ListMultipartUploadsResponse self = new ListMultipartUploadsResponse();
        return TeaModel.build(map, self);
    }

    public ListMultipartUploadsResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListMultipartUploadsResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListMultipartUploadsResponse setBody(ListMultipartUploadsResponseBody body) {
        this.body = body;
        return this;
    }
    public ListMultipartUploadsResponseBody getBody() {
        return this.body;
    }

}
