// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ReplicationEncryptionConfiguration extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>c4d49f85-ee30-426b-a5ed-95e9139dxxxxx</p>
     */
    @NameInMap("ReplicaKmsKeyID")
    public String replicaKmsKeyID;

    public static ReplicationEncryptionConfiguration build(java.util.Map<String, ?> map) throws Exception {
        ReplicationEncryptionConfiguration self = new ReplicationEncryptionConfiguration();
        return TeaModel.build(map, self);
    }

    public ReplicationEncryptionConfiguration setReplicaKmsKeyID(String replicaKmsKeyID) {
        this.replicaKmsKeyID = replicaKmsKeyID;
        return this;
    }
    public String getReplicaKmsKeyID() {
        return this.replicaKmsKeyID;
    }

}
