// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DeleteCacheRequest extends TeaModel {
    @NameInMap("x-oss-datalake-cache-available-zone")
    public String xOssDatalakeCacheAvailableZone;

    @NameInMap("x-oss-datalake-cache-name")
    public String xOssDatalakeCacheName;

    public static DeleteCacheRequest build(java.util.Map<String, ?> map) throws Exception {
        DeleteCacheRequest self = new DeleteCacheRequest();
        return TeaModel.build(map, self);
    }

    public DeleteCacheRequest setXOssDatalakeCacheAvailableZone(String xOssDatalakeCacheAvailableZone) {
        this.xOssDatalakeCacheAvailableZone = xOssDatalakeCacheAvailableZone;
        return this;
    }
    public String getXOssDatalakeCacheAvailableZone() {
        return this.xOssDatalakeCacheAvailableZone;
    }

    public DeleteCacheRequest setXOssDatalakeCacheName(String xOssDatalakeCacheName) {
        this.xOssDatalakeCacheName = xOssDatalakeCacheName;
        return this;
    }
    public String getXOssDatalakeCacheName() {
        return this.xOssDatalakeCacheName;
    }

}
