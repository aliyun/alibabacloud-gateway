// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class DeleteBucketReplicationRequest extends TeaModel {
    /**
     * <p>The container that is used to store the data replication rule to delete.</p>
     */
    @NameInMap("ReplicationRules")
    public ReplicationRules replicationRules;

    public static DeleteBucketReplicationRequest build(java.util.Map<String, ?> map) throws Exception {
        DeleteBucketReplicationRequest self = new DeleteBucketReplicationRequest();
        return TeaModel.build(map, self);
    }

    public DeleteBucketReplicationRequest setReplicationRules(ReplicationRules replicationRules) {
        this.replicationRules = replicationRules;
        return this;
    }
    public ReplicationRules getReplicationRules() {
        return this.replicationRules;
    }

    public static class ReplicationRules extends TeaModel {
        @NameInMap("ID")
        public String ID;

        public static ReplicationRules build(java.util.Map<String, ?> map) throws Exception {
            ReplicationRules self = new ReplicationRules();
            return TeaModel.build(map, self);
        }

        public ReplicationRules setID(String ID) {
            this.ID = ID;
            return this;
        }
        public String getID() {
            return this.ID;
        }

    }

}
