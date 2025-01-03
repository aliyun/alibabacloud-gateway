// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListDataLakeStorageTransferJobHistoryRequest extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("x-oss-datalake-job-id")
    public String xOssDatalakeJobId;

    public static ListDataLakeStorageTransferJobHistoryRequest build(java.util.Map<String, ?> map) throws Exception {
        ListDataLakeStorageTransferJobHistoryRequest self = new ListDataLakeStorageTransferJobHistoryRequest();
        return TeaModel.build(map, self);
    }

    public ListDataLakeStorageTransferJobHistoryRequest setXOssDatalakeJobId(String xOssDatalakeJobId) {
        this.xOssDatalakeJobId = xOssDatalakeJobId;
        return this;
    }
    public String getXOssDatalakeJobId() {
        return this.xOssDatalakeJobId;
    }

}
