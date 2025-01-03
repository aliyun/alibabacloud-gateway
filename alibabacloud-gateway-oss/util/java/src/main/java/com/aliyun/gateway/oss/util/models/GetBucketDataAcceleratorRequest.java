// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketDataAcceleratorRequest extends TeaModel {
    @NameInMap("verbose")
    public String verbose;

    public static GetBucketDataAcceleratorRequest build(java.util.Map<String, ?> map) throws Exception {
        GetBucketDataAcceleratorRequest self = new GetBucketDataAcceleratorRequest();
        return TeaModel.build(map, self);
    }

    public GetBucketDataAcceleratorRequest setVerbose(String verbose) {
        this.verbose = verbose;
        return this;
    }
    public String getVerbose() {
        return this.verbose;
    }

}
