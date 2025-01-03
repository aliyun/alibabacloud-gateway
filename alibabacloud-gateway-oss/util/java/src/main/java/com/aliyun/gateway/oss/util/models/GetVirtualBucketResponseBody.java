// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetVirtualBucketResponseBody extends TeaModel {
    @NameInMap("VirtualBucketConfiguration")
    public VirtualBucket virtualBucketConfiguration;

    public static GetVirtualBucketResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetVirtualBucketResponseBody self = new GetVirtualBucketResponseBody();
        return TeaModel.build(map, self);
    }

    public GetVirtualBucketResponseBody setVirtualBucketConfiguration(VirtualBucket virtualBucketConfiguration) {
        this.virtualBucketConfiguration = virtualBucketConfiguration;
        return this;
    }
    public VirtualBucket getVirtualBucketConfiguration() {
        return this.virtualBucketConfiguration;
    }

}
