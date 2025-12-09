// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class VectorIndexPut extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>float32</p>
     */
    @NameInMap("dataType")
    public String dataType;

    /**
     * <strong>example:</strong>
     * <p>100</p>
     */
    @NameInMap("dimension")
    public Long dimension;

    /**
     * <strong>example:</strong>
     * <p>euclidean</p>
     */
    @NameInMap("distanceMetric")
    public String distanceMetric;

    /**
     * <strong>example:</strong>
     * <p>test-vector</p>
     */
    @NameInMap("indexName")
    public String indexName;

    @NameInMap("metadata")
    public VectorIndexMetaData metadata;

    public static VectorIndexPut build(java.util.Map<String, ?> map) throws Exception {
        VectorIndexPut self = new VectorIndexPut();
        return TeaModel.build(map, self);
    }

    public VectorIndexPut setDataType(String dataType) {
        this.dataType = dataType;
        return this;
    }
    public String getDataType() {
        return this.dataType;
    }

    public VectorIndexPut setDimension(Long dimension) {
        this.dimension = dimension;
        return this;
    }
    public Long getDimension() {
        return this.dimension;
    }

    public VectorIndexPut setDistanceMetric(String distanceMetric) {
        this.distanceMetric = distanceMetric;
        return this;
    }
    public String getDistanceMetric() {
        return this.distanceMetric;
    }

    public VectorIndexPut setIndexName(String indexName) {
        this.indexName = indexName;
        return this;
    }
    public String getIndexName() {
        return this.indexName;
    }

    public VectorIndexPut setMetadata(VectorIndexMetaData metadata) {
        this.metadata = metadata;
        return this;
    }
    public VectorIndexMetaData getMetadata() {
        return this.metadata;
    }

}
