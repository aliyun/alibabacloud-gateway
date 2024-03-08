// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class DescribeRegionsResponseBody extends TeaModel {
    /**
     * <p>The information about the regions.</p>
     */
    @NameInMap("RegionInfo")
    public java.util.List<RegionInfo> regionInfo;

    public static DescribeRegionsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        DescribeRegionsResponseBody self = new DescribeRegionsResponseBody();
        return TeaModel.build(map, self);
    }

    public DescribeRegionsResponseBody setRegionInfo(java.util.List<RegionInfo> regionInfo) {
        this.regionInfo = regionInfo;
        return this;
    }
    public java.util.List<RegionInfo> getRegionInfo() {
        return this.regionInfo;
    }

}
