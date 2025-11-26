// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.hcs_mgw_models;

import com.aliyun.tea.*;

public class DeleteAddressResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    public static DeleteAddressResponse build(java.util.Map<String, ?> map) throws Exception {
        DeleteAddressResponse self = new DeleteAddressResponse();
        return TeaModel.build(map, self);
    }

    public DeleteAddressResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public DeleteAddressResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

}
