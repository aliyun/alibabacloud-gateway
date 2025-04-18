// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListResourcePoolRequesterQoSInfosRequest extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>26753xxxxxxxx14340</p>
     */
    @NameInMap("continuation-token")
    public String continuationToken;

    /**
     * <strong>example:</strong>
     * <p>100</p>
     */
    @NameInMap("max-keys")
    public Long maxKeys;

    /**
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>rp-01</p>
     */
    @NameInMap("resourcePool")
    public String resourcePool;

    public static ListResourcePoolRequesterQoSInfosRequest build(java.util.Map<String, ?> map) throws Exception {
        ListResourcePoolRequesterQoSInfosRequest self = new ListResourcePoolRequesterQoSInfosRequest();
        return TeaModel.build(map, self);
    }

    public ListResourcePoolRequesterQoSInfosRequest setContinuationToken(String continuationToken) {
        this.continuationToken = continuationToken;
        return this;
    }
    public String getContinuationToken() {
        return this.continuationToken;
    }

    public ListResourcePoolRequesterQoSInfosRequest setMaxKeys(Long maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Long getMaxKeys() {
        return this.maxKeys;
    }

    public ListResourcePoolRequesterQoSInfosRequest setResourcePool(String resourcePool) {
        this.resourcePool = resourcePool;
        return this;
    }
    public String getResourcePool() {
        return this.resourcePool;
    }

}
