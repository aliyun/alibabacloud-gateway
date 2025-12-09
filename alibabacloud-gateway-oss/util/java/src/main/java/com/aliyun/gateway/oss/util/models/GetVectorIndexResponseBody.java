// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetVectorIndexResponseBody extends TeaModel {
    @NameInMap("index")
    public VectorIndex index;

    public static GetVectorIndexResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetVectorIndexResponseBody self = new GetVectorIndexResponseBody();
        return TeaModel.build(map, self);
    }

    public GetVectorIndexResponseBody setIndex(VectorIndex index) {
        this.index = index;
        return this;
    }
    public VectorIndex getIndex() {
        return this.index;
    }

}
