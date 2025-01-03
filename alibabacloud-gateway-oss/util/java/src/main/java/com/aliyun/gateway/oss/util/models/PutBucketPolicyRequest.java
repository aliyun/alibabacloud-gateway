// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketPolicyRequest extends TeaModel {
    /**
     * <p>The request parameters.</p>
     * <p>This parameter is required.</p>
     */
    @NameInMap("body")
    public String policy;

    public static PutBucketPolicyRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketPolicyRequest self = new PutBucketPolicyRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketPolicyRequest setPolicy(String policy) {
        this.policy = policy;
        return this;
    }
    public String getPolicy() {
        return this.policy;
    }

}
