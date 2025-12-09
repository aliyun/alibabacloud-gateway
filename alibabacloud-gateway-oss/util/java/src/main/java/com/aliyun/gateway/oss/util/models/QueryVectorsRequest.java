// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class QueryVectorsRequest extends TeaModel {
    @NameInMap("filter")
    public Object filter;

    @NameInMap("indexName")
    public String indexName;

    @NameInMap("queryVector")
    public VectorData queryVector;

    @NameInMap("returnDistance")
    public Boolean returnDistance;

    @NameInMap("returnMetadata")
    public Boolean returnMetadata;

    @NameInMap("topK")
    public Long topK;

    public static QueryVectorsRequest build(java.util.Map<String, ?> map) throws Exception {
        QueryVectorsRequest self = new QueryVectorsRequest();
        return TeaModel.build(map, self);
    }

    public QueryVectorsRequest setFilter(Object filter) {
        this.filter = filter;
        return this;
    }
    public Object getFilter() {
        return this.filter;
    }

    public QueryVectorsRequest setIndexName(String indexName) {
        this.indexName = indexName;
        return this;
    }
    public String getIndexName() {
        return this.indexName;
    }

    public QueryVectorsRequest setQueryVector(VectorData queryVector) {
        this.queryVector = queryVector;
        return this;
    }
    public VectorData getQueryVector() {
        return this.queryVector;
    }

    public QueryVectorsRequest setReturnDistance(Boolean returnDistance) {
        this.returnDistance = returnDistance;
        return this;
    }
    public Boolean getReturnDistance() {
        return this.returnDistance;
    }

    public QueryVectorsRequest setReturnMetadata(Boolean returnMetadata) {
        this.returnMetadata = returnMetadata;
        return this;
    }
    public Boolean getReturnMetadata() {
        return this.returnMetadata;
    }

    public QueryVectorsRequest setTopK(Long topK) {
        this.topK = topK;
        return this;
    }
    public Long getTopK() {
        return this.topK;
    }

}
