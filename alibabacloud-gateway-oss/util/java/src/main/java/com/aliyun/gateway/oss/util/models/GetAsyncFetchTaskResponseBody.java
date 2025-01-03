// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetAsyncFetchTaskResponseBody extends TeaModel {
    @NameInMap("AsyncFetchTaskInfo")
    public AsyncFetchTaskInfo asyncFetchTaskInfo;

    public static GetAsyncFetchTaskResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetAsyncFetchTaskResponseBody self = new GetAsyncFetchTaskResponseBody();
        return TeaModel.build(map, self);
    }

    public GetAsyncFetchTaskResponseBody setAsyncFetchTaskInfo(AsyncFetchTaskInfo asyncFetchTaskInfo) {
        this.asyncFetchTaskInfo = asyncFetchTaskInfo;
        return this;
    }
    public AsyncFetchTaskInfo getAsyncFetchTaskInfo() {
        return this.asyncFetchTaskInfo;
    }

}
