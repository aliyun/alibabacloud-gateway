// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketInventoryResponseBody extends TeaModel {
    /**
     * <p>The inventory task configured for a bucket.</p>
     */
    @NameInMap("InventoryConfiguration")
    public InventoryConfiguration inventoryConfiguration;

    public static GetBucketInventoryResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketInventoryResponseBody self = new GetBucketInventoryResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketInventoryResponseBody setInventoryConfiguration(InventoryConfiguration inventoryConfiguration) {
        this.inventoryConfiguration = inventoryConfiguration;
        return this;
    }
    public InventoryConfiguration getInventoryConfiguration() {
        return this.inventoryConfiguration;
    }

}
