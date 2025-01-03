// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class RequesterQoSInfo extends TeaModel {
    @NameInMap("QoSConfiguration")
    public QoSConfiguration qoSConfiguration;

    /**
     * <strong>example:</strong>
     * <p>21234567890090</p>
     */
    @NameInMap("Requester")
    public String requester;

    public static RequesterQoSInfo build(java.util.Map<String, ?> map) throws Exception {
        RequesterQoSInfo self = new RequesterQoSInfo();
        return TeaModel.build(map, self);
    }

    public RequesterQoSInfo setQoSConfiguration(QoSConfiguration qoSConfiguration) {
        this.qoSConfiguration = qoSConfiguration;
        return this;
    }
    public QoSConfiguration getQoSConfiguration() {
        return this.qoSConfiguration;
    }

    public RequesterQoSInfo setRequester(String requester) {
        this.requester = requester;
        return this;
    }
    public String getRequester() {
        return this.requester;
    }

}
