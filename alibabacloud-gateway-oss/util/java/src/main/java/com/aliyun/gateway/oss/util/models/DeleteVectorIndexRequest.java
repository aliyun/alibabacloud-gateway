// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DeleteVectorIndexRequest extends TeaModel {
    @NameInMap("indexName")
    public String indexName;

    public static DeleteVectorIndexRequest build(java.util.Map<String, ?> map) throws Exception {
        DeleteVectorIndexRequest self = new DeleteVectorIndexRequest();
        return TeaModel.build(map, self);
    }

    public DeleteVectorIndexRequest setIndexName(String indexName) {
        this.indexName = indexName;
        return this;
    }
    public String getIndexName() {
        return this.indexName;
    }

}
