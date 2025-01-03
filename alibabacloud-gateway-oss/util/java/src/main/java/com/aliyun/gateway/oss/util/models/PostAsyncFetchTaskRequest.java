// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PostAsyncFetchTaskRequest extends TeaModel {
    @NameInMap("AsyncFetchTaskConfiguration")
    public AsyncFetchTaskConfiguration asyncFetchTaskConfiguration;

    public static PostAsyncFetchTaskRequest build(java.util.Map<String, ?> map) throws Exception {
        PostAsyncFetchTaskRequest self = new PostAsyncFetchTaskRequest();
        return TeaModel.build(map, self);
    }

    public PostAsyncFetchTaskRequest setAsyncFetchTaskConfiguration(AsyncFetchTaskConfiguration asyncFetchTaskConfiguration) {
        this.asyncFetchTaskConfiguration = asyncFetchTaskConfiguration;
        return this;
    }
    public AsyncFetchTaskConfiguration getAsyncFetchTaskConfiguration() {
        return this.asyncFetchTaskConfiguration;
    }

}
