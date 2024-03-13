// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketCallbackPolicyResponseBody extends TeaModel {
    @NameInMap("BucketCallbackPolicy")
    public CallbackPolicy bucketCallbackPolicy;

    public static GetBucketCallbackPolicyResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketCallbackPolicyResponseBody self = new GetBucketCallbackPolicyResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketCallbackPolicyResponseBody setBucketCallbackPolicy(CallbackPolicy bucketCallbackPolicy) {
        this.bucketCallbackPolicy = bucketCallbackPolicy;
        return this;
    }
    public CallbackPolicy getBucketCallbackPolicy() {
        return this.bucketCallbackPolicy;
    }

}
