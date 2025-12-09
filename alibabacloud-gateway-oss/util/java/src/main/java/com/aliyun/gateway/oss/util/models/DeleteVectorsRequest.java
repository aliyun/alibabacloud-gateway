// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DeleteVectorsRequest extends TeaModel {
    @NameInMap("indexName")
    public String indexName;

    @NameInMap("kyes")
    public java.util.List<String> kyes;

    public static DeleteVectorsRequest build(java.util.Map<String, ?> map) throws Exception {
        DeleteVectorsRequest self = new DeleteVectorsRequest();
        return TeaModel.build(map, self);
    }

    public DeleteVectorsRequest setIndexName(String indexName) {
        this.indexName = indexName;
        return this;
    }
    public String getIndexName() {
        return this.indexName;
    }

    public DeleteVectorsRequest setKyes(java.util.List<String> kyes) {
        this.kyes = kyes;
        return this;
    }
    public java.util.List<String> getKyes() {
        return this.kyes;
    }

}
