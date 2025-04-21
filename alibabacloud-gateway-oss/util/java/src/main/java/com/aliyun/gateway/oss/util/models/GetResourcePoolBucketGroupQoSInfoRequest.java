// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetResourcePoolBucketGroupQoSInfoRequest extends TeaModel {
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

    public static GetResourcePoolBucketGroupQoSInfoRequest build(java.util.Map<String, ?> map) throws Exception {
        GetResourcePoolBucketGroupQoSInfoRequest self = new GetResourcePoolBucketGroupQoSInfoRequest();
        return TeaModel.build(map, self);
    }

    public GetResourcePoolBucketGroupQoSInfoRequest setResourcePool(String resourcePool) {
        this.resourcePool = resourcePool;
        return this;
    }
    public String getResourcePool() {
        return this.resourcePool;
    }

    public GetResourcePoolBucketGroupQoSInfoRequest setResourcePoolBucketGroup(String resourcePoolBucketGroup) {
        this.resourcePoolBucketGroup = resourcePoolBucketGroup;
        return this;
    }
    public String getResourcePoolBucketGroup() {
        return this.resourcePoolBucketGroup;
    }

}
