// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListDataLakeStorageTransferJobResponseBody extends TeaModel {
    @NameInMap("DataLakeStorageTransferJobs")
    public DataLakeStorageTransferJobs dataLakeStorageTransferJobs;

    public static ListDataLakeStorageTransferJobResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListDataLakeStorageTransferJobResponseBody self = new ListDataLakeStorageTransferJobResponseBody();
        return TeaModel.build(map, self);
    }

    public ListDataLakeStorageTransferJobResponseBody setDataLakeStorageTransferJobs(DataLakeStorageTransferJobs dataLakeStorageTransferJobs) {
        this.dataLakeStorageTransferJobs = dataLakeStorageTransferJobs;
        return this;
    }
    public DataLakeStorageTransferJobs getDataLakeStorageTransferJobs() {
        return this.dataLakeStorageTransferJobs;
    }

}
