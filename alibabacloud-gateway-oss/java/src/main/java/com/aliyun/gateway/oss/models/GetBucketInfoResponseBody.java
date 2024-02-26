// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class GetBucketInfoResponseBody extends TeaModel {
    /**
     * <p>The container that stores the information about the bucket.</p>
     */
    @NameInMap("BucketInfo")
    public BucketInfo bucketInfo;

    public static GetBucketInfoResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketInfoResponseBody self = new GetBucketInfoResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketInfoResponseBody setBucketInfo(BucketInfo bucketInfo) {
        this.bucketInfo = bucketInfo;
        return this;
    }
    public BucketInfo getBucketInfo() {
        return this.bucketInfo;
    }

}
