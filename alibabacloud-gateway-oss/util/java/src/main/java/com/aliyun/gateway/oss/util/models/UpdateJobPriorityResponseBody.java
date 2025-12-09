// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class UpdateJobPriorityResponseBody extends TeaModel {
    @NameInMap("UpdateJobPriorityResult")
    public BatchOperationUpdateJobPriorityResult updateJobPriorityResult;

    public static UpdateJobPriorityResponseBody build(java.util.Map<String, ?> map) throws Exception {
        UpdateJobPriorityResponseBody self = new UpdateJobPriorityResponseBody();
        return TeaModel.build(map, self);
    }

    public UpdateJobPriorityResponseBody setUpdateJobPriorityResult(BatchOperationUpdateJobPriorityResult updateJobPriorityResult) {
        this.updateJobPriorityResult = updateJobPriorityResult;
        return this;
    }
    public BatchOperationUpdateJobPriorityResult getUpdateJobPriorityResult() {
        return this.updateJobPriorityResult;
    }

}
