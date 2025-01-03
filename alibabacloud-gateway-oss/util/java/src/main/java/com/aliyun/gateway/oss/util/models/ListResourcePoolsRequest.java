// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListResourcePoolsRequest extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>rp-01</p>
     */
    @NameInMap("continuation-token")
    public String continuationToken;

    /**
     * <strong>example:</strong>
     * <p>100</p>
     */
    @NameInMap("max-keys")
    public Long maxKeys;

    public static ListResourcePoolsRequest build(java.util.Map<String, ?> map) throws Exception {
        ListResourcePoolsRequest self = new ListResourcePoolsRequest();
        return TeaModel.build(map, self);
    }

    public ListResourcePoolsRequest setContinuationToken(String continuationToken) {
        this.continuationToken = continuationToken;
        return this;
    }
    public String getContinuationToken() {
        return this.continuationToken;
    }

    public ListResourcePoolsRequest setMaxKeys(Long maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Long getMaxKeys() {
        return this.maxKeys;
    }

}
