// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PromoteDataLakeCacheRequest extends TeaModel {
    @NameInMap("PromoteDataLakeCacheRequest")
    public PromoteDataLakeCacheReq promoteDataLakeCacheRequest;

    public static PromoteDataLakeCacheRequest build(java.util.Map<String, ?> map) throws Exception {
        PromoteDataLakeCacheRequest self = new PromoteDataLakeCacheRequest();
        return TeaModel.build(map, self);
    }

    public PromoteDataLakeCacheRequest setPromoteDataLakeCacheRequest(PromoteDataLakeCacheReq promoteDataLakeCacheRequest) {
        this.promoteDataLakeCacheRequest = promoteDataLakeCacheRequest;
        return this;
    }
    public PromoteDataLakeCacheReq getPromoteDataLakeCacheRequest() {
        return this.promoteDataLakeCacheRequest;
    }

}
