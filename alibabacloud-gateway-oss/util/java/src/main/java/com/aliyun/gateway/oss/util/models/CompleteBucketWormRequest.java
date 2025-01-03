// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CompleteBucketWormRequest extends TeaModel {
    /**
     * <p>The ID of the retention policy.</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>1666E2CFB2B3418****</p>
     */
    @NameInMap("wormId")
    public String wormId;

    public static CompleteBucketWormRequest build(java.util.Map<String, ?> map) throws Exception {
        CompleteBucketWormRequest self = new CompleteBucketWormRequest();
        return TeaModel.build(map, self);
    }

    public CompleteBucketWormRequest setWormId(String wormId) {
        this.wormId = wormId;
        return this;
    }
    public String getWormId() {
        return this.wormId;
    }

}
