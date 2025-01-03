// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetResourcePoolRequesterQoSInfoRequest extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>26753xxxxxxxx14340</p>
     */
    @NameInMap("qosRequester")
    public String qosRequester;

    /**
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>rp-01</p>
     */
    @NameInMap("resourcePool")
    public String resourcePool;

    public static GetResourcePoolRequesterQoSInfoRequest build(java.util.Map<String, ?> map) throws Exception {
        GetResourcePoolRequesterQoSInfoRequest self = new GetResourcePoolRequesterQoSInfoRequest();
        return TeaModel.build(map, self);
    }

    public GetResourcePoolRequesterQoSInfoRequest setQosRequester(String qosRequester) {
        this.qosRequester = qosRequester;
        return this;
    }
    public String getQosRequester() {
        return this.qosRequester;
    }

    public GetResourcePoolRequesterQoSInfoRequest setResourcePool(String resourcePool) {
        this.resourcePool = resourcePool;
        return this;
    }
    public String getResourcePool() {
        return this.resourcePool;
    }

}
