// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketStatResponseBody extends TeaModel {
    /**
     * <p>The container that stores all information returned for the GetBucketStat request.</p>
     */
    @NameInMap("BucketStat")
    public BucketStat bucketStat;

    public static GetBucketStatResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketStatResponseBody self = new GetBucketStatResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketStatResponseBody setBucketStat(BucketStat bucketStat) {
        this.bucketStat = bucketStat;
        return this;
    }
    public BucketStat getBucketStat() {
        return this.bucketStat;
    }

}
