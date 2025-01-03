// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketNotificationRequest extends TeaModel {
    @NameInMap("NotificationConfiguration")
    public NotificationConfiguration notificationConfiguration;

    public static PutBucketNotificationRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketNotificationRequest self = new PutBucketNotificationRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketNotificationRequest setNotificationConfiguration(NotificationConfiguration notificationConfiguration) {
        this.notificationConfiguration = notificationConfiguration;
        return this;
    }
    public NotificationConfiguration getNotificationConfiguration() {
        return this.notificationConfiguration;
    }

}
