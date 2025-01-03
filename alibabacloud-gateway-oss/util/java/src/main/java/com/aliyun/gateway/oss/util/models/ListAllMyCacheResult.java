// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListAllMyCacheResult extends TeaModel {
    @NameInMap("Caches")
    public Caches caches;

    /**
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("IsTruncated")
    public Boolean isTruncated;

    /**
     * <strong>example:</strong>
     * <p>abc</p>
     */
    @NameInMap("Marker")
    public String marker;

    /**
     * <strong>example:</strong>
     * <p>2</p>
     */
    @NameInMap("MaxKeys")
    public String maxKeys;

    /**
     * <strong>example:</strong>
     * <p>xyz</p>
     */
    @NameInMap("NextMarker")
    public String nextMarker;

    @NameInMap("Owner")
    public Owner owner;

    /**
     * <strong>example:</strong>
     * <p>abc</p>
     */
    @NameInMap("Prefix")
    public String prefix;

    public static ListAllMyCacheResult build(java.util.Map<String, ?> map) throws Exception {
        ListAllMyCacheResult self = new ListAllMyCacheResult();
        return TeaModel.build(map, self);
    }

    public ListAllMyCacheResult setCaches(Caches caches) {
        this.caches = caches;
        return this;
    }
    public Caches getCaches() {
        return this.caches;
    }

    public ListAllMyCacheResult setIsTruncated(Boolean isTruncated) {
        this.isTruncated = isTruncated;
        return this;
    }
    public Boolean getIsTruncated() {
        return this.isTruncated;
    }

    public ListAllMyCacheResult setMarker(String marker) {
        this.marker = marker;
        return this;
    }
    public String getMarker() {
        return this.marker;
    }

    public ListAllMyCacheResult setMaxKeys(String maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public String getMaxKeys() {
        return this.maxKeys;
    }

    public ListAllMyCacheResult setNextMarker(String nextMarker) {
        this.nextMarker = nextMarker;
        return this;
    }
    public String getNextMarker() {
        return this.nextMarker;
    }

    public ListAllMyCacheResult setOwner(Owner owner) {
        this.owner = owner;
        return this;
    }
    public Owner getOwner() {
        return this.owner;
    }

    public ListAllMyCacheResult setPrefix(String prefix) {
        this.prefix = prefix;
        return this;
    }
    public String getPrefix() {
        return this.prefix;
    }

    public static class Caches extends TeaModel {
        @NameInMap("Cache")
        public java.util.List<CacheBaseInfo> cache;

        public static Caches build(java.util.Map<String, ?> map) throws Exception {
            Caches self = new Caches();
            return TeaModel.build(map, self);
        }

        public Caches setCache(java.util.List<CacheBaseInfo> cache) {
            this.cache = cache;
            return this;
        }
        public java.util.List<CacheBaseInfo> getCache() {
            return this.cache;
        }

    }

}
