// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketResourceGroupRequest extends TeaModel {
    /**
     * <p>The container that contains the ID of the resource group.</p>
     */
    @NameInMap("BucketResourceGroupConfiguration")
    public BucketResourceGroupConfiguration bucketResourceGroupConfiguration;

    public static PutBucketResourceGroupRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketResourceGroupRequest self = new PutBucketResourceGroupRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketResourceGroupRequest setBucketResourceGroupConfiguration(BucketResourceGroupConfiguration bucketResourceGroupConfiguration) {
        this.bucketResourceGroupConfiguration = bucketResourceGroupConfiguration;
        return this;
    }
    public BucketResourceGroupConfiguration getBucketResourceGroupConfiguration() {
        return this.bucketResourceGroupConfiguration;
    }

}
