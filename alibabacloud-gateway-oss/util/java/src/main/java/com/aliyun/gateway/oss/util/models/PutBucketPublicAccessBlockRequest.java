// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketPublicAccessBlockRequest extends TeaModel {
    @NameInMap("PublicAccessBlockConfiguration")
    public PublicAccessBlockConfiguration publicAccessBlockConfiguration;

    public static PutBucketPublicAccessBlockRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketPublicAccessBlockRequest self = new PutBucketPublicAccessBlockRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketPublicAccessBlockRequest setPublicAccessBlockConfiguration(PublicAccessBlockConfiguration publicAccessBlockConfiguration) {
        this.publicAccessBlockConfiguration = publicAccessBlockConfiguration;
        return this;
    }
    public PublicAccessBlockConfiguration getPublicAccessBlockConfiguration() {
        return this.publicAccessBlockConfiguration;
    }

}
