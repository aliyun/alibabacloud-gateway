// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutCnameRequest extends TeaModel {
    /**
     * <p>The container that stores the CNAME record.</p>
     */
    @NameInMap("BucketCnameConfiguration")
    public BucketCnameConfiguration bucketCnameConfiguration;

    public static PutCnameRequest build(java.util.Map<String, ?> map) throws Exception {
        PutCnameRequest self = new PutCnameRequest();
        return TeaModel.build(map, self);
    }

    public PutCnameRequest setBucketCnameConfiguration(BucketCnameConfiguration bucketCnameConfiguration) {
        this.bucketCnameConfiguration = bucketCnameConfiguration;
        return this;
    }
    public BucketCnameConfiguration getBucketCnameConfiguration() {
        return this.bucketCnameConfiguration;
    }

}
