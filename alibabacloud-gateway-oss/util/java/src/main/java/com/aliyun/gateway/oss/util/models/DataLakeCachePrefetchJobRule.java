// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DataLakeCachePrefetchJobRule extends TeaModel {
    @NameInMap("PrefixFilter")
    public PrefixFilter prefixFilter;

    @NameInMap("Tag")
    public String tag;

    public static DataLakeCachePrefetchJobRule build(java.util.Map<String, ?> map) throws Exception {
        DataLakeCachePrefetchJobRule self = new DataLakeCachePrefetchJobRule();
        return TeaModel.build(map, self);
    }

    public DataLakeCachePrefetchJobRule setPrefixFilter(PrefixFilter prefixFilter) {
        this.prefixFilter = prefixFilter;
        return this;
    }
    public PrefixFilter getPrefixFilter() {
        return this.prefixFilter;
    }

    public DataLakeCachePrefetchJobRule setTag(String tag) {
        this.tag = tag;
        return this;
    }
    public String getTag() {
        return this.tag;
    }

    public static class Excludes extends TeaModel {
        @NameInMap("Exclude")
        public java.util.List<String> exclude;

        public static Excludes build(java.util.Map<String, ?> map) throws Exception {
            Excludes self = new Excludes();
            return TeaModel.build(map, self);
        }

        public Excludes setExclude(java.util.List<String> exclude) {
            this.exclude = exclude;
            return this;
        }
        public java.util.List<String> getExclude() {
            return this.exclude;
        }

    }

    public static class Includes extends TeaModel {
        @NameInMap("Include")
        public java.util.List<String> include;

        public static Includes build(java.util.Map<String, ?> map) throws Exception {
            Includes self = new Includes();
            return TeaModel.build(map, self);
        }

        public Includes setInclude(java.util.List<String> include) {
            this.include = include;
            return this;
        }
        public java.util.List<String> getInclude() {
            return this.include;
        }

    }

    public static class PrefixFilter extends TeaModel {
        @NameInMap("Excludes")
        public Excludes excludes;

        @NameInMap("Includes")
        public Includes includes;

        public static PrefixFilter build(java.util.Map<String, ?> map) throws Exception {
            PrefixFilter self = new PrefixFilter();
            return TeaModel.build(map, self);
        }

        public PrefixFilter setExcludes(Excludes excludes) {
            this.excludes = excludes;
            return this;
        }
        public Excludes getExcludes() {
            return this.excludes;
        }

        public PrefixFilter setIncludes(Includes includes) {
            this.includes = includes;
            return this;
        }
        public Includes getIncludes() {
            return this.includes;
        }

    }

}
