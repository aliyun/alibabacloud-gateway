// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListUserRegionsResult extends TeaModel {
    @NameInMap("Regions")
    public Regions regions;

    public static ListUserRegionsResult build(java.util.Map<String, ?> map) throws Exception {
        ListUserRegionsResult self = new ListUserRegionsResult();
        return TeaModel.build(map, self);
    }

    public ListUserRegionsResult setRegions(Regions regions) {
        this.regions = regions;
        return this;
    }
    public Regions getRegions() {
        return this.regions;
    }

    public static class Regions extends TeaModel {
        @NameInMap("Region")
        public java.util.List<String> region;

        public static Regions build(java.util.Map<String, ?> map) throws Exception {
            Regions self = new Regions();
            return TeaModel.build(map, self);
        }

        public Regions setRegion(java.util.List<String> region) {
            this.region = region;
            return this;
        }
        public java.util.List<String> getRegion() {
            return this.region;
        }

    }

}
