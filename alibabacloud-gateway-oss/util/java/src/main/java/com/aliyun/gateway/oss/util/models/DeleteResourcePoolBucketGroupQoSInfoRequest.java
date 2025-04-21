// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DeleteResourcePoolBucketGroupQoSInfoRequest extends TeaModel {
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

    public static DeleteResourcePoolBucketGroupQoSInfoRequest build(java.util.Map<String, ?> map) throws Exception {
        DeleteResourcePoolBucketGroupQoSInfoRequest self = new DeleteResourcePoolBucketGroupQoSInfoRequest();
        return TeaModel.build(map, self);
    }

    public DeleteResourcePoolBucketGroupQoSInfoRequest setResourcePool(String resourcePool) {
        this.resourcePool = resourcePool;
        return this;
    }
    public String getResourcePool() {
        return this.resourcePool;
    }

    public DeleteResourcePoolBucketGroupQoSInfoRequest setResourcePoolBucketGroup(String resourcePoolBucketGroup) {
        this.resourcePoolBucketGroup = resourcePoolBucketGroup;
        return this;
    }
    public String getResourcePoolBucketGroup() {
        return this.resourcePoolBucketGroup;
    }

}
