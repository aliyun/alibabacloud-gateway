// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetResourcePoolRequesterQoSInfoResponseBody extends TeaModel {
    @NameInMap("RequesterQoSInfo")
    public RequesterQoSInfo requesterQoSInfo;

    public static GetResourcePoolRequesterQoSInfoResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetResourcePoolRequesterQoSInfoResponseBody self = new GetResourcePoolRequesterQoSInfoResponseBody();
        return TeaModel.build(map, self);
    }

    public GetResourcePoolRequesterQoSInfoResponseBody setRequesterQoSInfo(RequesterQoSInfo requesterQoSInfo) {
        this.requesterQoSInfo = requesterQoSInfo;
        return this;
    }
    public RequesterQoSInfo getRequesterQoSInfo() {
        return this.requesterQoSInfo;
    }

}
