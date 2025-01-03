// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ReplicationDestination extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>destbucket</p>
     */
    @NameInMap("Bucket")
    public String bucket;

    /**
     * <strong>example:</strong>
     * <p>oss-cn-hangzhou</p>
     */
    @NameInMap("Location")
    public String location;

    /**
     * <strong>example:</strong>
     * <p>internal</p>
     */
    @NameInMap("TransferType")
    public String transferType;

    public static ReplicationDestination build(java.util.Map<String, ?> map) throws Exception {
        ReplicationDestination self = new ReplicationDestination();
        return TeaModel.build(map, self);
    }

    public ReplicationDestination setBucket(String bucket) {
        this.bucket = bucket;
        return this;
    }
    public String getBucket() {
        return this.bucket;
    }

    public ReplicationDestination setLocation(String location) {
        this.location = location;
        return this;
    }
    public String getLocation() {
        return this.location;
    }

    public ReplicationDestination setTransferType(String transferType) {
        this.transferType = transferType;
        return this;
    }
    public String getTransferType() {
        return this.transferType;
    }

}
