// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetDataLakeStorageTransferJobRequest extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("x-oss-datalake-job-id")
    public String xOssDatalakeJobId;

    @NameInMap("x-oss-datalake-job-progress")
    public String xOssDatalakeJobProgress;

    public static GetDataLakeStorageTransferJobRequest build(java.util.Map<String, ?> map) throws Exception {
        GetDataLakeStorageTransferJobRequest self = new GetDataLakeStorageTransferJobRequest();
        return TeaModel.build(map, self);
    }

    public GetDataLakeStorageTransferJobRequest setXOssDatalakeJobId(String xOssDatalakeJobId) {
        this.xOssDatalakeJobId = xOssDatalakeJobId;
        return this;
    }
    public String getXOssDatalakeJobId() {
        return this.xOssDatalakeJobId;
    }

    public GetDataLakeStorageTransferJobRequest setXOssDatalakeJobProgress(String xOssDatalakeJobProgress) {
        this.xOssDatalakeJobProgress = xOssDatalakeJobProgress;
        return this;
    }
    public String getXOssDatalakeJobProgress() {
        return this.xOssDatalakeJobProgress;
    }

}
