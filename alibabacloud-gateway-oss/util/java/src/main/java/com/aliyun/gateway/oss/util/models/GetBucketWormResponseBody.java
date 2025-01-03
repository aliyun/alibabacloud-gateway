// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketWormResponseBody extends TeaModel {
    /**
     * <p>The container that stores the information about retention policies of the bucket.</p>
     */
    @NameInMap("WormConfiguration")
    public WormConfiguration wormConfiguration;

    public static GetBucketWormResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketWormResponseBody self = new GetBucketWormResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketWormResponseBody setWormConfiguration(WormConfiguration wormConfiguration) {
        this.wormConfiguration = wormConfiguration;
        return this;
    }
    public WormConfiguration getWormConfiguration() {
        return this.wormConfiguration;
    }

    public static class WormConfiguration extends TeaModel {
        /**
         * <p>The time at which the retention policy was created.</p>
         * 
         * <strong>example:</strong>
         * <p>2020-10-15T15:50:32</p>
         */
        @NameInMap("CreationDate")
        public String creationDate;

        /**
         * <p>The time at which the retention policy will be expired.</p>
         * 
         * <strong>example:</strong>
         * <p>2020-10-16T15:50:32</p>
         */
        @NameInMap("ExpirationDate")
        public String expirationDate;

        /**
         * <p>The number of days for which objects can be retained.</p>
         * 
         * <strong>example:</strong>
         * <p>20</p>
         */
        @NameInMap("RetentionPeriodInDays")
        public Integer retentionPeriodInDays;

        /**
         * <p>The status of the retention policy. Valid values:</p>
         * <ul>
         * <li>InProgress: indicates that the retention policy is in the InProgress state. By default, a retention policy is in the InProgress state after it is created. The policy remains in this state for 24 hours.</li>
         * <li>Locked: indicates that the retention policy is in the Locked state.</li>
         * </ul>
         */
        @NameInMap("State")
        public String state;

        /**
         * <p>The ID of the retention policy.</p>
         * <blockquote>
         * <p>Note If the specified retention policy ID that is used to query the retention policy configurations of the bucket does not exist, OSS returns the 404 error code.</p>
         * </blockquote>
         * 
         * <strong>example:</strong>
         * <p>1666E2CFB2B3418****</p>
         */
        @NameInMap("WormId")
        public String wormId;

        public static WormConfiguration build(java.util.Map<String, ?> map) throws Exception {
            WormConfiguration self = new WormConfiguration();
            return TeaModel.build(map, self);
        }

        public WormConfiguration setCreationDate(String creationDate) {
            this.creationDate = creationDate;
            return this;
        }
        public String getCreationDate() {
            return this.creationDate;
        }

        public WormConfiguration setExpirationDate(String expirationDate) {
            this.expirationDate = expirationDate;
            return this;
        }
        public String getExpirationDate() {
            return this.expirationDate;
        }

        public WormConfiguration setRetentionPeriodInDays(Integer retentionPeriodInDays) {
            this.retentionPeriodInDays = retentionPeriodInDays;
            return this;
        }
        public Integer getRetentionPeriodInDays() {
            return this.retentionPeriodInDays;
        }

        public WormConfiguration setState(String state) {
            this.state = state;
            return this;
        }
        public String getState() {
            return this.state;
        }

        public WormConfiguration setWormId(String wormId) {
            this.wormId = wormId;
            return this;
        }
        public String getWormId() {
            return this.wormId;
        }

    }

}
