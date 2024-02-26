// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class GetBucketLoggingResponseBody extends TeaModel {
    /**
     * <p>Indicates the container used to store access logging configuration of a bucket.</p>
     */
    @NameInMap("BucketLoggingStatus")
    public BucketLoggingStatus bucketLoggingStatus;

    public static GetBucketLoggingResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketLoggingResponseBody self = new GetBucketLoggingResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketLoggingResponseBody setBucketLoggingStatus(BucketLoggingStatus bucketLoggingStatus) {
        this.bucketLoggingStatus = bucketLoggingStatus;
        return this;
    }
    public BucketLoggingStatus getBucketLoggingStatus() {
        return this.bucketLoggingStatus;
    }

    public static class BucketLoggingStatus extends TeaModel {
        /**
         * <p>Indicates the container used to store access logging information. This element is returned if it is enabled and is not returned if it is disabled.</p>
         */
        @NameInMap("LoggingEnabled")
        public LoggingEnabled loggingEnabled;

        public static BucketLoggingStatus build(java.util.Map<String, ?> map) throws Exception {
            BucketLoggingStatus self = new BucketLoggingStatus();
            return TeaModel.build(map, self);
        }

        public BucketLoggingStatus setLoggingEnabled(LoggingEnabled loggingEnabled) {
            this.loggingEnabled = loggingEnabled;
            return this;
        }
        public LoggingEnabled getLoggingEnabled() {
            return this.loggingEnabled;
        }

    }

}
