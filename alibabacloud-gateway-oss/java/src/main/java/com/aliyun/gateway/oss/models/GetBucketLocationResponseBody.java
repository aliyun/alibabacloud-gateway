// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class GetBucketLocationResponseBody extends TeaModel {
    /**
     * <p>The region in which the bucket resides.\</p>
     * <p>Examples: oss-cn-hangzhou, oss-cn-shanghai, oss-cn-qingdao, oss-cn-beijing, oss-cn-zhangjiakou, oss-cn-hongkong, oss-cn-shenzhen, oss-us-west-1, oss-us-east-1, and oss-ap-southeast-1.</p>
     * <br>
     * <p>\</p>
     * <p>For more information about the regions in which buckets reside, see [Regions and endpoints](~~31837~~).</p>
     */
    @NameInMap("LocationConstraint")
    public String locationConstraint;

    public static GetBucketLocationResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketLocationResponseBody self = new GetBucketLocationResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketLocationResponseBody setLocationConstraint(String locationConstraint) {
        this.locationConstraint = locationConstraint;
        return this;
    }
    public String getLocationConstraint() {
        return this.locationConstraint;
    }

}
