// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class PutBucketLifecycleHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    @NameInMap("x-oss-allow-same-action-overlap")
    public String xOssAllowSameActionOverlap;

    public static PutBucketLifecycleHeaders build(java.util.Map<String, ?> map) throws Exception {
        PutBucketLifecycleHeaders self = new PutBucketLifecycleHeaders();
        return TeaModel.build(map, self);
    }

    public PutBucketLifecycleHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public PutBucketLifecycleHeaders setXOssAllowSameActionOverlap(String xOssAllowSameActionOverlap) {
        this.xOssAllowSameActionOverlap = xOssAllowSameActionOverlap;
        return this;
    }
    public String getXOssAllowSameActionOverlap() {
        return this.xOssAllowSameActionOverlap;
    }

}
