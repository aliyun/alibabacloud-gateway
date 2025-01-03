// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketDataAcceleratorRequest extends TeaModel {
    @NameInMap("DataAcceleratorConfiguration")
    public DataAcceleratorConfiguration dataAcceleratorConfiguration;

    public static PutBucketDataAcceleratorRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketDataAcceleratorRequest self = new PutBucketDataAcceleratorRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketDataAcceleratorRequest setDataAcceleratorConfiguration(DataAcceleratorConfiguration dataAcceleratorConfiguration) {
        this.dataAcceleratorConfiguration = dataAcceleratorConfiguration;
        return this;
    }
    public DataAcceleratorConfiguration getDataAcceleratorConfiguration() {
        return this.dataAcceleratorConfiguration;
    }

}
