// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CacheBucketInfo extends TeaModel {
    @NameInMap("AcceleratePaths")
    public AcceleratePaths acceleratePaths;

    /**
     * <strong>example:</strong>
     * <p>sync-warmup</p>
     */
    @NameInMap("CachePolicy")
    public String cachePolicy;

    /**
     * <strong>example:</strong>
     * <p>test-bucket</p>
     */
    @NameInMap("Name")
    public String name;

    public static CacheBucketInfo build(java.util.Map<String, ?> map) throws Exception {
        CacheBucketInfo self = new CacheBucketInfo();
        return TeaModel.build(map, self);
    }

    public CacheBucketInfo setAcceleratePaths(AcceleratePaths acceleratePaths) {
        this.acceleratePaths = acceleratePaths;
        return this;
    }
    public AcceleratePaths getAcceleratePaths() {
        return this.acceleratePaths;
    }

    public CacheBucketInfo setCachePolicy(String cachePolicy) {
        this.cachePolicy = cachePolicy;
        return this;
    }
    public String getCachePolicy() {
        return this.cachePolicy;
    }

    public CacheBucketInfo setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return this.name;
    }

}
