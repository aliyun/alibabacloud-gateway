// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutDataLakeCachePrefetchJobRequest extends TeaModel {
    @NameInMap("CreateDataLakeCachePrefetchJob")
    public CreateDataLakeCachePrefetchJob createDataLakeCachePrefetchJob;

    @NameInMap("x-oss-datalake-job-id")
    public String xOssDatalakeJobId;

    public static PutDataLakeCachePrefetchJobRequest build(java.util.Map<String, ?> map) throws Exception {
        PutDataLakeCachePrefetchJobRequest self = new PutDataLakeCachePrefetchJobRequest();
        return TeaModel.build(map, self);
    }

    public PutDataLakeCachePrefetchJobRequest setCreateDataLakeCachePrefetchJob(CreateDataLakeCachePrefetchJob createDataLakeCachePrefetchJob) {
        this.createDataLakeCachePrefetchJob = createDataLakeCachePrefetchJob;
        return this;
    }
    public CreateDataLakeCachePrefetchJob getCreateDataLakeCachePrefetchJob() {
        return this.createDataLakeCachePrefetchJob;
    }

    public PutDataLakeCachePrefetchJobRequest setXOssDatalakeJobId(String xOssDatalakeJobId) {
        this.xOssDatalakeJobId = xOssDatalakeJobId;
        return this;
    }
    public String getXOssDatalakeJobId() {
        return this.xOssDatalakeJobId;
    }

}
