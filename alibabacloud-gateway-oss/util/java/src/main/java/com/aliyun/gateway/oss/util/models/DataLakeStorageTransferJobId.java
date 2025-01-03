// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DataLakeStorageTransferJobId extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>abcde787e3c04af2af290ec52d123456</p>
     */
    @NameInMap("Id")
    public String id;

    public static DataLakeStorageTransferJobId build(java.util.Map<String, ?> map) throws Exception {
        DataLakeStorageTransferJobId self = new DataLakeStorageTransferJobId();
        return TeaModel.build(map, self);
    }

    public DataLakeStorageTransferJobId setId(String id) {
        this.id = id;
        return this;
    }
    public String getId() {
        return this.id;
    }

}
