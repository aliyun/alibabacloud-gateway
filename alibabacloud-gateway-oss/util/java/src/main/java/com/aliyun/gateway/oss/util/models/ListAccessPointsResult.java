// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListAccessPointsResult extends TeaModel {
    @NameInMap("AccessPoints")
    public AccessPoints accessPoints;

    @NameInMap("AccountId")
    public String accountId;

    @NameInMap("IsTruncated")
    public String isTruncated;

    @NameInMap("MaxKeys")
    public Integer maxKeys;

    @NameInMap("NextContinuationToken")
    public String nextContinuationToken;

    public static ListAccessPointsResult build(java.util.Map<String, ?> map) throws Exception {
        ListAccessPointsResult self = new ListAccessPointsResult();
        return TeaModel.build(map, self);
    }

    public ListAccessPointsResult setAccessPoints(AccessPoints accessPoints) {
        this.accessPoints = accessPoints;
        return this;
    }
    public AccessPoints getAccessPoints() {
        return this.accessPoints;
    }

    public ListAccessPointsResult setAccountId(String accountId) {
        this.accountId = accountId;
        return this;
    }
    public String getAccountId() {
        return this.accountId;
    }

    public ListAccessPointsResult setIsTruncated(String isTruncated) {
        this.isTruncated = isTruncated;
        return this;
    }
    public String getIsTruncated() {
        return this.isTruncated;
    }

    public ListAccessPointsResult setMaxKeys(Integer maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Integer getMaxKeys() {
        return this.maxKeys;
    }

    public ListAccessPointsResult setNextContinuationToken(String nextContinuationToken) {
        this.nextContinuationToken = nextContinuationToken;
        return this;
    }
    public String getNextContinuationToken() {
        return this.nextContinuationToken;
    }

    public static class AccessPoints extends TeaModel {
        @NameInMap("AccessPoint")
        public java.util.List<AccessPoint> accessPoint;

        public static AccessPoints build(java.util.Map<String, ?> map) throws Exception {
            AccessPoints self = new AccessPoints();
            return TeaModel.build(map, self);
        }

        public AccessPoints setAccessPoint(java.util.List<AccessPoint> accessPoint) {
            this.accessPoint = accessPoint;
            return this;
        }
        public java.util.List<AccessPoint> getAccessPoint() {
            return this.accessPoint;
        }

    }

}
