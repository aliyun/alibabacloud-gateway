// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class CreateCnameTokenRequest extends TeaModel {
    /**
     * <p>The container that stores the CNAME record.</p>
     */
    @NameInMap("BucketCnameConfiguration")
    public BucketCnameConfiguration bucketCnameConfiguration;

    public static CreateCnameTokenRequest build(java.util.Map<String, ?> map) throws Exception {
        CreateCnameTokenRequest self = new CreateCnameTokenRequest();
        return TeaModel.build(map, self);
    }

    public CreateCnameTokenRequest setBucketCnameConfiguration(BucketCnameConfiguration bucketCnameConfiguration) {
        this.bucketCnameConfiguration = bucketCnameConfiguration;
        return this;
    }
    public BucketCnameConfiguration getBucketCnameConfiguration() {
        return this.bucketCnameConfiguration;
    }

}
