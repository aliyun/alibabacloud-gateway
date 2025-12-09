// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class QueryVectorsResponseBody extends TeaModel {
    @NameInMap("vectors")
    public java.util.List<QueriedVector> vectors;

    public static QueryVectorsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        QueryVectorsResponseBody self = new QueryVectorsResponseBody();
        return TeaModel.build(map, self);
    }

    public QueryVectorsResponseBody setVectors(java.util.List<QueriedVector> vectors) {
        this.vectors = vectors;
        return this;
    }
    public java.util.List<QueriedVector> getVectors() {
        return this.vectors;
    }

}
