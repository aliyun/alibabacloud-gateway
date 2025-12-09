// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListVectorsRequest extends TeaModel {
    @NameInMap("indexName")
    public String indexName;

    @NameInMap("maxResults")
    public Long maxResults;

    @NameInMap("nextToken")
    public String nextToken;

    @NameInMap("returnData")
    public Boolean returnData;

    @NameInMap("returnMetadata")
    public Boolean returnMetadata;

    @NameInMap("segmentCount")
    public Long segmentCount;

    @NameInMap("segmentIndex")
    public Long segmentIndex;

    public static ListVectorsRequest build(java.util.Map<String, ?> map) throws Exception {
        ListVectorsRequest self = new ListVectorsRequest();
        return TeaModel.build(map, self);
    }

    public ListVectorsRequest setIndexName(String indexName) {
        this.indexName = indexName;
        return this;
    }
    public String getIndexName() {
        return this.indexName;
    }

    public ListVectorsRequest setMaxResults(Long maxResults) {
        this.maxResults = maxResults;
        return this;
    }
    public Long getMaxResults() {
        return this.maxResults;
    }

    public ListVectorsRequest setNextToken(String nextToken) {
        this.nextToken = nextToken;
        return this;
    }
    public String getNextToken() {
        return this.nextToken;
    }

    public ListVectorsRequest setReturnData(Boolean returnData) {
        this.returnData = returnData;
        return this;
    }
    public Boolean getReturnData() {
        return this.returnData;
    }

    public ListVectorsRequest setReturnMetadata(Boolean returnMetadata) {
        this.returnMetadata = returnMetadata;
        return this;
    }
    public Boolean getReturnMetadata() {
        return this.returnMetadata;
    }

    public ListVectorsRequest setSegmentCount(Long segmentCount) {
        this.segmentCount = segmentCount;
        return this;
    }
    public Long getSegmentCount() {
        return this.segmentCount;
    }

    public ListVectorsRequest setSegmentIndex(Long segmentIndex) {
        this.segmentIndex = segmentIndex;
        return this;
    }
    public Long getSegmentIndex() {
        return this.segmentIndex;
    }

}
