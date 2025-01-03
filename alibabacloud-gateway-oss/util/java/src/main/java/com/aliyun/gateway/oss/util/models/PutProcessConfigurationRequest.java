// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutProcessConfigurationRequest extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("BucketProcessConfiguration")
    public BucketProcessConfiguration bucketProcessConfiguration;

    public static PutProcessConfigurationRequest build(java.util.Map<String, ?> map) throws Exception {
        PutProcessConfigurationRequest self = new PutProcessConfigurationRequest();
        return TeaModel.build(map, self);
    }

    public PutProcessConfigurationRequest setBucketProcessConfiguration(BucketProcessConfiguration bucketProcessConfiguration) {
        this.bucketProcessConfiguration = bucketProcessConfiguration;
        return this;
    }
    public BucketProcessConfiguration getBucketProcessConfiguration() {
        return this.bucketProcessConfiguration;
    }

}
