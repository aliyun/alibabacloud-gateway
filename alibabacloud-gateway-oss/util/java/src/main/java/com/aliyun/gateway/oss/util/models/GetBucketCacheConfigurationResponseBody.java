// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketCacheConfigurationResponseBody extends TeaModel {
    @NameInMap("CacheConfiguration")
    public CacheConfiguration cacheConfiguration;

    public static GetBucketCacheConfigurationResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketCacheConfigurationResponseBody self = new GetBucketCacheConfigurationResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketCacheConfigurationResponseBody setCacheConfiguration(CacheConfiguration cacheConfiguration) {
        this.cacheConfiguration = cacheConfiguration;
        return this;
    }
    public CacheConfiguration getCacheConfiguration() {
        return this.cacheConfiguration;
    }

}
