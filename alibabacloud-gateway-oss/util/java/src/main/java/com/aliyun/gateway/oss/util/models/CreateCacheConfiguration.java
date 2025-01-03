// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateCacheConfiguration extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>cn-hangzhou-j</p>
     */
    @NameInMap("AvailableZone")
    public String availableZone;

    /**
     * <strong>example:</strong>
     * <p>cache2</p>
     */
    @NameInMap("Name")
    public String name;

    @NameInMap("QuotaConfiguration")
    public CacheQuotaConfiguration quotaConfiguration;

    public static CreateCacheConfiguration build(java.util.Map<String, ?> map) throws Exception {
        CreateCacheConfiguration self = new CreateCacheConfiguration();
        return TeaModel.build(map, self);
    }

    public CreateCacheConfiguration setAvailableZone(String availableZone) {
        this.availableZone = availableZone;
        return this;
    }
    public String getAvailableZone() {
        return this.availableZone;
    }

    public CreateCacheConfiguration setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return this.name;
    }

    public CreateCacheConfiguration setQuotaConfiguration(CacheQuotaConfiguration quotaConfiguration) {
        this.quotaConfiguration = quotaConfiguration;
        return this;
    }
    public CacheQuotaConfiguration getQuotaConfiguration() {
        return this.quotaConfiguration;
    }

}
