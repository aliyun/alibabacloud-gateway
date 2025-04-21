// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketResourcePoolBucketGroupRequest extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("resourcePool")
    public String resourcePool;

    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("resourcePoolBucketGroup")
    public String resourcePoolBucketGroup;

    public static PutBucketResourcePoolBucketGroupRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketResourcePoolBucketGroupRequest self = new PutBucketResourcePoolBucketGroupRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketResourcePoolBucketGroupRequest setResourcePool(String resourcePool) {
        this.resourcePool = resourcePool;
        return this;
    }
    public String getResourcePool() {
        return this.resourcePool;
    }

    public PutBucketResourcePoolBucketGroupRequest setResourcePoolBucketGroup(String resourcePoolBucketGroup) {
        this.resourcePoolBucketGroup = resourcePoolBucketGroup;
        return this;
    }
    public String getResourcePoolBucketGroup() {
        return this.resourcePoolBucketGroup;
    }

}
