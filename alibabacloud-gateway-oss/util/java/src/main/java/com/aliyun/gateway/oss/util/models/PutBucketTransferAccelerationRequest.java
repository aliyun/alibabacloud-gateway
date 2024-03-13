// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketTransferAccelerationRequest extends TeaModel {
    /**
     * <p>The container that stores the transfer acceleration configurations.</p>
     */
    @NameInMap("TransferAccelerationConfiguration")
    public TransferAccelerationConfiguration transferAccelerationConfiguration;

    public static PutBucketTransferAccelerationRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketTransferAccelerationRequest self = new PutBucketTransferAccelerationRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketTransferAccelerationRequest setTransferAccelerationConfiguration(TransferAccelerationConfiguration transferAccelerationConfiguration) {
        this.transferAccelerationConfiguration = transferAccelerationConfiguration;
        return this;
    }
    public TransferAccelerationConfiguration getTransferAccelerationConfiguration() {
        return this.transferAccelerationConfiguration;
    }

}
