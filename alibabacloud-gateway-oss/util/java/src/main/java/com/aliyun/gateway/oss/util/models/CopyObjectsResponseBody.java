// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CopyObjectsResponseBody extends TeaModel {
    @NameInMap("CopyObjectsResult")
    public CopyObjectsResult copyObjectsResult;

    public static CopyObjectsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        CopyObjectsResponseBody self = new CopyObjectsResponseBody();
        return TeaModel.build(map, self);
    }

    public CopyObjectsResponseBody setCopyObjectsResult(CopyObjectsResult copyObjectsResult) {
        this.copyObjectsResult = copyObjectsResult;
        return this;
    }
    public CopyObjectsResult getCopyObjectsResult() {
        return this.copyObjectsResult;
    }

}
