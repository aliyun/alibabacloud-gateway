// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GroupBucketInfo extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>test-bucket</p>
     */
    @NameInMap("BucketName")
    public String bucketName;

    public static GroupBucketInfo build(java.util.Map<String, ?> map) throws Exception {
        GroupBucketInfo self = new GroupBucketInfo();
        return TeaModel.build(map, self);
    }

    public GroupBucketInfo setBucketName(String bucketName) {
        this.bucketName = bucketName;
        return this;
    }
    public String getBucketName() {
        return this.bucketName;
    }

}
