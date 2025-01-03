// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketPublicAccessBlockResponseBody extends TeaModel {
    /**
     * <p>The container in which the Block Public Access configurations are stored.</p>
     */
    @NameInMap("PublicAccessBlockConfiguration")
    public PublicAccessBlockConfiguration publicAccessBlockConfiguration;

    public static GetBucketPublicAccessBlockResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketPublicAccessBlockResponseBody self = new GetBucketPublicAccessBlockResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketPublicAccessBlockResponseBody setPublicAccessBlockConfiguration(PublicAccessBlockConfiguration publicAccessBlockConfiguration) {
        this.publicAccessBlockConfiguration = publicAccessBlockConfiguration;
        return this;
    }
    public PublicAccessBlockConfiguration getPublicAccessBlockConfiguration() {
        return this.publicAccessBlockConfiguration;
    }

}
