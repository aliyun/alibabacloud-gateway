// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketEventNotificationResponseBody extends TeaModel {
    /**
     * <p>Function Compute service configuration for a bucket.</p>
     */
    @NameInMap("NotificationConfiguration")
    public EventNotificationConfiguration notificationConfiguration;

    public static GetBucketEventNotificationResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketEventNotificationResponseBody self = new GetBucketEventNotificationResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketEventNotificationResponseBody setNotificationConfiguration(EventNotificationConfiguration notificationConfiguration) {
        this.notificationConfiguration = notificationConfiguration;
        return this;
    }
    public EventNotificationConfiguration getNotificationConfiguration() {
        return this.notificationConfiguration;
    }

}
