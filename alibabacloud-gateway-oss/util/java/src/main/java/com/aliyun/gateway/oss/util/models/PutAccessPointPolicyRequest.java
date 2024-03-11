// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class PutAccessPointPolicyRequest extends TeaModel {
    /**
     * <p>The configurations of the access point policy.</p>
     */
    @NameInMap("body")
    public String body;

    public static PutAccessPointPolicyRequest build(java.util.Map<String, ?> map) throws Exception {
        PutAccessPointPolicyRequest self = new PutAccessPointPolicyRequest();
        return TeaModel.build(map, self);
    }

    public PutAccessPointPolicyRequest setBody(String body) {
        this.body = body;
        return this;
    }
    public String getBody() {
        return this.body;
    }

}
