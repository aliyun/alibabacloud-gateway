// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The access control list (ACL) of the bucket to be created. Valid values:</p>
     * <ul>
     * <li>public-read-write</li>
     * <li>public-read</li>
     * <li>private (default)</li>
     * </ul>
     * <p>For more information, see <a href="https://help.aliyun.com/document_detail/31843.html">Bucket ACL</a>.</p>
     */
    @NameInMap("x-oss-acl")
    public String acl;

    @NameInMap("x-oss-bucket-tagging")
    public String xOssBucketTagging;

    /**
     * <p>The ID of the resource group.</p>
     * <ul>
     * <li>If you include the header in the request and specify the ID of the resource group, the bucket that you create belongs to the resource group. If the specified resource group ID is rg-default-id, the bucket that you create belongs to the default resource group.</li>
     * <li>If you do not include the header in the request, the bucket that you create belongs to the default resource group.</li>
     * </ul>
     * <p>You can obtain the ID of a resource group in the Resource Management console or by calling the ListResourceGroups operation. For more information, see <a href="https://help.aliyun.com/document_detail/151181.html">View basic information of a resource group</a> and <a href="https://help.aliyun.com/document_detail/158855.html">ListResourceGroups</a>.</p>
     * <blockquote>
     * <p> You cannot configure a resource group for an Anywhere Bucket.</p>
     * </blockquote>
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

    public PutBucketHeaders setXOssBucketTagging(String xOssBucketTagging) {
        this.xOssBucketTagging = xOssBucketTagging;
        return this;
    }
    public String getXOssBucketTagging() {
        return this.xOssBucketTagging;
    }

    public PutBucketHeaders setXOssResourceGroupId(String xOssResourceGroupId) {
        this.xOssResourceGroupId = xOssResourceGroupId;
        return this;
    }
    public String getXOssResourceGroupId() {
        return this.xOssResourceGroupId;
    }

}
