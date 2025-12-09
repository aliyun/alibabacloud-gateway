// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BatchPutObjectAcl extends TeaModel {
    @NameInMap("ObjectAcl")
    public String objectAcl;

    public static BatchPutObjectAcl build(java.util.Map<String, ?> map) throws Exception {
        BatchPutObjectAcl self = new BatchPutObjectAcl();
        return TeaModel.build(map, self);
    }

    public BatchPutObjectAcl setObjectAcl(String objectAcl) {
        this.objectAcl = objectAcl;
        return this;
    }
    public String getObjectAcl() {
        return this.objectAcl;
    }

}
