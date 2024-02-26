// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class PutBucketHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The access control list (ACL) of the bucket to be created. Valid values:</p>
     * <br>
     * <p>*   public-read-write</p>
     * <p>*   public-read</p>
     * <p>*   private (default)</p>
     * <br>
     * <p>For more information, see [Bucket ACL](~~31843~~).</p>
     */
    @NameInMap("x-oss-acl")
    public String acl;

    /**
     * <p>The ID of the resource group.</p>
     * <br>
     * <p>*   If you include the header in the request and specify the ID of the resource group, the bucket that you create belongs to the resource group. If the specified resource group ID is rg-default-id, the bucket that you create belongs to the default resource group.</p>
     * <p>*   If you do not include the header in the request, the bucket that you create belongs to the default resource group.</p>
     * <br>
     * <p>You can obtain the ID of a resource group in the Resource Management console or by calling the ListResourceGroups operation. For more information, see [View basic information of a resource group](~~151181~~) and [ListResourceGroups](~~158855~~).</p>
     * <br>
     * <p>>  You cannot configure a resource group for an Anywhere Bucket.</p>
     */
    @NameInMap("x-oss-resource-group-id")
    public String xOssResourceGroupId;

    public static PutBucketHeaders build(java.util.Map<String, ?> map) throws Exception {
        PutBucketHeaders self = new PutBucketHeaders();
        return TeaModel.build(map, self);
    }

    public PutBucketHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public PutBucketHeaders setAcl(String acl) {
        this.acl = acl;
        return this;
    }
    public String getAcl() {
        return this.acl;
    }

    public PutBucketHeaders setXOssResourceGroupId(String xOssResourceGroupId) {
        this.xOssResourceGroupId = xOssResourceGroupId;
        return this;
    }
    public String getXOssResourceGroupId() {
        return this.xOssResourceGroupId;
    }

}
