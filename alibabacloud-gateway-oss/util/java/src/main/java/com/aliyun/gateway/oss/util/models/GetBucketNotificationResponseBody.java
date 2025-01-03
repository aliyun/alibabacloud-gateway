// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketNotificationResponseBody extends TeaModel {
    @NameInMap("NotificationConfiguration")
    public NotificationConfiguration notificationConfiguration;

    public static GetBucketNotificationResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketNotificationResponseBody self = new GetBucketNotificationResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketNotificationResponseBody setNotificationConfiguration(NotificationConfiguration notificationConfiguration) {
        this.notificationConfiguration = notificationConfiguration;
        return this;
    }
    public NotificationConfiguration getNotificationConfiguration() {
        return this.notificationConfiguration;
    }

}
