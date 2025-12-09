// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BatchOperationListJobs extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>1730250699</p>
     */
    @NameInMap("CreationTime")
    public Long creationTime;

    @NameInMap("Description")
    public String description;

    /**
     * <strong>example:</strong>
     * <p>IDBkODc3M2RlZjgyNjQ0NDViZGV****</p>
     */
    @NameInMap("JobId")
    public String jobId;

    /**
     * <strong>example:</strong>
     * <p>CopyObject</p>
     */
    @NameInMap("Operation")
    public String operation;

    /**
     * <strong>example:</strong>
     * <p>5</p>
     */
    @NameInMap("Priority")
    public Integer priority;

    @NameInMap("ProgressSummary")
    public BatchOperationJobProgressSummary progressSummary;

    /**
     * <strong>example:</strong>
     * <p>Preparing</p>
     */
    @NameInMap("Status")
    public String status;

    /**
     * <strong>example:</strong>
     * <p>1730250699</p>
     */
    @NameInMap("TerminationDate")
    public Long terminationDate;

    public static BatchOperationListJobs build(java.util.Map<String, ?> map) throws Exception {
        BatchOperationListJobs self = new BatchOperationListJobs();
        return TeaModel.build(map, self);
    }

    public BatchOperationListJobs setCreationTime(Long creationTime) {
        this.creationTime = creationTime;
        return this;
    }
    public Long getCreationTime() {
        return this.creationTime;
    }

    public BatchOperationListJobs setDescription(String description) {
        this.description = description;
        return this;
    }
    public String getDescription() {
        return this.description;
    }

    public BatchOperationListJobs setJobId(String jobId) {
        this.jobId = jobId;
        return this;
    }
    public String getJobId() {
        return this.jobId;
    }

    public BatchOperationListJobs setOperation(String operation) {
        this.operation = operation;
        return this;
    }
    public String getOperation() {
        return this.operation;
    }

    public BatchOperationListJobs setPriority(Integer priority) {
        this.priority = priority;
        return this;
    }
    public Integer getPriority() {
        return this.priority;
    }

    public BatchOperationListJobs setProgressSummary(BatchOperationJobProgressSummary progressSummary) {
        this.progressSummary = progressSummary;
        return this;
    }
    public BatchOperationJobProgressSummary getProgressSummary() {
        return this.progressSummary;
    }

    public BatchOperationListJobs setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

    public BatchOperationListJobs setTerminationDate(Long terminationDate) {
        this.terminationDate = terminationDate;
        return this;
    }
    public Long getTerminationDate() {
        return this.terminationDate;
    }

}
