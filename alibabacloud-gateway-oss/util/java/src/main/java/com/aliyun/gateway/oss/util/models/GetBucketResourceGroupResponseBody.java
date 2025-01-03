// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketResourceGroupResponseBody extends TeaModel {
    /**
     * <p>The container that stores the ID of the resource group.</p>
     */
    @NameInMap("BucketResourceGroupConfiguration")
    public BucketResourceGroupConfiguration bucketResourceGroupConfiguration;

    public static GetBucketResourceGroupResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketResourceGroupResponseBody self = new GetBucketResourceGroupResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketResourceGroupResponseBody setBucketResourceGroupConfiguration(BucketResourceGroupConfiguration bucketResourceGroupConfiguration) {
        this.bucketResourceGroupConfiguration = bucketResourceGroupConfiguration;
        return this;
    }
    public BucketResourceGroupConfiguration getBucketResourceGroupConfiguration() {
        return this.bucketResourceGroupConfiguration;
    }

    public static class BucketResourceGroupConfiguration extends TeaModel {
        /**
         * <p>The ID of the resource group to which the bucket belongs.</p>
         * <p>If this element is not specified, the bucket is moved to the default resource group.</p>
         * 
         * <strong>example:</strong>
         * <p>rg-asdfklj***</p>
         */
        @NameInMap("ResourceGroupId")
        public String resourceGroupId;

        public static BucketResourceGroupConfiguration build(java.util.Map<String, ?> map) throws Exception {
            BucketResourceGroupConfiguration self = new BucketResourceGroupConfiguration();
            return TeaModel.build(map, self);
        }

        public BucketResourceGroupConfiguration setResourceGroupId(String resourceGroupId) {
            this.resourceGroupId = resourceGroupId;
            return this;
        }
        public String getResourceGroupId() {
            return this.resourceGroupId;
        }

    }

}
