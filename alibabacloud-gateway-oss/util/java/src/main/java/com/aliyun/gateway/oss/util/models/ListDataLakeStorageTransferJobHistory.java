// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListDataLakeStorageTransferJobHistory extends TeaModel {
    @NameInMap("DataLakeStorageTransferJobHistory")
    public java.util.List<DataLakeStorageTransferJobHistory> dataLakeStorageTransferJobHistory;

    public static ListDataLakeStorageTransferJobHistory build(java.util.Map<String, ?> map) throws Exception {
        ListDataLakeStorageTransferJobHistory self = new ListDataLakeStorageTransferJobHistory();
        return TeaModel.build(map, self);
    }

    public ListDataLakeStorageTransferJobHistory setDataLakeStorageTransferJobHistory(java.util.List<DataLakeStorageTransferJobHistory> dataLakeStorageTransferJobHistory) {
        this.dataLakeStorageTransferJobHistory = dataLakeStorageTransferJobHistory;
        return this;
    }
    public java.util.List<DataLakeStorageTransferJobHistory> getDataLakeStorageTransferJobHistory() {
        return this.dataLakeStorageTransferJobHistory;
    }

}
