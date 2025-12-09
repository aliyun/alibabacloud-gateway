// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetVectorsResponseBody extends TeaModel {
    @NameInMap("vectors")
    public java.util.List<Vector> vectors;

    public static GetVectorsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetVectorsResponseBody self = new GetVectorsResponseBody();
        return TeaModel.build(map, self);
    }

    public GetVectorsResponseBody setVectors(java.util.List<Vector> vectors) {
        this.vectors = vectors;
        return this;
    }
    public java.util.List<Vector> getVectors() {
        return this.vectors;
    }

}
