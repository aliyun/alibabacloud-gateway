// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetCacheResponseBody extends TeaModel {
    @NameInMap("Cache")
    public CacheDetailInfo cache;

    public static GetCacheResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetCacheResponseBody self = new GetCacheResponseBody();
        return TeaModel.build(map, self);
    }

    public GetCacheResponseBody setCache(CacheDetailInfo cache) {
        this.cache = cache;
        return this;
    }
    public CacheDetailInfo getCache() {
        return this.cache;
    }

}
