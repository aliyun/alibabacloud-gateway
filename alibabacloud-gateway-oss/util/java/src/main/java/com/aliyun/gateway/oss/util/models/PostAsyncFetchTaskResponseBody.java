// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PostAsyncFetchTaskResponseBody extends TeaModel {
    @NameInMap("AsyncFetchTaskResult")
    public AsyncFetchTaskResult asyncFetchTaskResult;

    public static PostAsyncFetchTaskResponseBody build(java.util.Map<String, ?> map) throws Exception {
        PostAsyncFetchTaskResponseBody self = new PostAsyncFetchTaskResponseBody();
        return TeaModel.build(map, self);
    }

    public PostAsyncFetchTaskResponseBody setAsyncFetchTaskResult(AsyncFetchTaskResult asyncFetchTaskResult) {
        this.asyncFetchTaskResult = asyncFetchTaskResult;
        return this;
    }
    public AsyncFetchTaskResult getAsyncFetchTaskResult() {
        return this.asyncFetchTaskResult;
    }

}
