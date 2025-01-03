// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketQoSInfoRequest extends TeaModel {
    @NameInMap("QoSConfiguration")
    public QoSConfiguration qoSConfiguration;

    public static PutBucketQoSInfoRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketQoSInfoRequest self = new PutBucketQoSInfoRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketQoSInfoRequest setQoSConfiguration(QoSConfiguration qoSConfiguration) {
        this.qoSConfiguration = qoSConfiguration;
        return this;
    }
    public QoSConfiguration getQoSConfiguration() {
        return this.qoSConfiguration;
    }

}
