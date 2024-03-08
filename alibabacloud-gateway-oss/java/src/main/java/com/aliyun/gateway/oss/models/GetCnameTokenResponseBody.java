// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class GetCnameTokenResponseBody extends TeaModel {
    /**
     * <p>The container in which the CNAME token is stored.</p>
     */
    @NameInMap("CnameToken")
    public CnameToken cnameToken;

    public static GetCnameTokenResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetCnameTokenResponseBody self = new GetCnameTokenResponseBody();
        return TeaModel.build(map, self);
    }

    public GetCnameTokenResponseBody setCnameToken(CnameToken cnameToken) {
        this.cnameToken = cnameToken;
        return this;
    }
    public CnameToken getCnameToken() {
        return this.cnameToken;
    }

}
