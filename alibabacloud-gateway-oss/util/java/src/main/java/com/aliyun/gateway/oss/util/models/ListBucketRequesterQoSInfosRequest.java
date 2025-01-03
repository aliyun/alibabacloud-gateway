// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListBucketRequesterQoSInfosRequest extends TeaModel {
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

    public static ListBucketRequesterQoSInfosRequest build(java.util.Map<String, ?> map) throws Exception {
        ListBucketRequesterQoSInfosRequest self = new ListBucketRequesterQoSInfosRequest();
        return TeaModel.build(map, self);
    }

    public ListBucketRequesterQoSInfosRequest setContinuationToken(String continuationToken) {
        this.continuationToken = continuationToken;
        return this;
    }
    public String getContinuationToken() {
        return this.continuationToken;
    }

    public ListBucketRequesterQoSInfosRequest setMaxKeys(Long maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Long getMaxKeys() {
        return this.maxKeys;
    }

}
