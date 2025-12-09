// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketDataAcceleratorRequest extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("verbose")
    public String verbose;

    /**
     * <strong>example:</strong>
     * <p>cn-hangzhou-a</p>
     */
    @NameInMap("x-oss-datalake-cache-available-zone")
    public String xOssDatalakeCacheAvailableZone;

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

    public GetBucketDataAcceleratorRequest setXOssDatalakeCacheAvailableZone(String xOssDatalakeCacheAvailableZone) {
        this.xOssDatalakeCacheAvailableZone = xOssDatalakeCacheAvailableZone;
        return this;
    }
    public String getXOssDatalakeCacheAvailableZone() {
        return this.xOssDatalakeCacheAvailableZone;
    }

}
