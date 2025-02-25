// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DeleteBucketInventoryRequest extends TeaModel {
    /**
     * <p>The name of the inventory that you want to delete.</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>list1</p>
     */
    @NameInMap("inventoryId")
    public String inventoryId;

    public static DeleteBucketInventoryRequest build(java.util.Map<String, ?> map) throws Exception {
        DeleteBucketInventoryRequest self = new DeleteBucketInventoryRequest();
        return TeaModel.build(map, self);
    }

    public DeleteBucketInventoryRequest setInventoryId(String inventoryId) {
        this.inventoryId = inventoryId;
        return this;
    }
    public String getInventoryId() {
        return this.inventoryId;
    }

}
