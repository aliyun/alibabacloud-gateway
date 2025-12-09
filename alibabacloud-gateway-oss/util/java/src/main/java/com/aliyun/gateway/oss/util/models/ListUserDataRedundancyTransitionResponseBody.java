// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListUserDataRedundancyTransitionResponseBody extends TeaModel {
    /**
     * <p>Container for listing storage redundancy transition tasks.</p>
     */
    @NameInMap("ListBucketDataRedundancyTransition")
    public ListBucketDataRedundancyTransition listBucketDataRedundancyTransition;

    public static ListUserDataRedundancyTransitionResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListUserDataRedundancyTransitionResponseBody self = new ListUserDataRedundancyTransitionResponseBody();
        return TeaModel.build(map, self);
    }

    public ListUserDataRedundancyTransitionResponseBody setListBucketDataRedundancyTransition(ListBucketDataRedundancyTransition listBucketDataRedundancyTransition) {
        this.listBucketDataRedundancyTransition = listBucketDataRedundancyTransition;
        return this;
    }
    public ListBucketDataRedundancyTransition getListBucketDataRedundancyTransition() {
        return this.listBucketDataRedundancyTransition;
    }

    public static class ListBucketDataRedundancyTransition extends TeaModel {
        /**
         * <p>Container for storage redundancy transition tasks.</p>
         */
        @NameInMap("BucketDataRedundancyTransition")
        public java.util.List<BucketDataRedundancyTransition> bucketDataRedundancyTransition;

        /**
         * <p>Indicates whether the results in the request are truncated. The values are as follows:</p>
         * <p>true: Indicates that not all results are returned in this request.</p>
         * <p>false: Indicates that all results are returned in this request.</p>
         * 
         * <strong>example:</strong>
         * <p>true</p>
         */
        @NameInMap("IsTruncated")
        public Boolean isTruncated;

        /**
         * <p>Indicates that the current ListUserDataRedundancyTransition request contains subsequent results, and you need to specify NextContinuationToken as continuation-token to continue retrieving the results.</p>
         * 
         * <strong>example:</strong>
         * <p>def</p>
         */
        @NameInMap("NextContinuationToken")
        public String nextContinuationToken;

        public static ListBucketDataRedundancyTransition build(java.util.Map<String, ?> map) throws Exception {
            ListBucketDataRedundancyTransition self = new ListBucketDataRedundancyTransition();
            return TeaModel.build(map, self);
        }

        public ListBucketDataRedundancyTransition setBucketDataRedundancyTransition(java.util.List<BucketDataRedundancyTransition> bucketDataRedundancyTransition) {
            this.bucketDataRedundancyTransition = bucketDataRedundancyTransition;
            return this;
        }
        public java.util.List<BucketDataRedundancyTransition> getBucketDataRedundancyTransition() {
            return this.bucketDataRedundancyTransition;
        }

        public ListBucketDataRedundancyTransition setIsTruncated(Boolean isTruncated) {
            this.isTruncated = isTruncated;
            return this;
        }
        public Boolean getIsTruncated() {
            return this.isTruncated;
        }

        public ListBucketDataRedundancyTransition setNextContinuationToken(String nextContinuationToken) {
            this.nextContinuationToken = nextContinuationToken;
            return this;
        }
        public String getNextContinuationToken() {
            return this.nextContinuationToken;
        }

    }

}
