// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListVirtualBucketResult extends TeaModel {
    @NameInMap("VirtualBucket")
    public java.util.List<VirtualBucket> virtualBucket;

    public static ListVirtualBucketResult build(java.util.Map<String, ?> map) throws Exception {
        ListVirtualBucketResult self = new ListVirtualBucketResult();
        return TeaModel.build(map, self);
    }

    public ListVirtualBucketResult setVirtualBucket(java.util.List<VirtualBucket> virtualBucket) {
        this.virtualBucket = virtualBucket;
        return this;
    }
    public java.util.List<VirtualBucket> getVirtualBucket() {
        return this.virtualBucket;
    }

}
