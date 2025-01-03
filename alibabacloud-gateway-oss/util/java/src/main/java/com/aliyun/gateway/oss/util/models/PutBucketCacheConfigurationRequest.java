// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketCacheConfigurationRequest extends TeaModel {
    @NameInMap("CacheConfiguration")
    public CacheConfiguration cacheConfiguration;

    public static PutBucketCacheConfigurationRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketCacheConfigurationRequest self = new PutBucketCacheConfigurationRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketCacheConfigurationRequest setCacheConfiguration(CacheConfiguration cacheConfiguration) {
        this.cacheConfiguration = cacheConfiguration;
        return this;
    }
    public CacheConfiguration getCacheConfiguration() {
        return this.cacheConfiguration;
    }

}
