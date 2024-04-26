// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ReplicationEncryptionConfiguration extends TeaModel {
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
