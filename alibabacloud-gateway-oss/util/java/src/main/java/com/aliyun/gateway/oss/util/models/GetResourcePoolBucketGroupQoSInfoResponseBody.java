// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetResourcePoolBucketGroupQoSInfoResponseBody extends TeaModel {
    @NameInMap("ResourcePoolBucketGroupQoSInfo")
    public ResourcePoolBucketGroupQoSInfo resourcePoolBucketGroupQoSInfo;

    public static GetResourcePoolBucketGroupQoSInfoResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetResourcePoolBucketGroupQoSInfoResponseBody self = new GetResourcePoolBucketGroupQoSInfoResponseBody();
        return TeaModel.build(map, self);
    }

    public GetResourcePoolBucketGroupQoSInfoResponseBody setResourcePoolBucketGroupQoSInfo(ResourcePoolBucketGroupQoSInfo resourcePoolBucketGroupQoSInfo) {
        this.resourcePoolBucketGroupQoSInfo = resourcePoolBucketGroupQoSInfo;
        return this;
    }
    public ResourcePoolBucketGroupQoSInfo getResourcePoolBucketGroupQoSInfo() {
        return this.resourcePoolBucketGroupQoSInfo;
    }

}
