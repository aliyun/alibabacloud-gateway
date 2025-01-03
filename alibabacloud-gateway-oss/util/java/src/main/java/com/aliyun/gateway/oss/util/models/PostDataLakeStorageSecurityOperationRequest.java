// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PostDataLakeStorageSecurityOperationRequest extends TeaModel {
    @NameInMap("body")
    public String body;

    public static PostDataLakeStorageSecurityOperationRequest build(java.util.Map<String, ?> map) throws Exception {
        PostDataLakeStorageSecurityOperationRequest self = new PostDataLakeStorageSecurityOperationRequest();
        return TeaModel.build(map, self);
    }

    public PostDataLakeStorageSecurityOperationRequest setBody(String body) {
        this.body = body;
        return this;
    }
    public String getBody() {
        return this.body;
    }

}
