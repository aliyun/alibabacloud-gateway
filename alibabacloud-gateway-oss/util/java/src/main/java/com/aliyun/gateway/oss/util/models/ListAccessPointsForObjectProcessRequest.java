// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListAccessPointsForObjectProcessRequest extends TeaModel {
    /**
     * <p>The token from which the list operation must start. You can obtain this token from the NextContinuationToken element in the returned result.</p>
     * 
     * <strong>example:</strong>
     * <p>abc</p>
     */
    @NameInMap("continuation-token")
    public String continuationToken;

    /**
     * <p>The maximum number of Object FC Access Points to return.</p>
     * <p>Valid values: 1 to 1000</p>
     * <blockquote>
     * <p>If the list cannot be complete at a time due to the configurations of the max-keys element, the NextContinuationToken element is included in the response as the token for the next list.</p>
     * </blockquote>
     * 
     * <strong>example:</strong>
     * <p>10</p>
     */
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
