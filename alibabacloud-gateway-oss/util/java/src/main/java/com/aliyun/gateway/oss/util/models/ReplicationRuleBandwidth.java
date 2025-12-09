// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ReplicationRuleBandwidth extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>209715200</p>
     */
    @NameInMap("Bandwidth")
    public Long bandwidth;

    /**
     * <strong>example:</strong>
     * <p>test_replication_1</p>
     */
    @NameInMap("ID")
    public String ID;

    public static ReplicationRuleBandwidth build(java.util.Map<String, ?> map) throws Exception {
        ReplicationRuleBandwidth self = new ReplicationRuleBandwidth();
        return TeaModel.build(map, self);
    }

    public ReplicationRuleBandwidth setBandwidth(Long bandwidth) {
        this.bandwidth = bandwidth;
        return this;
    }
    public Long getBandwidth() {
        return this.bandwidth;
    }

    public ReplicationRuleBandwidth setID(String ID) {
        this.ID = ID;
        return this;
    }
    public String getID() {
        return this.ID;
    }

}
