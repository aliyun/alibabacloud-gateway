// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutPublicAccessBlockRequest extends TeaModel {
    /**
     * <p>The container in which the Block Public Access configurations are stored.</p>
     */
    @NameInMap("PublicAccessBlockConfiguration")
    public PublicAccessBlockConfiguration publicAccessBlockConfiguration;

    public static PutPublicAccessBlockRequest build(java.util.Map<String, ?> map) throws Exception {
        PutPublicAccessBlockRequest self = new PutPublicAccessBlockRequest();
        return TeaModel.build(map, self);
    }

    public PutPublicAccessBlockRequest setPublicAccessBlockConfiguration(PublicAccessBlockConfiguration publicAccessBlockConfiguration) {
        this.publicAccessBlockConfiguration = publicAccessBlockConfiguration;
        return this;
    }
    public PublicAccessBlockConfiguration getPublicAccessBlockConfiguration() {
        return this.publicAccessBlockConfiguration;
    }

}
