// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketReplicationBandwidthResponseBody extends TeaModel {
    @NameInMap("ReplicationRule")
    public ReplicationRule replicationRule;

    public static PutBucketReplicationBandwidthResponseBody build(java.util.Map<String, ?> map) throws Exception {
        PutBucketReplicationBandwidthResponseBody self = new PutBucketReplicationBandwidthResponseBody();
        return TeaModel.build(map, self);
    }

    public PutBucketReplicationBandwidthResponseBody setReplicationRule(ReplicationRule replicationRule) {
        this.replicationRule = replicationRule;
        return this;
    }
    public ReplicationRule getReplicationRule() {
        return this.replicationRule;
    }

    public static class ReplicationRule extends TeaModel {
        @NameInMap("Bandwith")
        public Long bandwith;

        @NameInMap("ID")
        public String ID;

        public static ReplicationRule build(java.util.Map<String, ?> map) throws Exception {
            ReplicationRule self = new ReplicationRule();
            return TeaModel.build(map, self);
        }

        public ReplicationRule setBandwith(Long bandwith) {
            this.bandwith = bandwith;
            return this;
        }
        public Long getBandwith() {
            return this.bandwith;
        }

        public ReplicationRule setID(String ID) {
            this.ID = ID;
            return this;
        }
        public String getID() {
            return this.ID;
        }

    }

}
