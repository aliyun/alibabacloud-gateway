// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DataLakeStorageTransferJobHistoryId extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>abcdef03370a4b32ac6bc69eb1123456</p>
     */
    @NameInMap("HistoryId")
    public String historyId;

    public static DataLakeStorageTransferJobHistoryId build(java.util.Map<String, ?> map) throws Exception {
        DataLakeStorageTransferJobHistoryId self = new DataLakeStorageTransferJobHistoryId();
        return TeaModel.build(map, self);
    }

    public DataLakeStorageTransferJobHistoryId setHistoryId(String historyId) {
        this.historyId = historyId;
        return this;
    }
    public String getHistoryId() {
        return this.historyId;
    }

}
