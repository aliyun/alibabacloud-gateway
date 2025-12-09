// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketAccessMonitorResponseBody extends TeaModel {
    /**
     * <p>The container that stores access monitor configuration.</p>
     */
    @NameInMap("AccessMonitorConfiguration")
    public AccessMonitorConfiguration accessMonitorConfiguration;

    public static GetBucketAccessMonitorResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketAccessMonitorResponseBody self = new GetBucketAccessMonitorResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketAccessMonitorResponseBody setAccessMonitorConfiguration(AccessMonitorConfiguration accessMonitorConfiguration) {
        this.accessMonitorConfiguration = accessMonitorConfiguration;
        return this;
    }
    public AccessMonitorConfiguration getAccessMonitorConfiguration() {
        return this.accessMonitorConfiguration;
    }

}
