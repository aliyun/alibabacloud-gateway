// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListResourcePoolRequesterQoSInfosResult extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>29387293298</p>
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
     * <p>29387293298</p>
     */
    @NameInMap("NextContinuationToken")
    public String nextContinuationToken;

    @NameInMap("RequesterQoSInfo")
    public java.util.List<RequesterQoSInfo> requesterQoSInfo;

    /**
     * <strong>example:</strong>
     * <p>rp-for-ai</p>
     */
    @NameInMap("ResourcePool")
    public String resourcePool;

    public static ListResourcePoolRequesterQoSInfosResult build(java.util.Map<String, ?> map) throws Exception {
        ListResourcePoolRequesterQoSInfosResult self = new ListResourcePoolRequesterQoSInfosResult();
        return TeaModel.build(map, self);
    }

    public ListResourcePoolRequesterQoSInfosResult setContinuationToken(String continuationToken) {
        this.continuationToken = continuationToken;
        return this;
    }
    public String getContinuationToken() {
        return this.continuationToken;
    }

    public ListResourcePoolRequesterQoSInfosResult setIsTruncated(Boolean isTruncated) {
        this.isTruncated = isTruncated;
        return this;
    }
    public Boolean getIsTruncated() {
        return this.isTruncated;
    }

    public ListResourcePoolRequesterQoSInfosResult setNextContinuationToken(String nextContinuationToken) {
        this.nextContinuationToken = nextContinuationToken;
        return this;
    }
    public String getNextContinuationToken() {
        return this.nextContinuationToken;
    }

    public ListResourcePoolRequesterQoSInfosResult setRequesterQoSInfo(java.util.List<RequesterQoSInfo> requesterQoSInfo) {
        this.requesterQoSInfo = requesterQoSInfo;
        return this;
    }
    public java.util.List<RequesterQoSInfo> getRequesterQoSInfo() {
        return this.requesterQoSInfo;
    }

    public ListResourcePoolRequesterQoSInfosResult setResourcePool(String resourcePool) {
        this.resourcePool = resourcePool;
        return this;
    }
    public String getResourcePool() {
        return this.resourcePool;
    }

}
