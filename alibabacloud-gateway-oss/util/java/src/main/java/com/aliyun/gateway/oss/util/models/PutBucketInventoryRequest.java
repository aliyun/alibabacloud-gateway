// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketInventoryRequest extends TeaModel {
    /**
     * <p>The container that stores the Inventory configuration.</p>
     */
    @NameInMap("InventoryConfiguration")
    public InventoryConfiguration inventoryConfiguration;

    /**
     * <p>The name of the inventory.</p>
     */
    @NameInMap("inventoryId")
    public String inventoryId;

    public static PutBucketInventoryRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketInventoryRequest self = new PutBucketInventoryRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketInventoryRequest setInventoryConfiguration(InventoryConfiguration inventoryConfiguration) {
        this.inventoryConfiguration = inventoryConfiguration;
        return this;
    }
    public InventoryConfiguration getInventoryConfiguration() {
        return this.inventoryConfiguration;
    }

    public PutBucketInventoryRequest setInventoryId(String inventoryId) {
        this.inventoryId = inventoryId;
        return this;
    }
    public String getInventoryId() {
        return this.inventoryId;
    }

}
