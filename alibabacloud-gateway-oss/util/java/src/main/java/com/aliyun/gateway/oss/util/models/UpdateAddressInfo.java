// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class UpdateAddressInfo extends TeaModel {
    @NameInMap("AgentList")
    public String agentList;

    public static UpdateAddressInfo build(java.util.Map<String, ?> map) throws Exception {
        UpdateAddressInfo self = new UpdateAddressInfo();
        return TeaModel.build(map, self);
    }

    public UpdateAddressInfo setAgentList(String agentList) {
        this.agentList = agentList;
        return this;
    }
    public String getAgentList() {
        return this.agentList;
    }

}
