// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class ListAccessPointsForObjectProcessRequest extends TeaModel {
    @NameInMap("continuation-token")
    public String continuationToken;

    @NameInMap("max-keys")
    public Long maxKeys;

    public static ListAccessPointsForObjectProcessRequest build(java.util.Map<String, ?> map) throws Exception {
        ListAccessPointsForObjectProcessRequest self = new ListAccessPointsForObjectProcessRequest();
        return TeaModel.build(map, self);
    }

    public ListAccessPointsForObjectProcessRequest setContinuationToken(String continuationToken) {
        this.continuationToken = continuationToken;
        return this;
    }
    public String getContinuationToken() {
        return this.continuationToken;
    }

    public ListAccessPointsForObjectProcessRequest setMaxKeys(Long maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Long getMaxKeys() {
        return this.maxKeys;
    }

}
