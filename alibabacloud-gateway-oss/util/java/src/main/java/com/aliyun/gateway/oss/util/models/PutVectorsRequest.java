// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutVectorsRequest extends TeaModel {
    @NameInMap("indexName")
    public String indexName;

    @NameInMap("vectors")
    public java.util.List<Vector> vectors;

    public static PutVectorsRequest build(java.util.Map<String, ?> map) throws Exception {
        PutVectorsRequest self = new PutVectorsRequest();
        return TeaModel.build(map, self);
    }

    public PutVectorsRequest setIndexName(String indexName) {
        this.indexName = indexName;
        return this;
    }
    public String getIndexName() {
        return this.indexName;
    }

    public PutVectorsRequest setVectors(java.util.List<Vector> vectors) {
        this.vectors = vectors;
        return this;
    }
    public java.util.List<Vector> getVectors() {
        return this.vectors;
    }

}
