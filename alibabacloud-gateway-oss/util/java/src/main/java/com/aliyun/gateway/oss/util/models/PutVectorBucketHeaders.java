// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutVectorBucketHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    @NameInMap("x-oss-resource-group-id")
    public String xOssResourceGroupId;

    public static PutVectorBucketHeaders build(java.util.Map<String, ?> map) throws Exception {
        PutVectorBucketHeaders self = new PutVectorBucketHeaders();
        return TeaModel.build(map, self);
    }

    public PutVectorBucketHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public PutVectorBucketHeaders setXOssResourceGroupId(String xOssResourceGroupId) {
        this.xOssResourceGroupId = xOssResourceGroupId;
        return this;
    }
    public String getXOssResourceGroupId() {
        return this.xOssResourceGroupId;
    }

}
