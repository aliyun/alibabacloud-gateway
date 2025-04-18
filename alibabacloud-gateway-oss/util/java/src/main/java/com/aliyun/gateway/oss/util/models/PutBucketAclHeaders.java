// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketAclHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The ACL that you want to configure or modify for the bucket. The x-oss-acl header is included in PutBucketAcl requests to configure or modify the ACL of the bucket. If this header is not included, the ACL configurations do not take effect.\
     * Valid values:</p>
     * <ul>
     * <li>public-read-write: All users can read and write objects in the bucket. Exercise caution when you set the value to public-read-write.</li>
     * <li>public-read: Only the owner and authorized users of the bucket can read and write objects in the bucket. Other users can only read objects in the bucket. Exercise caution when you set the value to public-read.</li>
     * <li>private: Only the owner and authorized users of this bucket can read and write objects in the bucket. Other users cannot access objects in the bucket.</li>
     * </ul>
     * <p>This parameter is required.</p>
     */
    @NameInMap("x-oss-acl")
    public String acl;

    public static PutBucketAclHeaders build(java.util.Map<String, ?> map) throws Exception {
        PutBucketAclHeaders self = new PutBucketAclHeaders();
        return TeaModel.build(map, self);
    }

    public PutBucketAclHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public PutBucketAclHeaders setAcl(String acl) {
        this.acl = acl;
        return this;
    }
    public String getAcl() {
        return this.acl;
    }

}
