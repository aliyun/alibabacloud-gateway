// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class VirtualBucketConfiguration extends TeaModel {
    @NameInMap("RealBucket")
    public java.util.List<RealBucket> realBucket;

    public static VirtualBucketConfiguration build(java.util.Map<String, ?> map) throws Exception {
        VirtualBucketConfiguration self = new VirtualBucketConfiguration();
        return TeaModel.build(map, self);
    }

    public VirtualBucketConfiguration setRealBucket(java.util.List<RealBucket> realBucket) {
        this.realBucket = realBucket;
        return this;
    }
    public java.util.List<RealBucket> getRealBucket() {
        return this.realBucket;
    }

    public static class RealBucket extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>realbucket-1</p>
         */
        @NameInMap("Name")
        public String name;

        public static RealBucket build(java.util.Map<String, ?> map) throws Exception {
            RealBucket self = new RealBucket();
            return TeaModel.build(map, self);
        }

        public RealBucket setName(String name) {
            this.name = name;
            return this;
        }
        public String getName() {
            return this.name;
        }

    }

}
