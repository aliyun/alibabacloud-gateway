// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketRequesterQoSInfoRequest extends TeaModel {
    @NameInMap("QoSConfiguration")
    public QoSConfiguration qoSConfiguration;

    /**
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>26753xxxxxxxx14340</p>
     */
    @NameInMap("qosRequester")
    public String qosRequester;

    public static PutBucketRequesterQoSInfoRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketRequesterQoSInfoRequest self = new PutBucketRequesterQoSInfoRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketRequesterQoSInfoRequest setQoSConfiguration(QoSConfiguration qoSConfiguration) {
        this.qoSConfiguration = qoSConfiguration;
        return this;
    }
    public QoSConfiguration getQoSConfiguration() {
        return this.qoSConfiguration;
    }

    public PutBucketRequesterQoSInfoRequest setQosRequester(String qosRequester) {
        this.qosRequester = qosRequester;
        return this;
    }
    public String getQosRequester() {
        return this.qosRequester;
    }

}
