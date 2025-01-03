// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutDataLakeStorageTransferJobRequest extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("CreateDataLakeStorageTransferJob")
    public CreateDataLakeStorageTransferJob createDataLakeStorageTransferJob;

    @NameInMap("x-oss-datalake-job-id")
    public String xOssDatalakeJobId;

    public static PutDataLakeStorageTransferJobRequest build(java.util.Map<String, ?> map) throws Exception {
        PutDataLakeStorageTransferJobRequest self = new PutDataLakeStorageTransferJobRequest();
        return TeaModel.build(map, self);
    }

    public PutDataLakeStorageTransferJobRequest setCreateDataLakeStorageTransferJob(CreateDataLakeStorageTransferJob createDataLakeStorageTransferJob) {
        this.createDataLakeStorageTransferJob = createDataLakeStorageTransferJob;
        return this;
    }
    public CreateDataLakeStorageTransferJob getCreateDataLakeStorageTransferJob() {
        return this.createDataLakeStorageTransferJob;
    }

    public PutDataLakeStorageTransferJobRequest setXOssDatalakeJobId(String xOssDatalakeJobId) {
        this.xOssDatalakeJobId = xOssDatalakeJobId;
        return this;
    }
    public String getXOssDatalakeJobId() {
        return this.xOssDatalakeJobId;
    }

}
