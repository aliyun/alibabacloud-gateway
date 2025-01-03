// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketLocationResponseBody extends TeaModel {
    /**
     * <p>The region in which the bucket resides.\
     * Examples: oss-cn-hangzhou, oss-cn-shanghai, oss-cn-qingdao, oss-cn-beijing, oss-cn-zhangjiakou, oss-cn-hongkong, oss-cn-shenzhen, oss-us-west-1, oss-us-east-1, and oss-ap-southeast-1.</p>
     * <p>\
     * For more information about the regions in which buckets reside, see <a href="https://help.aliyun.com/document_detail/31837.html">Regions and endpoints</a>.</p>
     * 
     * <strong>example:</strong>
     * <p>oss-cn-hangzhou</p>
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
