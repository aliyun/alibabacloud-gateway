// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class StopDataLakeStorageTransferJobRequest extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("x-oss-datalake-job-id")
    public String xOssDatalakeJobId;

    public static StopDataLakeStorageTransferJobRequest build(java.util.Map<String, ?> map) throws Exception {
        StopDataLakeStorageTransferJobRequest self = new StopDataLakeStorageTransferJobRequest();
        return TeaModel.build(map, self);
    }

    public StopDataLakeStorageTransferJobRequest setXOssDatalakeJobId(String xOssDatalakeJobId) {
        this.xOssDatalakeJobId = xOssDatalakeJobId;
        return this;
    }
    public String getXOssDatalakeJobId() {
        return this.xOssDatalakeJobId;
    }

}
