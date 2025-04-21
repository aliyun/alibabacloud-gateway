// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListResourcePoolBucketGroupsResult extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>test-</p>
     */
    @NameInMap("ContinuationToken")
    public String continuationToken;

    /**
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("IsTruncated")
    public Boolean isTruncated;

    /**
     * <strong>example:</strong>
     * <p>xyz-</p>
     */
    @NameInMap("NextContinuationToken")
    public String nextContinuationToken;

    /**
     * <strong>example:</strong>
     * <p>test-resource-pool</p>
     */
    @NameInMap("ResourcePool")
    public String resourcePool;

    @NameInMap("ResourcePoolBucketGroup")
    public java.util.List<ResourcePoolBucketGroup> resourcePoolBucketGroup;

    public static ListResourcePoolBucketGroupsResult build(java.util.Map<String, ?> map) throws Exception {
        ListResourcePoolBucketGroupsResult self = new ListResourcePoolBucketGroupsResult();
        return TeaModel.build(map, self);
    }

    public ListResourcePoolBucketGroupsResult setContinuationToken(String continuationToken) {
        this.continuationToken = continuationToken;
        return this;
    }
    public String getContinuationToken() {
        return this.continuationToken;
    }

    public ListResourcePoolBucketGroupsResult setIsTruncated(Boolean isTruncated) {
        this.isTruncated = isTruncated;
        return this;
    }
    public Boolean getIsTruncated() {
        return this.isTruncated;
    }

    public ListResourcePoolBucketGroupsResult setNextContinuationToken(String nextContinuationToken) {
        this.nextContinuationToken = nextContinuationToken;
        return this;
    }
    public String getNextContinuationToken() {
        return this.nextContinuationToken;
    }

    public ListResourcePoolBucketGroupsResult setResourcePool(String resourcePool) {
        this.resourcePool = resourcePool;
        return this;
    }
    public String getResourcePool() {
        return this.resourcePool;
    }

    public ListResourcePoolBucketGroupsResult setResourcePoolBucketGroup(java.util.List<ResourcePoolBucketGroup> resourcePoolBucketGroup) {
        this.resourcePoolBucketGroup = resourcePoolBucketGroup;
        return this;
    }
    public java.util.List<ResourcePoolBucketGroup> getResourcePoolBucketGroup() {
        return this.resourcePoolBucketGroup;
    }

}
