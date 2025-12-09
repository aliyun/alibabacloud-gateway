// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DeleteBucketDataAcceleratorRequest extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>cn-hangzhou-a</p>
     */
    @NameInMap("x-oss-datalake-cache-available-zone")
    public String xOssDatalakeCacheAvailableZone;

    public static DeleteBucketDataAcceleratorRequest build(java.util.Map<String, ?> map) throws Exception {
        DeleteBucketDataAcceleratorRequest self = new DeleteBucketDataAcceleratorRequest();
        return TeaModel.build(map, self);
    }

    public DeleteBucketDataAcceleratorRequest setXOssDatalakeCacheAvailableZone(String xOssDatalakeCacheAvailableZone) {
        this.xOssDatalakeCacheAvailableZone = xOssDatalakeCacheAvailableZone;
        return this;
    }
    public String getXOssDatalakeCacheAvailableZone() {
        return this.xOssDatalakeCacheAvailableZone;
    }

}
