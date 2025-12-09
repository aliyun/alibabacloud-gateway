// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BatchOperationReport extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>bucketname</p>
     */
    @NameInMap("Bucket")
    public String bucket;

    /**
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("Enabled")
    public Boolean enabled;

    /**
     * <strong>example:</strong>
     * <p>csv</p>
     */
    @NameInMap("Format")
    public String format;

    /**
     * <strong>example:</strong>
     * <p>/</p>
     */
    @NameInMap("Prefix")
    public String prefix;

    /**
     * <strong>example:</strong>
     * <p>AllTasks</p>
     */
    @NameInMap("ReportScope")
    public String reportScope;

    public static BatchOperationReport build(java.util.Map<String, ?> map) throws Exception {
        BatchOperationReport self = new BatchOperationReport();
        return TeaModel.build(map, self);
    }

    public BatchOperationReport setBucket(String bucket) {
        this.bucket = bucket;
        return this;
    }
    public String getBucket() {
        return this.bucket;
    }

    public BatchOperationReport setEnabled(Boolean enabled) {
        this.enabled = enabled;
        return this;
    }
    public Boolean getEnabled() {
        return this.enabled;
    }

    public BatchOperationReport setFormat(String format) {
        this.format = format;
        return this;
    }
    public String getFormat() {
        return this.format;
    }

    public BatchOperationReport setPrefix(String prefix) {
        this.prefix = prefix;
        return this;
    }
    public String getPrefix() {
        return this.prefix;
    }

    public BatchOperationReport setReportScope(String reportScope) {
        this.reportScope = reportScope;
        return this;
    }
    public String getReportScope() {
        return this.reportScope;
    }

}
