// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetVectorIndexRequest extends TeaModel {
    @NameInMap("indexName")
    public String indexName;

    public static GetVectorIndexRequest build(java.util.Map<String, ?> map) throws Exception {
        GetVectorIndexRequest self = new GetVectorIndexRequest();
        return TeaModel.build(map, self);
    }

    public GetVectorIndexRequest setIndexName(String indexName) {
        this.indexName = indexName;
        return this;
    }
    public String getIndexName() {
        return this.indexName;
    }

}
