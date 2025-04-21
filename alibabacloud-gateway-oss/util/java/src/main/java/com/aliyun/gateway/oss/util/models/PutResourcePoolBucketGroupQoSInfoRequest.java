// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutResourcePoolBucketGroupQoSInfoRequest extends TeaModel {
    @NameInMap("ResourcePoolBucketGroupQoSInfo")
    public ResourcePoolBucketGroupQoSInfo resourcePoolBucketGroupQoSInfo;

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

    public static PutResourcePoolBucketGroupQoSInfoRequest build(java.util.Map<String, ?> map) throws Exception {
        PutResourcePoolBucketGroupQoSInfoRequest self = new PutResourcePoolBucketGroupQoSInfoRequest();
        return TeaModel.build(map, self);
    }

    public PutResourcePoolBucketGroupQoSInfoRequest setResourcePoolBucketGroupQoSInfo(ResourcePoolBucketGroupQoSInfo resourcePoolBucketGroupQoSInfo) {
        this.resourcePoolBucketGroupQoSInfo = resourcePoolBucketGroupQoSInfo;
        return this;
    }
    public ResourcePoolBucketGroupQoSInfo getResourcePoolBucketGroupQoSInfo() {
        return this.resourcePoolBucketGroupQoSInfo;
    }

    public PutResourcePoolBucketGroupQoSInfoRequest setResourcePool(String resourcePool) {
        this.resourcePool = resourcePool;
        return this;
    }
    public String getResourcePool() {
        return this.resourcePool;
    }

    public PutResourcePoolBucketGroupQoSInfoRequest setResourcePoolBucketGroup(String resourcePoolBucketGroup) {
        this.resourcePoolBucketGroup = resourcePoolBucketGroup;
        return this;
    }
    public String getResourcePoolBucketGroup() {
        return this.resourcePoolBucketGroup;
    }

}
