// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketArchiveDirectReadResponseBody extends TeaModel {
    /**
     * <p>The container that stores the configurations for real-time access of Archive objects.</p>
     */
    @NameInMap("ArchiveDirectReadConfiguration")
    public ArchiveDirectReadConfiguration archiveDirectReadConfiguration;

    public static GetBucketArchiveDirectReadResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketArchiveDirectReadResponseBody self = new GetBucketArchiveDirectReadResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketArchiveDirectReadResponseBody setArchiveDirectReadConfiguration(ArchiveDirectReadConfiguration archiveDirectReadConfiguration) {
        this.archiveDirectReadConfiguration = archiveDirectReadConfiguration;
        return this;
    }
    public ArchiveDirectReadConfiguration getArchiveDirectReadConfiguration() {
        return this.archiveDirectReadConfiguration;
    }

}
