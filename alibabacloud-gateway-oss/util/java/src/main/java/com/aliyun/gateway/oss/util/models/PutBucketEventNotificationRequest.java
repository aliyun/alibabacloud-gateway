// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketEventNotificationRequest extends TeaModel {
    @NameInMap("NotificationConfiguration")
    public EventNotificationConfiguration notificationConfiguration;

    public static PutBucketEventNotificationRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketEventNotificationRequest self = new PutBucketEventNotificationRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketEventNotificationRequest setNotificationConfiguration(EventNotificationConfiguration notificationConfiguration) {
        this.notificationConfiguration = notificationConfiguration;
        return this;
    }
    public EventNotificationConfiguration getNotificationConfiguration() {
        return this.notificationConfiguration;
    }

}
