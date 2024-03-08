// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class PutBucketArchiveDirectReadRequest extends TeaModel {
    @NameInMap("ArchiveDirectReadConfiguration")
    public ArchiveDirectReadConfiguration archiveDirectReadConfiguration;

    public static PutBucketArchiveDirectReadRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketArchiveDirectReadRequest self = new PutBucketArchiveDirectReadRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketArchiveDirectReadRequest setArchiveDirectReadConfiguration(ArchiveDirectReadConfiguration archiveDirectReadConfiguration) {
        this.archiveDirectReadConfiguration = archiveDirectReadConfiguration;
        return this;
    }
    public ArchiveDirectReadConfiguration getArchiveDirectReadConfiguration() {
        return this.archiveDirectReadConfiguration;
    }

}
