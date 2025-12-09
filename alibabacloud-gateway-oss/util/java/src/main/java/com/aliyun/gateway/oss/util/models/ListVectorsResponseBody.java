// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListVectorsResponseBody extends TeaModel {
    @NameInMap("nextToken")
    public String nextToken;

    @NameInMap("vectors")
    public java.util.List<Vector> vectors;

    public static ListVectorsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListVectorsResponseBody self = new ListVectorsResponseBody();
        return TeaModel.build(map, self);
    }

    public ListVectorsResponseBody setNextToken(String nextToken) {
        this.nextToken = nextToken;
        return this;
    }
    public String getNextToken() {
        return this.nextToken;
    }

    public ListVectorsResponseBody setVectors(java.util.List<Vector> vectors) {
        this.vectors = vectors;
        return this;
    }
    public java.util.List<Vector> getVectors() {
        return this.vectors;
    }

}
