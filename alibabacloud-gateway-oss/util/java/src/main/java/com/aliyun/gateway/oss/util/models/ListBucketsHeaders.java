// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListBucketsHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The ID of the resource group to which the bucket belongs.</p>
     */
    @NameInMap("x-oss-resource-group-id")
    public String xOssResourceGroupId;

    public static ListBucketsHeaders build(java.util.Map<String, ?> map) throws Exception {
        ListBucketsHeaders self = new ListBucketsHeaders();
        return TeaModel.build(map, self);
    }

    public ListBucketsHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public ListBucketsHeaders setXOssResourceGroupId(String xOssResourceGroupId) {
        this.xOssResourceGroupId = xOssResourceGroupId;
        return this;
    }
    public String getXOssResourceGroupId() {
        return this.xOssResourceGroupId;
    }

}
