// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class GetBucketArchiveDirectReadResponseBody extends TeaModel {
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
