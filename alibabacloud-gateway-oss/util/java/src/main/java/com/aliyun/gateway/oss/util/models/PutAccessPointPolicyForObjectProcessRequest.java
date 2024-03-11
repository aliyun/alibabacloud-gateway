// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class PutAccessPointPolicyForObjectProcessRequest extends TeaModel {
    @NameInMap("body")
    public String body;

    public static PutAccessPointPolicyForObjectProcessRequest build(java.util.Map<String, ?> map) throws Exception {
        PutAccessPointPolicyForObjectProcessRequest self = new PutAccessPointPolicyForObjectProcessRequest();
        return TeaModel.build(map, self);
    }

    public PutAccessPointPolicyForObjectProcessRequest setBody(String body) {
        this.body = body;
        return this;
    }
    public String getBody() {
        return this.body;
    }

}
