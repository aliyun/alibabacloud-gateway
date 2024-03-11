// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class PutBucketRequest extends TeaModel {
    /**
     * <p>The container that stores the information about the bucket to be created.</p>
     */
    @NameInMap("CreateBucketConfiguration")
    public CreateBucketConfiguration createBucketConfiguration;

    public static PutBucketRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketRequest self = new PutBucketRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketRequest setCreateBucketConfiguration(CreateBucketConfiguration createBucketConfiguration) {
        this.createBucketConfiguration = createBucketConfiguration;
        return this;
    }
    public CreateBucketConfiguration getCreateBucketConfiguration() {
        return this.createBucketConfiguration;
    }

}
