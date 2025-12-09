// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListUserDataRedundancyTransitionRequest extends TeaModel {
    /**
     * <p>Specifies that the List operation should start from this token.</p>
     * 
     * <strong>example:</strong>
     * <p>abc</p>
     */
    @NameInMap("continuation-token")
    public String continuationToken;

    /**
     * <p>Limits the maximum number of tasks returned in this request. Value range: 1-100.</p>
     * 
     * <strong>example:</strong>
     * <p>10</p>
     */
    @NameInMap("max-keys")
    public Integer maxKeys;

    public static ListUserDataRedundancyTransitionRequest build(java.util.Map<String, ?> map) throws Exception {
        ListUserDataRedundancyTransitionRequest self = new ListUserDataRedundancyTransitionRequest();
        return TeaModel.build(map, self);
    }

    public ListUserDataRedundancyTransitionRequest setContinuationToken(String continuationToken) {
        this.continuationToken = continuationToken;
        return this;
    }
    public String getContinuationToken() {
        return this.continuationToken;
    }

    public ListUserDataRedundancyTransitionRequest setMaxKeys(Integer maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Integer getMaxKeys() {
        return this.maxKeys;
    }

}
