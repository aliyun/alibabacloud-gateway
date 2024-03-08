// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ListUserDataRedundancyTransitionResponseBody extends TeaModel {
    @NameInMap("BucketDataRedundancyTransition")
    public java.util.List<BucketDataRedundancyTransition> bucketDataRedundancyTransition;

    @NameInMap("IsTruncated")
    public Boolean isTruncated;

    @NameInMap("NextContinuationToken")
    public String nextContinuationToken;

    public static ListUserDataRedundancyTransitionResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListUserDataRedundancyTransitionResponseBody self = new ListUserDataRedundancyTransitionResponseBody();
        return TeaModel.build(map, self);
    }

    public ListUserDataRedundancyTransitionResponseBody setBucketDataRedundancyTransition(java.util.List<BucketDataRedundancyTransition> bucketDataRedundancyTransition) {
        this.bucketDataRedundancyTransition = bucketDataRedundancyTransition;
        return this;
    }
    public java.util.List<BucketDataRedundancyTransition> getBucketDataRedundancyTransition() {
        return this.bucketDataRedundancyTransition;
    }

    public ListUserDataRedundancyTransitionResponseBody setIsTruncated(Boolean isTruncated) {
        this.isTruncated = isTruncated;
        return this;
    }
    public Boolean getIsTruncated() {
        return this.isTruncated;
    }

    public ListUserDataRedundancyTransitionResponseBody setNextContinuationToken(String nextContinuationToken) {
        this.nextContinuationToken = nextContinuationToken;
        return this;
    }
    public String getNextContinuationToken() {
        return this.nextContinuationToken;
    }

}
