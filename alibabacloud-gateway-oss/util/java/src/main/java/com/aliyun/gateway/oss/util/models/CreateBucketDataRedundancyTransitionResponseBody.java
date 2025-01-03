// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateBucketDataRedundancyTransitionResponseBody extends TeaModel {
    /**
     * <p>The container in which the redundancy type conversion task is stored.</p>
     */
    @NameInMap("BucketDataRedundancyTransition")
    public BucketDataRedundancyTransition bucketDataRedundancyTransition;

    public static CreateBucketDataRedundancyTransitionResponseBody build(java.util.Map<String, ?> map) throws Exception {
        CreateBucketDataRedundancyTransitionResponseBody self = new CreateBucketDataRedundancyTransitionResponseBody();
        return TeaModel.build(map, self);
    }

    public CreateBucketDataRedundancyTransitionResponseBody setBucketDataRedundancyTransition(BucketDataRedundancyTransition bucketDataRedundancyTransition) {
        this.bucketDataRedundancyTransition = bucketDataRedundancyTransition;
        return this;
    }
    public BucketDataRedundancyTransition getBucketDataRedundancyTransition() {
        return this.bucketDataRedundancyTransition;
    }

    public static class BucketDataRedundancyTransition extends TeaModel {
        /**
         * <p>The ID of the redundancy type conversion task. The ID can be used to view and delete the redundancy type conversion task.</p>
         * 
         * <strong>example:</strong>
         * <p>4be5beb0f74f490186311b268bf6****</p>
         */
        @NameInMap("TaskId")
        public String taskId;

        public static BucketDataRedundancyTransition build(java.util.Map<String, ?> map) throws Exception {
            BucketDataRedundancyTransition self = new BucketDataRedundancyTransition();
            return TeaModel.build(map, self);
        }

        public BucketDataRedundancyTransition setTaskId(String taskId) {
            this.taskId = taskId;
            return this;
        }
        public String getTaskId() {
            return this.taskId;
        }

    }

}
