// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class DeleteBucketResponseHeaderResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    public static DeleteBucketResponseHeaderResponse build(java.util.Map<String, ?> map) throws Exception {
        DeleteBucketResponseHeaderResponse self = new DeleteBucketResponseHeaderResponse();
        return TeaModel.build(map, self);
    }

    public DeleteBucketResponseHeaderResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public DeleteBucketResponseHeaderResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

}
