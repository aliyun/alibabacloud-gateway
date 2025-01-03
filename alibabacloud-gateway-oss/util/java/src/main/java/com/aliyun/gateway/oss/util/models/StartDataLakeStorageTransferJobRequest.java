// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class StartDataLakeStorageTransferJobRequest extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("x-oss-datalake-job-id")
    public String xOssDatalakeJobId;

    public static StartDataLakeStorageTransferJobRequest build(java.util.Map<String, ?> map) throws Exception {
        StartDataLakeStorageTransferJobRequest self = new StartDataLakeStorageTransferJobRequest();
        return TeaModel.build(map, self);
    }

    public StartDataLakeStorageTransferJobRequest setXOssDatalakeJobId(String xOssDatalakeJobId) {
        this.xOssDatalakeJobId = xOssDatalakeJobId;
        return this;
    }
    public String getXOssDatalakeJobId() {
        return this.xOssDatalakeJobId;
    }

}
