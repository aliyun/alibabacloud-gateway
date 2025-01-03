// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketRequesterQoSInfoResponseBody extends TeaModel {
    @NameInMap("RequesterQoSInfo")
    public RequesterQoSInfo requesterQoSInfo;

    public static GetBucketRequesterQoSInfoResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketRequesterQoSInfoResponseBody self = new GetBucketRequesterQoSInfoResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketRequesterQoSInfoResponseBody setRequesterQoSInfo(RequesterQoSInfo requesterQoSInfo) {
        this.requesterQoSInfo = requesterQoSInfo;
        return this;
    }
    public RequesterQoSInfo getRequesterQoSInfo() {
        return this.requesterQoSInfo;
    }

}
