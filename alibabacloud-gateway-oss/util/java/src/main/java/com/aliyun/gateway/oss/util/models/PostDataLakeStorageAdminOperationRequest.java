// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PostDataLakeStorageAdminOperationRequest extends TeaModel {
    @NameInMap("body")
    public String body;

    public static PostDataLakeStorageAdminOperationRequest build(java.util.Map<String, ?> map) throws Exception {
        PostDataLakeStorageAdminOperationRequest self = new PostDataLakeStorageAdminOperationRequest();
        return TeaModel.build(map, self);
    }

    public PostDataLakeStorageAdminOperationRequest setBody(String body) {
        this.body = body;
        return this;
    }
    public String getBody() {
        return this.body;
    }

}
