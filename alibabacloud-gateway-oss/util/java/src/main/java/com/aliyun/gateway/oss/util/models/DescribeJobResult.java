// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DescribeJobResult extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>1730250699</p>
     */
    @NameInMap("CreationTime")
    public Long creationTime;

    @NameInMap("Description")
    public String description;

    @NameInMap("FailureReasons")
    public BatchOperationFailureReasons failureReasons;

    /**
     * <strong>example:</strong>
     * <p>IDBkODc3M2RlZjgyNjQ0NDViZGV****</p>
     */
    @NameInMap("JobId")
    public String jobId;

    @NameInMap("KeyPrefixManifestGenerator")
    public BatchOperationManifestGenerator keyPrefixManifestGenerator;

    @NameInMap("Manifest")
    public BatchOperationManifest manifest;

    @NameInMap("Operation")
    public BatchOperation operation;

    /**
     * <strong>example:</strong>
     * <p>5</p>
     */
    @NameInMap("Priority")
    public Integer priority;

    @NameInMap("ProgressSummary")
    public BatchOperationJobProgressSummary progressSummary;

    @NameInMap("Report")
    public BatchOperationReport report;

    /**
     * <strong>example:</strong>
     * <p>acs:ram::174649585760****:role/AliyunOSSRole</p>
     */
    @NameInMap("RoleArn")
    public String roleArn;

    /**
     * <strong>example:</strong>
     * <p>Preparing</p>
     */
    @NameInMap("Status")
    public String status;

    @NameInMap("StatusUpdateReason")
    public String statusUpdateReason;

    /**
     * <strong>example:</strong>
     * <p>1730250699</p>
     */
    @NameInMap("TerminationDate")
    public Long terminationDate;

    public static DescribeJobResult build(java.util.Map<String, ?> map) throws Exception {
        DescribeJobResult self = new DescribeJobResult();
        return TeaModel.build(map, self);
    }

    public DescribeJobResult setCreationTime(Long creationTime) {
        this.creationTime = creationTime;
        return this;
    }
    public Long getCreationTime() {
        return this.creationTime;
    }

    public DescribeJobResult setDescription(String description) {
        this.description = description;
        return this;
    }
    public String getDescription() {
        return this.description;
    }

    public DescribeJobResult setFailureReasons(BatchOperationFailureReasons failureReasons) {
        this.failureReasons = failureReasons;
        return this;
    }
    public BatchOperationFailureReasons getFailureReasons() {
        return this.failureReasons;
    }

    public DescribeJobResult setJobId(String jobId) {
        this.jobId = jobId;
        return this;
    }
    public String getJobId() {
        return this.jobId;
    }

    public DescribeJobResult setKeyPrefixManifestGenerator(BatchOperationManifestGenerator keyPrefixManifestGenerator) {
        this.keyPrefixManifestGenerator = keyPrefixManifestGenerator;
        return this;
    }
    public BatchOperationManifestGenerator getKeyPrefixManifestGenerator() {
        return this.keyPrefixManifestGenerator;
    }

    public DescribeJobResult setManifest(BatchOperationManifest manifest) {
        this.manifest = manifest;
        return this;
    }
    public BatchOperationManifest getManifest() {
        return this.manifest;
    }

    public DescribeJobResult setOperation(BatchOperation operation) {
        this.operation = operation;
        return this;
    }
    public BatchOperation getOperation() {
        return this.operation;
    }

    public DescribeJobResult setPriority(Integer priority) {
        this.priority = priority;
        return this;
    }
    public Integer getPriority() {
        return this.priority;
    }

    public DescribeJobResult setProgressSummary(BatchOperationJobProgressSummary progressSummary) {
        this.progressSummary = progressSummary;
        return this;
    }
    public BatchOperationJobProgressSummary getProgressSummary() {
        return this.progressSummary;
    }

    public DescribeJobResult setReport(BatchOperationReport report) {
        this.report = report;
        return this;
    }
    public BatchOperationReport getReport() {
        return this.report;
    }

    public DescribeJobResult setRoleArn(String roleArn) {
        this.roleArn = roleArn;
        return this;
    }
    public String getRoleArn() {
        return this.roleArn;
    }

    public DescribeJobResult setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

    public DescribeJobResult setStatusUpdateReason(String statusUpdateReason) {
        this.statusUpdateReason = statusUpdateReason;
        return this;
    }
    public String getStatusUpdateReason() {
        return this.statusUpdateReason;
    }

    public DescribeJobResult setTerminationDate(Long terminationDate) {
        this.terminationDate = terminationDate;
        return this;
    }
    public Long getTerminationDate() {
        return this.terminationDate;
    }

}
