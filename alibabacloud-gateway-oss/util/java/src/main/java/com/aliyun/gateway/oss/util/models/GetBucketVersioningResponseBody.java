// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketVersioningResponseBody extends TeaModel {
    /**
     * <p>The container that stores the versioning state of the bucket.</p>
     */
    @NameInMap("VersioningConfiguration")
    public VersioningConfiguration versioningConfiguration;

    public static GetBucketVersioningResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketVersioningResponseBody self = new GetBucketVersioningResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketVersioningResponseBody setVersioningConfiguration(VersioningConfiguration versioningConfiguration) {
        this.versioningConfiguration = versioningConfiguration;
        return this;
    }
    public VersioningConfiguration getVersioningConfiguration() {
        return this.versioningConfiguration;
    }

    public static class VersioningConfiguration extends TeaModel {
        /**
         * <p>The versioning state of the bucket.</p>
         */
        @NameInMap("Status")
        public String status;

        public static VersioningConfiguration build(java.util.Map<String, ?> map) throws Exception {
            VersioningConfiguration self = new VersioningConfiguration();
            return TeaModel.build(map, self);
        }

        public VersioningConfiguration setStatus(String status) {
            this.status = status;
            return this;
        }
        public String getStatus() {
            return this.status;
        }

    }

}
