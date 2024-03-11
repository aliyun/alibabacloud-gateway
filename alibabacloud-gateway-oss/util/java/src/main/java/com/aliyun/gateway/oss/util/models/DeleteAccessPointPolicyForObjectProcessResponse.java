// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class DeleteAccessPointPolicyForObjectProcessResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    public static DeleteAccessPointPolicyForObjectProcessResponse build(java.util.Map<String, ?> map) throws Exception {
        DeleteAccessPointPolicyForObjectProcessResponse self = new DeleteAccessPointPolicyForObjectProcessResponse();
        return TeaModel.build(map, self);
    }

    public DeleteAccessPointPolicyForObjectProcessResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public DeleteAccessPointPolicyForObjectProcessResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

}
