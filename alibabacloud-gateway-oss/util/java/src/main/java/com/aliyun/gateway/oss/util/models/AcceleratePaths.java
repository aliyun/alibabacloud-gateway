// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class AcceleratePaths extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>sync-warmup</p>
     */
    @NameInMap("DefaultCachePolicy")
    public String defaultCachePolicy;

    @NameInMap("Path")
    public java.util.List<Path> path;

    public static AcceleratePaths build(java.util.Map<String, ?> map) throws Exception {
        AcceleratePaths self = new AcceleratePaths();
        return TeaModel.build(map, self);
    }

    public AcceleratePaths setDefaultCachePolicy(String defaultCachePolicy) {
        this.defaultCachePolicy = defaultCachePolicy;
        return this;
    }
    public String getDefaultCachePolicy() {
        return this.defaultCachePolicy;
    }

    public AcceleratePaths setPath(java.util.List<Path> path) {
        this.path = path;
        return this;
    }
    public java.util.List<Path> getPath() {
        return this.path;
    }

    public static class Path extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>sync-warmup</p>
         */
        @NameInMap("CachePolicy")
        public String cachePolicy;

        /**
         * <strong>example:</strong>
         * <p>acc-path/</p>
         */
        @NameInMap("Name")
        public String name;

        public static Path build(java.util.Map<String, ?> map) throws Exception {
            Path self = new Path();
            return TeaModel.build(map, self);
        }

        public Path setCachePolicy(String cachePolicy) {
            this.cachePolicy = cachePolicy;
            return this;
        }
        public String getCachePolicy() {
            return this.cachePolicy;
        }

        public Path setName(String name) {
            this.name = name;
            return this;
        }
        public String getName() {
            return this.name;
        }

    }

}
