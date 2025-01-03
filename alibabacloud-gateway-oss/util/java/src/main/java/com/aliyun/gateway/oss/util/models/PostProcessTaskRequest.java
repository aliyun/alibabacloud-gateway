// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PostProcessTaskRequest extends TeaModel {
    @NameInMap("body")
    public java.io.InputStream body;

    public static PostProcessTaskRequest build(java.util.Map<String, ?> map) throws Exception {
        PostProcessTaskRequest self = new PostProcessTaskRequest();
        return TeaModel.build(map, self);
    }

    public PostProcessTaskRequest setBody(java.io.InputStream body) {
        this.body = body;
        return this;
    }
    public java.io.InputStream getBody() {
        return this.body;
    }

}
