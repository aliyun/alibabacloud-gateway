// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BatchOperationUpdateJobPriorityResult extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>IDBkODc3M2RlZjgyNjQ0NDViZGV****</p>
     */
    @NameInMap("JobId")
    public String jobId;

    /**
     * <strong>example:</strong>
     * <p>5</p>
     */
    @NameInMap("Priority")
    public Integer priority;

    public static BatchOperationUpdateJobPriorityResult build(java.util.Map<String, ?> map) throws Exception {
        BatchOperationUpdateJobPriorityResult self = new BatchOperationUpdateJobPriorityResult();
        return TeaModel.build(map, self);
    }

    public BatchOperationUpdateJobPriorityResult setJobId(String jobId) {
        this.jobId = jobId;
        return this;
    }
    public String getJobId() {
        return this.jobId;
    }

    public BatchOperationUpdateJobPriorityResult setPriority(Integer priority) {
        this.priority = priority;
        return this;
    }
    public Integer getPriority() {
        return this.priority;
    }

}
