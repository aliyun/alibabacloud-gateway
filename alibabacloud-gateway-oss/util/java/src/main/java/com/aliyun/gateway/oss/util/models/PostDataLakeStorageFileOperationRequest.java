// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PostDataLakeStorageFileOperationRequest extends TeaModel {
    @NameInMap("body")
    public String body;

    public static PostDataLakeStorageFileOperationRequest build(java.util.Map<String, ?> map) throws Exception {
        PostDataLakeStorageFileOperationRequest self = new PostDataLakeStorageFileOperationRequest();
        return TeaModel.build(map, self);
    }

    public PostDataLakeStorageFileOperationRequest setBody(String body) {
        this.body = body;
        return this;
    }
    public String getBody() {
        return this.body;
    }

}
