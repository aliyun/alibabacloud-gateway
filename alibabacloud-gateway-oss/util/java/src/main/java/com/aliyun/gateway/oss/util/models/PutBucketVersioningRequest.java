// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketVersioningRequest extends TeaModel {
    /**
     * <p>The container that stores the versioning state of the bucket.</p>
     */
    @NameInMap("VersioningConfiguration")
    public VersioningConfiguration versioningConfiguration;

    public static PutBucketVersioningRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketVersioningRequest self = new PutBucketVersioningRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketVersioningRequest setVersioningConfiguration(VersioningConfiguration versioningConfiguration) {
        this.versioningConfiguration = versioningConfiguration;
        return this;
    }
    public VersioningConfiguration getVersioningConfiguration() {
        return this.versioningConfiguration;
    }

}
