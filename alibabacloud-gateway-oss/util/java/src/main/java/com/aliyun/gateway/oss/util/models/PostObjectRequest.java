// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PostObjectRequest extends TeaModel {
    @NameInMap("key")
    public String key;

    public static PostObjectRequest build(java.util.Map<String, ?> map) throws Exception {
        PostObjectRequest self = new PostObjectRequest();
        return TeaModel.build(map, self);
    }

    public PostObjectRequest setKey(String key) {
        this.key = key;
        return this;
    }
    public String getKey() {
        return this.key;
    }

}
