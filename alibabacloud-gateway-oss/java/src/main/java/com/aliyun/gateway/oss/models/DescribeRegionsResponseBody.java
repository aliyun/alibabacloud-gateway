// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class DescribeRegionsResponseBody extends TeaModel {
    /**
     * <p>The information about the regions.</p>
     */
    @NameInMap("RegionInfoList")
    public RegionInfoList regionInfoList;

    public static DescribeRegionsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        DescribeRegionsResponseBody self = new DescribeRegionsResponseBody();
        return TeaModel.build(map, self);
    }

    public DescribeRegionsResponseBody setRegionInfoList(RegionInfoList regionInfoList) {
        this.regionInfoList = regionInfoList;
        return this;
    }
    public RegionInfoList getRegionInfoList() {
        return this.regionInfoList;
    }

    public static class RegionInfoList extends TeaModel {
        /**
         * <p>The information about the regions.</p>
         */
        @NameInMap("RegionInfo")
        public java.util.List<RegionInfo> regionInfos;

        public static RegionInfoList build(java.util.Map<String, ?> map) throws Exception {
            RegionInfoList self = new RegionInfoList();
            return TeaModel.build(map, self);
        }

        public RegionInfoList setRegionInfos(java.util.List<RegionInfo> regionInfos) {
            this.regionInfos = regionInfos;
            return this;
        }
        public java.util.List<RegionInfo> getRegionInfos() {
            return this.regionInfos;
        }

    }

}
