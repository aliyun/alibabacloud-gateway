// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketRefererRequest extends TeaModel {
    /**
     * <p>The container that stores the hotlink protection configurations.</p>
     */
    @NameInMap("RefererConfiguration")
    public RefererConfiguration refererConfiguration;

    public static PutBucketRefererRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketRefererRequest self = new PutBucketRefererRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketRefererRequest setRefererConfiguration(RefererConfiguration refererConfiguration) {
        this.refererConfiguration = refererConfiguration;
        return this;
    }
    public RefererConfiguration getRefererConfiguration() {
        return this.refererConfiguration;
    }

}
