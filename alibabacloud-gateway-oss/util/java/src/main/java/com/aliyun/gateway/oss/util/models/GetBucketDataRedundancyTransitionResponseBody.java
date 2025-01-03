// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketDataRedundancyTransitionResponseBody extends TeaModel {
    /**
     * <p>The container for a specific redundancy type change task.</p>
     */
    @NameInMap("BucketDataRedundancyTransition")
    public BucketDataRedundancyTransition bucketDataRedundancyTransition;

    public static GetBucketDataRedundancyTransitionResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketDataRedundancyTransitionResponseBody self = new GetBucketDataRedundancyTransitionResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketDataRedundancyTransitionResponseBody setBucketDataRedundancyTransition(BucketDataRedundancyTransition bucketDataRedundancyTransition) {
        this.bucketDataRedundancyTransition = bucketDataRedundancyTransition;
        return this;
    }
    public BucketDataRedundancyTransition getBucketDataRedundancyTransition() {
        return this.bucketDataRedundancyTransition;
    }

}
