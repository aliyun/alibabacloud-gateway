// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListBucketDataRedundancyTransitionResponseBody extends TeaModel {
    @NameInMap("ListBucketDataRedundancyTransition")
    public ListBucketDataRedundancyTransition listBucketDataRedundancyTransition;

    public static ListBucketDataRedundancyTransitionResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListBucketDataRedundancyTransitionResponseBody self = new ListBucketDataRedundancyTransitionResponseBody();
        return TeaModel.build(map, self);
    }

    public ListBucketDataRedundancyTransitionResponseBody setListBucketDataRedundancyTransition(ListBucketDataRedundancyTransition listBucketDataRedundancyTransition) {
        this.listBucketDataRedundancyTransition = listBucketDataRedundancyTransition;
        return this;
    }
    public ListBucketDataRedundancyTransition getListBucketDataRedundancyTransition() {
        return this.listBucketDataRedundancyTransition;
    }

    public static class ListBucketDataRedundancyTransition extends TeaModel {
        @NameInMap("BucketDataRedundancyTransition")
        public BucketDataRedundancyTransition bucketDataRedundancyTransition;

        public static ListBucketDataRedundancyTransition build(java.util.Map<String, ?> map) throws Exception {
            ListBucketDataRedundancyTransition self = new ListBucketDataRedundancyTransition();
            return TeaModel.build(map, self);
        }

        public ListBucketDataRedundancyTransition setBucketDataRedundancyTransition(BucketDataRedundancyTransition bucketDataRedundancyTransition) {
            this.bucketDataRedundancyTransition = bucketDataRedundancyTransition;
            return this;
        }
        public BucketDataRedundancyTransition getBucketDataRedundancyTransition() {
            return this.bucketDataRedundancyTransition;
        }

    }

}

