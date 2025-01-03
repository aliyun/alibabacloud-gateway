// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class VirtualBucket extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>virtual-bucket</p>
     */
    @NameInMap("Name")
    public String name;

    @NameInMap("RealBucket")
    public java.util.List<RealBucket> realBucket;

    public static VirtualBucket build(java.util.Map<String, ?> map) throws Exception {
        VirtualBucket self = new VirtualBucket();
        return TeaModel.build(map, self);
    }

    public VirtualBucket setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return this.name;
    }

    public VirtualBucket setRealBucket(java.util.List<RealBucket> realBucket) {
        this.realBucket = realBucket;
        return this;
    }
    public java.util.List<RealBucket> getRealBucket() {
        return this.realBucket;
    }

    public static class RealBucket extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>real-bucket</p>
         */
        @NameInMap("Name")
        public String name;

        /**
         * <strong>example:</strong>
         * <p>ok</p>
         */
        @NameInMap("Status")
        public String status;

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

        public RealBucket setStatus(String status) {
            this.status = status;
            return this;
        }
        public String getStatus() {
            return this.status;
        }

    }

}
