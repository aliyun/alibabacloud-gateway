// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BatchOperationUpdateJobStatusResult extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>IDBkODc3M2RlZjgyNjQ0NDViZGV****</p>
     */
    @NameInMap("JobId")
    public String jobId;

    /**
     * <strong>example:</strong>
     * <p>Preparing</p>
     */
    @NameInMap("Status")
    public String status;

    @NameInMap("StatusUpdateReason")
    public String statusUpdateReason;

    public static BatchOperationUpdateJobStatusResult build(java.util.Map<String, ?> map) throws Exception {
        BatchOperationUpdateJobStatusResult self = new BatchOperationUpdateJobStatusResult();
        return TeaModel.build(map, self);
    }

    public BatchOperationUpdateJobStatusResult setJobId(String jobId) {
        this.jobId = jobId;
        return this;
    }
    public String getJobId() {
        return this.jobId;
    }

    public BatchOperationUpdateJobStatusResult setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

    public BatchOperationUpdateJobStatusResult setStatusUpdateReason(String statusUpdateReason) {
        this.statusUpdateReason = statusUpdateReason;
        return this;
    }
    public String getStatusUpdateReason() {
        return this.statusUpdateReason;
    }

}
