// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class VectorIndex extends TeaModel {
    /**
     * <p>Use the UTC time format: yyyy-MM-ddTHH:mmZ</p>
     * 
     * <strong>example:</strong>
     * <p>2025-08-31T10:56:21.000Z</p>
     */
    @NameInMap("createTime")
    public String createTime;

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

    /**
     * <strong>example:</strong>
     * <p>running</p>
     */
    @NameInMap("status")
    public String status;

    /**
     * <strong>example:</strong>
     * <p>test-vector-bucket</p>
     */
    @NameInMap("vectorBucketName")
    public String vectorBucketName;

    public static VectorIndex build(java.util.Map<String, ?> map) throws Exception {
        VectorIndex self = new VectorIndex();
        return TeaModel.build(map, self);
    }

    public VectorIndex setCreateTime(String createTime) {
        this.createTime = createTime;
        return this;
    }
    public String getCreateTime() {
        return this.createTime;
    }

    public VectorIndex setDataType(String dataType) {
        this.dataType = dataType;
        return this;
    }
    public String getDataType() {
        return this.dataType;
    }

    public VectorIndex setDimension(Long dimension) {
        this.dimension = dimension;
        return this;
    }
    public Long getDimension() {
        return this.dimension;
    }

    public VectorIndex setDistanceMetric(String distanceMetric) {
        this.distanceMetric = distanceMetric;
        return this;
    }
    public String getDistanceMetric() {
        return this.distanceMetric;
    }

    public VectorIndex setIndexName(String indexName) {
        this.indexName = indexName;
        return this;
    }
    public String getIndexName() {
        return this.indexName;
    }

    public VectorIndex setMetadata(VectorIndexMetaData metadata) {
        this.metadata = metadata;
        return this;
    }
    public VectorIndexMetaData getMetadata() {
        return this.metadata;
    }

    public VectorIndex setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

    public VectorIndex setVectorBucketName(String vectorBucketName) {
        this.vectorBucketName = vectorBucketName;
        return this;
    }
    public String getVectorBucketName() {
        return this.vectorBucketName;
    }

}
