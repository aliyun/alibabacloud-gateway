// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CacheBaseInfo extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>cn-hangzhou-j</p>
     */
    @NameInMap("AvailableZone")
    public String availableZone;

    /**
     * <strong>example:</strong>
     * <p>2023-09-12T15:26:29.000Z</p>
     */
    @NameInMap("CreationDate")
    public String creationDate;

    /**
     * <strong>example:</strong>
     * <p>cache1</p>
     */
    @NameInMap("Name")
    public String name;

    @NameInMap("QuotaConfiguration")
    public CacheQuotaConfiguration quotaConfiguration;

    public static CacheBaseInfo build(java.util.Map<String, ?> map) throws Exception {
        CacheBaseInfo self = new CacheBaseInfo();
        return TeaModel.build(map, self);
    }

    public CacheBaseInfo setAvailableZone(String availableZone) {
        this.availableZone = availableZone;
        return this;
    }
    public String getAvailableZone() {
        return this.availableZone;
    }

    public CacheBaseInfo setCreationDate(String creationDate) {
        this.creationDate = creationDate;
        return this;
    }
    public String getCreationDate() {
        return this.creationDate;
    }

    public CacheBaseInfo setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return this.name;
    }

    public CacheBaseInfo setQuotaConfiguration(CacheQuotaConfiguration quotaConfiguration) {
        this.quotaConfiguration = quotaConfiguration;
        return this;
    }
    public CacheQuotaConfiguration getQuotaConfiguration() {
        return this.quotaConfiguration;
    }

}
