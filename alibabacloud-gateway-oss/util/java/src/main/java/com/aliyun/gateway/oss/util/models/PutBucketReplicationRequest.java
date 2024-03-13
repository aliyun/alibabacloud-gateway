// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketReplicationRequest extends TeaModel {
    /**
     * <p>The container that stores data replication configurations.</p>
     */
    @NameInMap("ReplicationConfiguration")
    public ReplicationConfiguration replicationConfiguration;

    public static PutBucketReplicationRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketReplicationRequest self = new PutBucketReplicationRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketReplicationRequest setReplicationConfiguration(ReplicationConfiguration replicationConfiguration) {
        this.replicationConfiguration = replicationConfiguration;
        return this;
    }
    public ReplicationConfiguration getReplicationConfiguration() {
        return this.replicationConfiguration;
    }

}
