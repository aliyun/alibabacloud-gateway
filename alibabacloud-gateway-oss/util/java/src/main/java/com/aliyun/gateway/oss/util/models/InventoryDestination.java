// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class InventoryDestination extends TeaModel {
    @NameInMap("OSSBucketDestination")
    public InventoryOSSBucketDestination OSSBucketDestination;

    public static InventoryDestination build(java.util.Map<String, ?> map) throws Exception {
        InventoryDestination self = new InventoryDestination();
        return TeaModel.build(map, self);
    }

    public InventoryDestination setOSSBucketDestination(InventoryOSSBucketDestination OSSBucketDestination) {
        this.OSSBucketDestination = OSSBucketDestination;
        return this;
    }
    public InventoryOSSBucketDestination getOSSBucketDestination() {
        return this.OSSBucketDestination;
    }

}
