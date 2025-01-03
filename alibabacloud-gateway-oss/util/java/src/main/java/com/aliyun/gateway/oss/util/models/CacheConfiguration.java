// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CacheConfiguration extends TeaModel {
    @NameInMap("Caches")
    public Caches caches;

    public static CacheConfiguration build(java.util.Map<String, ?> map) throws Exception {
        CacheConfiguration self = new CacheConfiguration();
        return TeaModel.build(map, self);
    }

    public CacheConfiguration setCaches(Caches caches) {
        this.caches = caches;
        return this;
    }
    public Caches getCaches() {
        return this.caches;
    }

    public static class Cache extends TeaModel {
        @NameInMap("AcceleratePaths")
        public AcceleratePaths acceleratePaths;

        /**
         * <strong>example:</strong>
         * <p>cn-hangzhou-j</p>
         */
        @NameInMap("AvailableZone")
        public String availableZone;

        /**
         * <strong>example:</strong>
         * <p>data-acc-test_data-acc</p>
         */
        @NameInMap("CacheName")
        public String cacheName;

        /**
         * <strong>example:</strong>
         * <p>sync-warmup</p>
         */
        @NameInMap("CachePolicy")
        public String cachePolicy;

        public static Cache build(java.util.Map<String, ?> map) throws Exception {
            Cache self = new Cache();
            return TeaModel.build(map, self);
        }

        public Cache setAcceleratePaths(AcceleratePaths acceleratePaths) {
            this.acceleratePaths = acceleratePaths;
            return this;
        }
        public AcceleratePaths getAcceleratePaths() {
            return this.acceleratePaths;
        }

        public Cache setAvailableZone(String availableZone) {
            this.availableZone = availableZone;
            return this;
        }
        public String getAvailableZone() {
            return this.availableZone;
        }

        public Cache setCacheName(String cacheName) {
            this.cacheName = cacheName;
            return this;
        }
        public String getCacheName() {
            return this.cacheName;
        }

        public Cache setCachePolicy(String cachePolicy) {
            this.cachePolicy = cachePolicy;
            return this;
        }
        public String getCachePolicy() {
            return this.cachePolicy;
        }

    }

    public static class Caches extends TeaModel {
        @NameInMap("Cache")
        public Cache cache;

        public static Caches build(java.util.Map<String, ?> map) throws Exception {
            Caches self = new Caches();
            return TeaModel.build(map, self);
        }

        public Caches setCache(Cache cache) {
            this.cache = cache;
            return this;
        }
        public Cache getCache() {
            return this.cache;
        }

    }

}
