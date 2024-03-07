// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class CreateSelectObjectMetaResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public SelectMetaStatus body;

    public static CreateSelectObjectMetaResponse build(java.util.Map<String, ?> map) throws Exception {
        CreateSelectObjectMetaResponse self = new CreateSelectObjectMetaResponse();
        return TeaModel.build(map, self);
    }

    public CreateSelectObjectMetaResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public CreateSelectObjectMetaResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public CreateSelectObjectMetaResponse setBody(SelectMetaStatus body) {
        this.body = body;
        return this;
    }
    public SelectMetaStatus getBody() {
        return this.body;
    }

}
