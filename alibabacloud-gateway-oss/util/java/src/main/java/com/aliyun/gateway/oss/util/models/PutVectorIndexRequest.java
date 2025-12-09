// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutVectorIndexRequest extends TeaModel {
    @NameInMap("body")
    public VectorIndexPut body;

    public static PutVectorIndexRequest build(java.util.Map<String, ?> map) throws Exception {
        PutVectorIndexRequest self = new PutVectorIndexRequest();
        return TeaModel.build(map, self);
    }

    public PutVectorIndexRequest setBody(VectorIndexPut body) {
        this.body = body;
        return this;
    }
    public VectorIndexPut getBody() {
        return this.body;
    }

}
