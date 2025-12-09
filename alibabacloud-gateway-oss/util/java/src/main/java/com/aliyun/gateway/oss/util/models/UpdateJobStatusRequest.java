// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class UpdateJobStatusRequest extends TeaModel {
    @NameInMap("batchJobId")
    public String batchJobId;

    @NameInMap("requestedJobStatus")
    public String requestedJobStatus;

    @NameInMap("statusUpdateReason")
    public String statusUpdateReason;

    public static UpdateJobStatusRequest build(java.util.Map<String, ?> map) throws Exception {
        UpdateJobStatusRequest self = new UpdateJobStatusRequest();
        return TeaModel.build(map, self);
    }

    public UpdateJobStatusRequest setBatchJobId(String batchJobId) {
        this.batchJobId = batchJobId;
        return this;
    }
    public String getBatchJobId() {
        return this.batchJobId;
    }

    public UpdateJobStatusRequest setRequestedJobStatus(String requestedJobStatus) {
        this.requestedJobStatus = requestedJobStatus;
        return this;
    }
    public String getRequestedJobStatus() {
        return this.requestedJobStatus;
    }

    public UpdateJobStatusRequest setStatusUpdateReason(String statusUpdateReason) {
        this.statusUpdateReason = statusUpdateReason;
        return this;
    }
    public String getStatusUpdateReason() {
        return this.statusUpdateReason;
    }

}
