// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListResourcePoolBucketGroupQoSInfosRequest extends TeaModel {
    @NameInMap("continuation-token")
    public String continuationToken;

    @NameInMap("max-keys")
    public String maxKeys;

    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("resourcePool")
    public String resourcePool;

    public static ListResourcePoolBucketGroupQoSInfosRequest build(java.util.Map<String, ?> map) throws Exception {
        ListResourcePoolBucketGroupQoSInfosRequest self = new ListResourcePoolBucketGroupQoSInfosRequest();
        return TeaModel.build(map, self);
    }

    public ListResourcePoolBucketGroupQoSInfosRequest setContinuationToken(String continuationToken) {
        this.continuationToken = continuationToken;
        return this;
    }
    public String getContinuationToken() {
        return this.continuationToken;
    }

    public ListResourcePoolBucketGroupQoSInfosRequest setMaxKeys(String maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public String getMaxKeys() {
        return this.maxKeys;
    }

    public ListResourcePoolBucketGroupQoSInfosRequest setResourcePool(String resourcePool) {
        this.resourcePool = resourcePool;
        return this;
    }
    public String getResourcePool() {
        return this.resourcePool;
    }

}
