// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutResourcePoolRequesterQoSInfoRequest extends TeaModel {
    @NameInMap("QoSConfiguration")
    public QoSConfiguration qoSConfiguration;

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

    public static PutResourcePoolRequesterQoSInfoRequest build(java.util.Map<String, ?> map) throws Exception {
        PutResourcePoolRequesterQoSInfoRequest self = new PutResourcePoolRequesterQoSInfoRequest();
        return TeaModel.build(map, self);
    }

    public PutResourcePoolRequesterQoSInfoRequest setQoSConfiguration(QoSConfiguration qoSConfiguration) {
        this.qoSConfiguration = qoSConfiguration;
        return this;
    }
    public QoSConfiguration getQoSConfiguration() {
        return this.qoSConfiguration;
    }

    public PutResourcePoolRequesterQoSInfoRequest setQosRequester(String qosRequester) {
        this.qosRequester = qosRequester;
        return this;
    }
    public String getQosRequester() {
        return this.qosRequester;
    }

    public PutResourcePoolRequesterQoSInfoRequest setResourcePool(String resourcePool) {
        this.resourcePool = resourcePool;
        return this;
    }
    public String getResourcePool() {
        return this.resourcePool;
    }

}
