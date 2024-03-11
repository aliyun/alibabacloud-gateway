// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class GetCnameTokenRequest extends TeaModel {
    /**
     * <p>The name of the CNAME record that is mapped to the bucket.</p>
     */
    @NameInMap("cname")
    public String cname;

    public static GetCnameTokenRequest build(java.util.Map<String, ?> map) throws Exception {
        GetCnameTokenRequest self = new GetCnameTokenRequest();
        return TeaModel.build(map, self);
    }

    public GetCnameTokenRequest setCname(String cname) {
        this.cname = cname;
        return this;
    }
    public String getCname() {
        return this.cname;
    }

}
