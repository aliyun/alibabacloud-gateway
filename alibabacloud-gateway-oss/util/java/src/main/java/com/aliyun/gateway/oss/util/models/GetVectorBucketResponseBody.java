// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetVectorBucketResponseBody extends TeaModel {
    @NameInMap("BucketInfo")
    public BucketInfo bucketInfo;

    public static GetVectorBucketResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetVectorBucketResponseBody self = new GetVectorBucketResponseBody();
        return TeaModel.build(map, self);
    }

    public GetVectorBucketResponseBody setBucketInfo(BucketInfo bucketInfo) {
        this.bucketInfo = bucketInfo;
        return this;
    }
    public BucketInfo getBucketInfo() {
        return this.bucketInfo;
    }

}
