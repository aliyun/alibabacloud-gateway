// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BatchCreateJobRequest extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>123456789012****</p>
     */
    @NameInMap("ClientRequestToken")
    public String clientRequestToken;

    /**
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("ConfirmationRequired")
    public Boolean confirmationRequired;

    @NameInMap("Description")
    public String description;

    @NameInMap("KeyPrefixManifestGenerator")
    public BatchOperationManifestGenerator keyPrefixManifestGenerator;

    @NameInMap("Manifest")
    public BatchOperationManifest manifest;

    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("Operation")
    public BatchOperation operation;

    /**
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>5</p>
     */
    @NameInMap("Priority")
    public Integer priority;

    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("Report")
    public BatchOperationReport report;

    /**
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>acs:ram::123456789012****:root</p>
     */
    @NameInMap("RoleArn")
    public String roleArn;

    public static BatchCreateJobRequest build(java.util.Map<String, ?> map) throws Exception {
        BatchCreateJobRequest self = new BatchCreateJobRequest();
        return TeaModel.build(map, self);
    }

    public BatchCreateJobRequest setClientRequestToken(String clientRequestToken) {
        this.clientRequestToken = clientRequestToken;
        return this;
    }
    public String getClientRequestToken() {
        return this.clientRequestToken;
    }

    public BatchCreateJobRequest setConfirmationRequired(Boolean confirmationRequired) {
        this.confirmationRequired = confirmationRequired;
        return this;
    }
    public Boolean getConfirmationRequired() {
        return this.confirmationRequired;
    }

    public BatchCreateJobRequest setDescription(String description) {
        this.description = description;
        return this;
    }
    public String getDescription() {
        return this.description;
    }

    public BatchCreateJobRequest setKeyPrefixManifestGenerator(BatchOperationManifestGenerator keyPrefixManifestGenerator) {
        this.keyPrefixManifestGenerator = keyPrefixManifestGenerator;
        return this;
    }
    public BatchOperationManifestGenerator getKeyPrefixManifestGenerator() {
        return this.keyPrefixManifestGenerator;
    }

    public BatchCreateJobRequest setManifest(BatchOperationManifest manifest) {
        this.manifest = manifest;
        return this;
    }
    public BatchOperationManifest getManifest() {
        return this.manifest;
    }

    public BatchCreateJobRequest setOperation(BatchOperation operation) {
        this.operation = operation;
        return this;
    }
    public BatchOperation getOperation() {
        return this.operation;
    }

    public BatchCreateJobRequest setPriority(Integer priority) {
        this.priority = priority;
        return this;
    }
    public Integer getPriority() {
        return this.priority;
    }

    public BatchCreateJobRequest setReport(BatchOperationReport report) {
        this.report = report;
        return this;
    }
    public BatchOperationReport getReport() {
        return this.report;
    }

    public BatchCreateJobRequest setRoleArn(String roleArn) {
        this.roleArn = roleArn;
        return this;
    }
    public String getRoleArn() {
        return this.roleArn;
    }

}
