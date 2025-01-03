// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketDataAcceleratorResponseBody extends TeaModel {
    @NameInMap("DataAccelerator")
    public DataAccelerator dataAccelerator;

    public static GetBucketDataAcceleratorResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketDataAcceleratorResponseBody self = new GetBucketDataAcceleratorResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketDataAcceleratorResponseBody setDataAccelerator(DataAccelerator dataAccelerator) {
        this.dataAccelerator = dataAccelerator;
        return this;
    }
    public DataAccelerator getDataAccelerator() {
        return this.dataAccelerator;
    }

}
