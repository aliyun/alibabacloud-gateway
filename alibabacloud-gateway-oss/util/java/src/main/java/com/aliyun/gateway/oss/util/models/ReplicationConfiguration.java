// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ReplicationConfiguration extends TeaModel {
    @NameInMap("Rule")
    public PutReplicationRule rule;

    public static ReplicationConfiguration build(java.util.Map<String, ?> map) throws Exception {
        ReplicationConfiguration self = new ReplicationConfiguration();
        return TeaModel.build(map, self);
    }

    public ReplicationConfiguration setRule(PutReplicationRule rule) {
        this.rule = rule;
        return this;
    }
    public PutReplicationRule getRule() {
        return this.rule;
    }

}
