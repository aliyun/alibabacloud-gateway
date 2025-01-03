// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetDataLakeStorageTransferJobResponseBody extends TeaModel {
    @NameInMap("DataLakeStorageTransferJob")
    public DataLakeStorageTransferJob dataLakeStorageTransferJob;

    public static GetDataLakeStorageTransferJobResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetDataLakeStorageTransferJobResponseBody self = new GetDataLakeStorageTransferJobResponseBody();
        return TeaModel.build(map, self);
    }

    public GetDataLakeStorageTransferJobResponseBody setDataLakeStorageTransferJob(DataLakeStorageTransferJob dataLakeStorageTransferJob) {
        this.dataLakeStorageTransferJob = dataLakeStorageTransferJob;
        return this;
    }
    public DataLakeStorageTransferJob getDataLakeStorageTransferJob() {
        return this.dataLakeStorageTransferJob;
    }

}
