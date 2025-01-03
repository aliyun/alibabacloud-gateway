// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CacheDetailInfo extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>cn-hangzhou-j</p>
     */
    @NameInMap("AvailableZone")
    public String availableZone;

    @NameInMap("Buckets")
    public Buckets buckets;

    /**
     * <strong>example:</strong>
     * <p>2023-09-12T15:26:29.000Z</p>
     */
    @NameInMap("CreationDate")
    public String creationDate;

    /**
     * <strong>example:</strong>
     * <p>test-acc</p>
     */
    @NameInMap("Name")
    public String name;

    @NameInMap("QuotaConfiguration")
    public CacheQuotaConfiguration quotaConfiguration;

    public static CacheDetailInfo build(java.util.Map<String, ?> map) throws Exception {
        CacheDetailInfo self = new CacheDetailInfo();
        return TeaModel.build(map, self);
    }

    public CacheDetailInfo setAvailableZone(String availableZone) {
        this.availableZone = availableZone;
        return this;
    }
    public String getAvailableZone() {
        return this.availableZone;
    }

    public CacheDetailInfo setBuckets(Buckets buckets) {
        this.buckets = buckets;
        return this;
    }
    public Buckets getBuckets() {
        return this.buckets;
    }

    public CacheDetailInfo setCreationDate(String creationDate) {
        this.creationDate = creationDate;
        return this;
    }
    public String getCreationDate() {
        return this.creationDate;
    }

    public CacheDetailInfo setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return this.name;
    }

    public CacheDetailInfo setQuotaConfiguration(CacheQuotaConfiguration quotaConfiguration) {
        this.quotaConfiguration = quotaConfiguration;
        return this;
    }
    public CacheQuotaConfiguration getQuotaConfiguration() {
        return this.quotaConfiguration;
    }

    public static class Buckets extends TeaModel {
        @NameInMap("Bucket")
        public java.util.List<CacheBucketInfo> bucket;

        public static Buckets build(java.util.Map<String, ?> map) throws Exception {
            Buckets self = new Buckets();
            return TeaModel.build(map, self);
        }

        public Buckets setBucket(java.util.List<CacheBucketInfo> bucket) {
            this.bucket = bucket;
            return this;
        }
        public java.util.List<CacheBucketInfo> getBucket() {
            return this.bucket;
        }

    }

}
