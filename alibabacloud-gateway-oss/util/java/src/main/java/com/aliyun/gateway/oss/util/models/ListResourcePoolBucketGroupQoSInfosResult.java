// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListResourcePoolBucketGroupQoSInfosResult extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>abc-</p>
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
     * <p>rp-for-ai</p>
     */
    @NameInMap("ResourcePool")
    public String resourcePool;

    @NameInMap("ResourcePoolBucketGroupQoSInfo")
    public java.util.List<ResourcePoolBucketGroupQoSInfo> resourcePoolBucketGroupQoSInfo;

    public static ListResourcePoolBucketGroupQoSInfosResult build(java.util.Map<String, ?> map) throws Exception {
        ListResourcePoolBucketGroupQoSInfosResult self = new ListResourcePoolBucketGroupQoSInfosResult();
        return TeaModel.build(map, self);
    }

    public ListResourcePoolBucketGroupQoSInfosResult setContinuationToken(String continuationToken) {
        this.continuationToken = continuationToken;
        return this;
    }
    public String getContinuationToken() {
        return this.continuationToken;
    }

    public ListResourcePoolBucketGroupQoSInfosResult setIsTruncated(Boolean isTruncated) {
        this.isTruncated = isTruncated;
        return this;
    }
    public Boolean getIsTruncated() {
        return this.isTruncated;
    }

    public ListResourcePoolBucketGroupQoSInfosResult setNextContinuationToken(String nextContinuationToken) {
        this.nextContinuationToken = nextContinuationToken;
        return this;
    }
    public String getNextContinuationToken() {
        return this.nextContinuationToken;
    }

    public ListResourcePoolBucketGroupQoSInfosResult setResourcePool(String resourcePool) {
        this.resourcePool = resourcePool;
        return this;
    }
    public String getResourcePool() {
        return this.resourcePool;
    }

    public ListResourcePoolBucketGroupQoSInfosResult setResourcePoolBucketGroupQoSInfo(java.util.List<ResourcePoolBucketGroupQoSInfo> resourcePoolBucketGroupQoSInfo) {
        this.resourcePoolBucketGroupQoSInfo = resourcePoolBucketGroupQoSInfo;
        return this;
    }
    public java.util.List<ResourcePoolBucketGroupQoSInfo> getResourcePoolBucketGroupQoSInfo() {
        return this.resourcePoolBucketGroupQoSInfo;
    }

}
