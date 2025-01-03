// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CopyObjectsRequest extends TeaModel {
    @NameInMap("Copy")
    public CopyObjectsCopy copy;

    public static CopyObjectsRequest build(java.util.Map<String, ?> map) throws Exception {
        CopyObjectsRequest self = new CopyObjectsRequest();
        return TeaModel.build(map, self);
    }

    public CopyObjectsRequest setCopy(CopyObjectsCopy copy) {
        this.copy = copy;
        return this;
    }
    public CopyObjectsCopy getCopy() {
        return this.copy;
    }

}
