// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutVirtualBucketRequest extends TeaModel {
    @NameInMap("VirtualBucketConfiguration")
    public VirtualBucketConfiguration virtualBucketConfiguration;

    public static PutVirtualBucketRequest build(java.util.Map<String, ?> map) throws Exception {
        PutVirtualBucketRequest self = new PutVirtualBucketRequest();
        return TeaModel.build(map, self);
    }

    public PutVirtualBucketRequest setVirtualBucketConfiguration(VirtualBucketConfiguration virtualBucketConfiguration) {
        this.virtualBucketConfiguration = virtualBucketConfiguration;
        return this;
    }
    public VirtualBucketConfiguration getVirtualBucketConfiguration() {
        return this.virtualBucketConfiguration;
    }

}
