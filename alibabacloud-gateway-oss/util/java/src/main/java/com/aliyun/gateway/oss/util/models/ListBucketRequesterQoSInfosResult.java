// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListBucketRequesterQoSInfosResult extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>test-bucket</p>
     */
    @NameInMap("Bucket")
    public String bucket;

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

    public static ListBucketRequesterQoSInfosResult build(java.util.Map<String, ?> map) throws Exception {
        ListBucketRequesterQoSInfosResult self = new ListBucketRequesterQoSInfosResult();
        return TeaModel.build(map, self);
    }

    public ListBucketRequesterQoSInfosResult setBucket(String bucket) {
        this.bucket = bucket;
        return this;
    }
    public String getBucket() {
        return this.bucket;
    }

    public ListBucketRequesterQoSInfosResult setContinuationToken(String continuationToken) {
        this.continuationToken = continuationToken;
        return this;
    }
    public String getContinuationToken() {
        return this.continuationToken;
    }

    public ListBucketRequesterQoSInfosResult setIsTruncated(Boolean isTruncated) {
        this.isTruncated = isTruncated;
        return this;
    }
    public Boolean getIsTruncated() {
        return this.isTruncated;
    }

    public ListBucketRequesterQoSInfosResult setNextContinuationToken(String nextContinuationToken) {
        this.nextContinuationToken = nextContinuationToken;
        return this;
    }
    public String getNextContinuationToken() {
        return this.nextContinuationToken;
    }

    public ListBucketRequesterQoSInfosResult setRequesterQoSInfo(java.util.List<RequesterQoSInfo> requesterQoSInfo) {
        this.requesterQoSInfo = requesterQoSInfo;
        return this;
    }
    public java.util.List<RequesterQoSInfo> getRequesterQoSInfo() {
        return this.requesterQoSInfo;
    }

}
