// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class UpdateJobPriorityRequest extends TeaModel {
    @NameInMap("batchJobId")
    public String batchJobId;

    @NameInMap("targetPriority")
    public Integer targetPriority;

    public static UpdateJobPriorityRequest build(java.util.Map<String, ?> map) throws Exception {
        UpdateJobPriorityRequest self = new UpdateJobPriorityRequest();
        return TeaModel.build(map, self);
    }

    public UpdateJobPriorityRequest setBatchJobId(String batchJobId) {
        this.batchJobId = batchJobId;
        return this;
    }
    public String getBatchJobId() {
        return this.batchJobId;
    }

    public UpdateJobPriorityRequest setTargetPriority(Integer targetPriority) {
        this.targetPriority = targetPriority;
        return this;
    }
    public Integer getTargetPriority() {
        return this.targetPriority;
    }

}
