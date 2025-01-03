// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetResourcePoolInfoRequest extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>rp-01</p>
     */
    @NameInMap("resourcePool")
    public String resourcePool;

    public static GetResourcePoolInfoRequest build(java.util.Map<String, ?> map) throws Exception {
        GetResourcePoolInfoRequest self = new GetResourcePoolInfoRequest();
        return TeaModel.build(map, self);
    }

    public GetResourcePoolInfoRequest setResourcePool(String resourcePool) {
        this.resourcePool = resourcePool;
        return this;
    }
    public String getResourcePool() {
        return this.resourcePool;
    }

}
