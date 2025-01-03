// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListDataLakeStorageTransferJobHistoryResponseBody extends TeaModel {
    @NameInMap("ListDataLakeStorageTransferJobHistory")
    public ListDataLakeStorageTransferJobHistory listDataLakeStorageTransferJobHistory;

    public static ListDataLakeStorageTransferJobHistoryResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListDataLakeStorageTransferJobHistoryResponseBody self = new ListDataLakeStorageTransferJobHistoryResponseBody();
        return TeaModel.build(map, self);
    }

    public ListDataLakeStorageTransferJobHistoryResponseBody setListDataLakeStorageTransferJobHistory(ListDataLakeStorageTransferJobHistory listDataLakeStorageTransferJobHistory) {
        this.listDataLakeStorageTransferJobHistory = listDataLakeStorageTransferJobHistory;
        return this;
    }
    public ListDataLakeStorageTransferJobHistory getListDataLakeStorageTransferJobHistory() {
        return this.listDataLakeStorageTransferJobHistory;
    }

}
