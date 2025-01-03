// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class StartDataLakeStorageTransferJobResponseBody extends TeaModel {
    @NameInMap("DataLakeStorageTransferJobHistoryId")
    public DataLakeStorageTransferJobHistoryId dataLakeStorageTransferJobHistoryId;

    public static StartDataLakeStorageTransferJobResponseBody build(java.util.Map<String, ?> map) throws Exception {
        StartDataLakeStorageTransferJobResponseBody self = new StartDataLakeStorageTransferJobResponseBody();
        return TeaModel.build(map, self);
    }

    public StartDataLakeStorageTransferJobResponseBody setDataLakeStorageTransferJobHistoryId(DataLakeStorageTransferJobHistoryId dataLakeStorageTransferJobHistoryId) {
        this.dataLakeStorageTransferJobHistoryId = dataLakeStorageTransferJobHistoryId;
        return this;
    }
    public DataLakeStorageTransferJobHistoryId getDataLakeStorageTransferJobHistoryId() {
        return this.dataLakeStorageTransferJobHistoryId;
    }

}
