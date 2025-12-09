// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateJobResult extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>IDBkODc3M2RlZjgyNjQ0NDViZGV****</p>
     */
    @NameInMap("JobId")
    public String jobId;

    public static CreateJobResult build(java.util.Map<String, ?> map) throws Exception {
        CreateJobResult self = new CreateJobResult();
        return TeaModel.build(map, self);
    }

    public CreateJobResult setJobId(String jobId) {
        this.jobId = jobId;
        return this;
    }
    public String getJobId() {
        return this.jobId;
    }

}
