// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class MetaQueryOpenRequest extends TeaModel {
    @NameInMap("Filters")
    public Filters filters;

    public static MetaQueryOpenRequest build(java.util.Map<String, ?> map) throws Exception {
        MetaQueryOpenRequest self = new MetaQueryOpenRequest();
        return TeaModel.build(map, self);
    }

    public MetaQueryOpenRequest setFilters(Filters filters) {
        this.filters = filters;
        return this;
    }
    public Filters getFilters() {
        return this.filters;
    }

    public static class Filters extends TeaModel {
        @NameInMap("Filter")
        public java.util.List<String> filter;

        public static Filters build(java.util.Map<String, ?> map) throws Exception {
            Filters self = new Filters();
            return TeaModel.build(map, self);
        }

        public Filters setFilter(java.util.List<String> filter) {
            this.filter = filter;
            return this;
        }
        public java.util.List<String> getFilter() {
            return this.filter;
        }

    }

}
