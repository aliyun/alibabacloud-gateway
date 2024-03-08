// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketInventoryRequest extends TeaModel {
    /**
     * <p>The container that stores the configurations of the inventory.</p>
     */
    @NameInMap("body")
    public InventoryConfiguration body;

    /**
     * <p>The name of the inventory.</p>
     */
    @NameInMap("inventoryId")
    public String inventoryId;

    public static PutBucketInventoryRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketInventoryRequest self = new PutBucketInventoryRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketInventoryRequest setBody(InventoryConfiguration body) {
        this.body = body;
        return this;
    }
    public InventoryConfiguration getBody() {
        return this.body;
    }

    public PutBucketInventoryRequest setInventoryId(String inventoryId) {
        this.inventoryId = inventoryId;
        return this;
    }
    public String getInventoryId() {
        return this.inventoryId;
    }

}
