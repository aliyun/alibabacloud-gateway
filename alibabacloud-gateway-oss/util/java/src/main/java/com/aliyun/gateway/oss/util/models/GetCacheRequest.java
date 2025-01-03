// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetCacheRequest extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("x-oss-datalake-cache-available-zone")
    public String xOssDatalakeCacheAvailableZone;

    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("x-oss-datalake-cache-name")
    public String xOssDatalakeCacheName;

    @NameInMap("x-oss-datalake-cache-verbose")
    public Boolean xOssDatalakeCacheVerbose;

    public static GetCacheRequest build(java.util.Map<String, ?> map) throws Exception {
        GetCacheRequest self = new GetCacheRequest();
        return TeaModel.build(map, self);
    }

    public GetCacheRequest setXOssDatalakeCacheAvailableZone(String xOssDatalakeCacheAvailableZone) {
        this.xOssDatalakeCacheAvailableZone = xOssDatalakeCacheAvailableZone;
        return this;
    }
    public String getXOssDatalakeCacheAvailableZone() {
        return this.xOssDatalakeCacheAvailableZone;
    }

    public GetCacheRequest setXOssDatalakeCacheName(String xOssDatalakeCacheName) {
        this.xOssDatalakeCacheName = xOssDatalakeCacheName;
        return this;
    }
    public String getXOssDatalakeCacheName() {
        return this.xOssDatalakeCacheName;
    }

    public GetCacheRequest setXOssDatalakeCacheVerbose(Boolean xOssDatalakeCacheVerbose) {
        this.xOssDatalakeCacheVerbose = xOssDatalakeCacheVerbose;
        return this;
    }
    public Boolean getXOssDatalakeCacheVerbose() {
        return this.xOssDatalakeCacheVerbose;
    }

}
