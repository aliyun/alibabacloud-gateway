// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketCallbackPolicyRequest extends TeaModel {
    @NameInMap("BucketCallbackPolicy")
    public CallbackPolicy bucketCallbackPolicy;

    public static PutBucketCallbackPolicyRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketCallbackPolicyRequest self = new PutBucketCallbackPolicyRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketCallbackPolicyRequest setBucketCallbackPolicy(CallbackPolicy bucketCallbackPolicy) {
        this.bucketCallbackPolicy = bucketCallbackPolicy;
        return this;
    }
    public CallbackPolicy getBucketCallbackPolicy() {
        return this.bucketCallbackPolicy;
    }

}
