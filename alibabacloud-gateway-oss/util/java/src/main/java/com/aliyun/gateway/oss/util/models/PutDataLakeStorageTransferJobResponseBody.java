// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutDataLakeStorageTransferJobResponseBody extends TeaModel {
    @NameInMap("DataLakeStorageTransferJobId")
    public DataLakeStorageTransferJobId dataLakeStorageTransferJobId;

    public static PutDataLakeStorageTransferJobResponseBody build(java.util.Map<String, ?> map) throws Exception {
        PutDataLakeStorageTransferJobResponseBody self = new PutDataLakeStorageTransferJobResponseBody();
        return TeaModel.build(map, self);
    }

    public PutDataLakeStorageTransferJobResponseBody setDataLakeStorageTransferJobId(DataLakeStorageTransferJobId dataLakeStorageTransferJobId) {
        this.dataLakeStorageTransferJobId = dataLakeStorageTransferJobId;
        return this;
    }
    public DataLakeStorageTransferJobId getDataLakeStorageTransferJobId() {
        return this.dataLakeStorageTransferJobId;
    }

}
