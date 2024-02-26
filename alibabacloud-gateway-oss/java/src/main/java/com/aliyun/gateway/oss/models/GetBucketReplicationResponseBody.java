// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class GetBucketReplicationResponseBody extends TeaModel {
    /**
     * <p>The container that stores data replication configurations.</p>
     */
    @NameInMap("ReplicationConfiguration")
    public ReplicationConfiguration replicationConfiguration;

    public static GetBucketReplicationResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketReplicationResponseBody self = new GetBucketReplicationResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketReplicationResponseBody setReplicationConfiguration(ReplicationConfiguration replicationConfiguration) {
        this.replicationConfiguration = replicationConfiguration;
        return this;
    }
    public ReplicationConfiguration getReplicationConfiguration() {
        return this.replicationConfiguration;
    }

    public static class ReplicationConfiguration extends TeaModel {
        /**
         * <p>The container that stores the data replication rules.</p>
         */
        @NameInMap("Rule")
        public java.util.List<ReplicationRule> rules;

        public static ReplicationConfiguration build(java.util.Map<String, ?> map) throws Exception {
            ReplicationConfiguration self = new ReplicationConfiguration();
            return TeaModel.build(map, self);
        }

        public ReplicationConfiguration setRules(java.util.List<ReplicationRule> rules) {
            this.rules = rules;
            return this;
        }
        public java.util.List<ReplicationRule> getRules() {
            return this.rules;
        }

    }

}
