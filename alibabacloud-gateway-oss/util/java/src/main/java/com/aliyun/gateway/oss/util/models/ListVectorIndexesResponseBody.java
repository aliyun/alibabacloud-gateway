// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListVectorIndexesResponseBody extends TeaModel {
    @NameInMap("indexes")
    public java.util.List<VectorIndex> indexes;

    @NameInMap("nextToken")
    public String nextToken;

    public static ListVectorIndexesResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListVectorIndexesResponseBody self = new ListVectorIndexesResponseBody();
        return TeaModel.build(map, self);
    }

    public ListVectorIndexesResponseBody setIndexes(java.util.List<VectorIndex> indexes) {
        this.indexes = indexes;
        return this;
    }
    public java.util.List<VectorIndex> getIndexes() {
        return this.indexes;
    }

    public ListVectorIndexesResponseBody setNextToken(String nextToken) {
        this.nextToken = nextToken;
        return this;
    }
    public String getNextToken() {
        return this.nextToken;
    }

}
