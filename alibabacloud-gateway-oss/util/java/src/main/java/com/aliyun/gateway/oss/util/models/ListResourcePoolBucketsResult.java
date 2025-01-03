// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListResourcePoolBucketsResult extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>abcd</p>
     */
    @NameInMap("ContionuationToken")
    public String contionuationToken;

    /**
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("IsTruncated")
    public Boolean isTruncated;

    /**
     * <strong>example:</strong>
     * <p>defg</p>
     */
    @NameInMap("NextContionuationToken")
    public String nextContionuationToken;

    /**
     * <strong>example:</strong>
     * <p>rp-for-api</p>
     */
    @NameInMap("ResourcePool")
    public String resourcePool;

    @NameInMap("ResourcePoolBucket")
    public java.util.List<ResourcePoolBucket> resourcePoolBucket;

    public static ListResourcePoolBucketsResult build(java.util.Map<String, ?> map) throws Exception {
        ListResourcePoolBucketsResult self = new ListResourcePoolBucketsResult();
        return TeaModel.build(map, self);
    }

    public ListResourcePoolBucketsResult setContionuationToken(String contionuationToken) {
        this.contionuationToken = contionuationToken;
        return this;
    }
    public String getContionuationToken() {
        return this.contionuationToken;
    }

    public ListResourcePoolBucketsResult setIsTruncated(Boolean isTruncated) {
        this.isTruncated = isTruncated;
        return this;
    }
    public Boolean getIsTruncated() {
        return this.isTruncated;
    }

    public ListResourcePoolBucketsResult setNextContionuationToken(String nextContionuationToken) {
        this.nextContionuationToken = nextContionuationToken;
        return this;
    }
    public String getNextContionuationToken() {
        return this.nextContionuationToken;
    }

    public ListResourcePoolBucketsResult setResourcePool(String resourcePool) {
        this.resourcePool = resourcePool;
        return this;
    }
    public String getResourcePool() {
        return this.resourcePool;
    }

    public ListResourcePoolBucketsResult setResourcePoolBucket(java.util.List<ResourcePoolBucket> resourcePoolBucket) {
        this.resourcePoolBucket = resourcePoolBucket;
        return this;
    }
    public java.util.List<ResourcePoolBucket> getResourcePoolBucket() {
        return this.resourcePoolBucket;
    }

}
