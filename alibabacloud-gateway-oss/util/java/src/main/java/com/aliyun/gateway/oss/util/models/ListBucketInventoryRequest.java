// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class ListBucketInventoryRequest extends TeaModel {
    /**
     * <p>Specify the start position of the list operation. You can obtain this token from the NextContinuationToken field of last ListBucketInventory\"s result.</p>
     */
    @NameInMap("continuation-token")
    public String continuationToken;

    public static ListBucketInventoryRequest build(java.util.Map<String, ?> map) throws Exception {
        ListBucketInventoryRequest self = new ListBucketInventoryRequest();
        return TeaModel.build(map, self);
    }

    public ListBucketInventoryRequest setContinuationToken(String continuationToken) {
        this.continuationToken = continuationToken;
        return this;
    }
    public String getContinuationToken() {
        return this.continuationToken;
    }

}
