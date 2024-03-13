// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketReplicationProgressResponseBody extends TeaModel {
    /**
     * <p>The container that is used to store the progress of data replication tasks.</p>
     */
    @NameInMap("ReplicationProgress")
    public ReplicationProgress replicationProgress;

    public static GetBucketReplicationProgressResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketReplicationProgressResponseBody self = new GetBucketReplicationProgressResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketReplicationProgressResponseBody setReplicationProgress(ReplicationProgress replicationProgress) {
        this.replicationProgress = replicationProgress;
        return this;
    }
    public ReplicationProgress getReplicationProgress() {
        return this.replicationProgress;
    }

    public static class ReplicationProgress extends TeaModel {
        /**
         * <p>The container that stores the progress of the data replication task corresponding to each data replication rule.</p>
         */
        @NameInMap("Rule")
        public java.util.List<ReplicationProgressRule> rule;

        public static ReplicationProgress build(java.util.Map<String, ?> map) throws Exception {
            ReplicationProgress self = new ReplicationProgress();
            return TeaModel.build(map, self);
        }

        public ReplicationProgress setRule(java.util.List<ReplicationProgressRule> rule) {
            this.rule = rule;
            return this;
        }
        public java.util.List<ReplicationProgressRule> getRule() {
            return this.rule;
        }

    }

}
