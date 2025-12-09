// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetVectorsRequest extends TeaModel {
    @NameInMap("indexName")
    public String indexName;

    @NameInMap("keys")
    public java.util.List<String> keys;

    @NameInMap("returnData")
    public Boolean returnData;

    @NameInMap("returnMetadata")
    public Boolean returnMetadata;

    public static GetVectorsRequest build(java.util.Map<String, ?> map) throws Exception {
        GetVectorsRequest self = new GetVectorsRequest();
        return TeaModel.build(map, self);
    }

    public GetVectorsRequest setIndexName(String indexName) {
        this.indexName = indexName;
        return this;
    }
    public String getIndexName() {
        return this.indexName;
    }

    public GetVectorsRequest setKeys(java.util.List<String> keys) {
        this.keys = keys;
        return this;
    }
    public java.util.List<String> getKeys() {
        return this.keys;
    }

    public GetVectorsRequest setReturnData(Boolean returnData) {
        this.returnData = returnData;
        return this;
    }
    public Boolean getReturnData() {
        return this.returnData;
    }

    public GetVectorsRequest setReturnMetadata(Boolean returnMetadata) {
        this.returnMetadata = returnMetadata;
        return this;
    }
    public Boolean getReturnMetadata() {
        return this.returnMetadata;
    }

}
