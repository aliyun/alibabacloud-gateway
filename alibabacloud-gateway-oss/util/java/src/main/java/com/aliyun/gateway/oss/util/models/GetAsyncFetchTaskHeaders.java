// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetAsyncFetchTaskHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("x-oss-task-id")
    public String xOssTaskId;

    public static GetAsyncFetchTaskHeaders build(java.util.Map<String, ?> map) throws Exception {
        GetAsyncFetchTaskHeaders self = new GetAsyncFetchTaskHeaders();
        return TeaModel.build(map, self);
    }

    public GetAsyncFetchTaskHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public GetAsyncFetchTaskHeaders setXOssTaskId(String xOssTaskId) {
        this.xOssTaskId = xOssTaskId;
        return this;
    }
    public String getXOssTaskId() {
        return this.xOssTaskId;
    }

}
