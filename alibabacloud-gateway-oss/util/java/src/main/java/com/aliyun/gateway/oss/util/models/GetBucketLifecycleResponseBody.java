// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketLifecycleResponseBody extends TeaModel {
    /**
     * <p>The container that stores the lifecycle rules configured for the bucket.</p>
     */
    @NameInMap("LifecycleConfiguration")
    public LifecycleConfiguration lifecycleConfiguration;

    public static GetBucketLifecycleResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketLifecycleResponseBody self = new GetBucketLifecycleResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketLifecycleResponseBody setLifecycleConfiguration(LifecycleConfiguration lifecycleConfiguration) {
        this.lifecycleConfiguration = lifecycleConfiguration;
        return this;
    }
    public LifecycleConfiguration getLifecycleConfiguration() {
        return this.lifecycleConfiguration;
    }

}
