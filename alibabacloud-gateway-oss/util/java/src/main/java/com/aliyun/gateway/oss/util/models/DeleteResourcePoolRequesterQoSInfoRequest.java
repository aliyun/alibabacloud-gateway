// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DeleteResourcePoolRequesterQoSInfoRequest extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("qosRequester")
    public String qosRequester;

    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("resourcePool")
    public String resourcePool;

    public static DeleteResourcePoolRequesterQoSInfoRequest build(java.util.Map<String, ?> map) throws Exception {
        DeleteResourcePoolRequesterQoSInfoRequest self = new DeleteResourcePoolRequesterQoSInfoRequest();
        return TeaModel.build(map, self);
    }

    public DeleteResourcePoolRequesterQoSInfoRequest setQosRequester(String qosRequester) {
        this.qosRequester = qosRequester;
        return this;
    }
    public String getQosRequester() {
        return this.qosRequester;
    }

    public DeleteResourcePoolRequesterQoSInfoRequest setResourcePool(String resourcePool) {
        this.resourcePool = resourcePool;
        return this;
    }
    public String getResourcePool() {
        return this.resourcePool;
    }

}
