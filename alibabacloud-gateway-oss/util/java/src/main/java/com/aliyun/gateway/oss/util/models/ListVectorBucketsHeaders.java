// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListVectorBucketsHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    @NameInMap("x-oss-resource-group-id")
    public String xOssResourceGroupId;

    public static ListVectorBucketsHeaders build(java.util.Map<String, ?> map) throws Exception {
        ListVectorBucketsHeaders self = new ListVectorBucketsHeaders();
        return TeaModel.build(map, self);
    }

    public ListVectorBucketsHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public ListVectorBucketsHeaders setXOssResourceGroupId(String xOssResourceGroupId) {
        this.xOssResourceGroupId = xOssResourceGroupId;
        return this;
    }
    public String getXOssResourceGroupId() {
        return this.xOssResourceGroupId;
    }

}
