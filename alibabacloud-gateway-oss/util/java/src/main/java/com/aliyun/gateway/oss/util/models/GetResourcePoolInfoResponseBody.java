// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetResourcePoolInfoResponseBody extends TeaModel {
    @NameInMap("GetResourcePoolInfoResponse")
    public GetResourcePoolInfoResp getResourcePoolInfoResponse;

    public static GetResourcePoolInfoResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetResourcePoolInfoResponseBody self = new GetResourcePoolInfoResponseBody();
        return TeaModel.build(map, self);
    }

    public GetResourcePoolInfoResponseBody setGetResourcePoolInfoResponse(GetResourcePoolInfoResp getResourcePoolInfoResponse) {
        this.getResourcePoolInfoResponse = getResourcePoolInfoResponse;
        return this;
    }
    public GetResourcePoolInfoResp getGetResourcePoolInfoResponse() {
        return this.getResourcePoolInfoResponse;
    }

}
