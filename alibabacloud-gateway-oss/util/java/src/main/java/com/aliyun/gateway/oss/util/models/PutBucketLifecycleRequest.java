// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class PutBucketLifecycleRequest extends TeaModel {
    /**
     * <p>The container that stores the lifecycle configuration.</p>
     */
    @NameInMap("LifecycleConfiguration")
    public LifecycleConfiguration lifecycleConfiguration;

    public static PutBucketLifecycleRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketLifecycleRequest self = new PutBucketLifecycleRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketLifecycleRequest setLifecycleConfiguration(LifecycleConfiguration lifecycleConfiguration) {
        this.lifecycleConfiguration = lifecycleConfiguration;
        return this;
    }
    public LifecycleConfiguration getLifecycleConfiguration() {
        return this.lifecycleConfiguration;
    }

}
