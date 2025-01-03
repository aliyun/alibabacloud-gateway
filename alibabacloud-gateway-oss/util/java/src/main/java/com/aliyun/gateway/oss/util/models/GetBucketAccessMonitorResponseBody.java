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

    public static class AccessMonitorConfiguration extends TeaModel {
        /**
         * <p>The access tracking status of the bucket. Valid values:</p>
         * <ul>
         * <li><p>Enabled: Access tracking is enabled.</p>
         * </li>
         * <li><p>Disabled: Access tracking is disabled.</p>
         * </li>
         * </ul>
         */
        @NameInMap("Status")
        public String status;

        public static AccessMonitorConfiguration build(java.util.Map<String, ?> map) throws Exception {
            AccessMonitorConfiguration self = new AccessMonitorConfiguration();
            return TeaModel.build(map, self);
        }

        public AccessMonitorConfiguration setStatus(String status) {
            this.status = status;
            return this;
        }
        public String getStatus() {
            return this.status;
        }

    }

}
