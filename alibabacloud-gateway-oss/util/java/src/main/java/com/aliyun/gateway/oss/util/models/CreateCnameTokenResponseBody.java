// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateCnameTokenResponseBody extends TeaModel {
    /**
     * <p>The container in which the CNAME token is stored.</p>
     */
    @NameInMap("CnameToken")
    public CnameToken cnameToken;

    public static CreateCnameTokenResponseBody build(java.util.Map<String, ?> map) throws Exception {
        CreateCnameTokenResponseBody self = new CreateCnameTokenResponseBody();
        return TeaModel.build(map, self);
    }

    public CreateCnameTokenResponseBody setCnameToken(CnameToken cnameToken) {
        this.cnameToken = cnameToken;
        return this;
    }
    public CnameToken getCnameToken() {
        return this.cnameToken;
    }

}
