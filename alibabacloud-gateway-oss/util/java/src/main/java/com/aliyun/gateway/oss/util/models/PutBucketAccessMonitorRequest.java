// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketAccessMonitorRequest extends TeaModel {
    /**
     * <p>The access tracking configurations of the bucket.</p>
     */
    @NameInMap("AccessMonitorConfiguration")
    public AccessMonitorConfiguration accessMonitorConfiguration;

    public static PutBucketAccessMonitorRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketAccessMonitorRequest self = new PutBucketAccessMonitorRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketAccessMonitorRequest setAccessMonitorConfiguration(AccessMonitorConfiguration accessMonitorConfiguration) {
        this.accessMonitorConfiguration = accessMonitorConfiguration;
        return this;
    }
    public AccessMonitorConfiguration getAccessMonitorConfiguration() {
        return this.accessMonitorConfiguration;
    }

}
