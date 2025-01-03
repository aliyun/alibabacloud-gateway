// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListAccessPointsRequest extends TeaModel {
    /**
     * <p>The token from which the listing operation starts. You must specify the value of NextContinuationToken that is obtained from the previous query as the value of continuation-token.</p>
     * 
     * <strong>example:</strong>
     * <p>abc</p>
     */
    @NameInMap("continuation-token")
    public String continuationToken;

    /**
     * <p>The maximum number of access points that can be returned. Valid values:</p>
     * <ul>
     * <li>For user-level access points: (0,1000].</li>
     * <li>For bucket-level access points: (0,100].</li>
     * </ul>
     * 
     * <strong>example:</strong>
     * <p>10</p>
     */
    @NameInMap("max-keys")
    public Long maxKeys;

    public static ListAccessPointsRequest build(java.util.Map<String, ?> map) throws Exception {
        ListAccessPointsRequest self = new ListAccessPointsRequest();
        return TeaModel.build(map, self);
    }

    public ListAccessPointsRequest setContinuationToken(String continuationToken) {
        this.continuationToken = continuationToken;
        return this;
    }
    public String getContinuationToken() {
        return this.continuationToken;
    }

    public ListAccessPointsRequest setMaxKeys(Long maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Long getMaxKeys() {
        return this.maxKeys;
    }

}
