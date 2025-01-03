// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class AccessPointVpcConfiguration extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>vpc-t4nlw426y44rd3iq4****</p>
     */
    @NameInMap("VpcId")
    public String vpcId;

    public static AccessPointVpcConfiguration build(java.util.Map<String, ?> map) throws Exception {
        AccessPointVpcConfiguration self = new AccessPointVpcConfiguration();
        return TeaModel.build(map, self);
    }

    public AccessPointVpcConfiguration setVpcId(String vpcId) {
        this.vpcId = vpcId;
        return this;
    }
    public String getVpcId() {
        return this.vpcId;
    }

}
