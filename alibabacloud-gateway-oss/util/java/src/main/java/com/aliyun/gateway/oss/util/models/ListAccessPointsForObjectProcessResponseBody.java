// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class ListAccessPointsForObjectProcessResponseBody extends TeaModel {
    @NameInMap("ListAccessPointsForObjectProcessResult")
    public ListAccessPointsForObjectProcessResult listAccessPointsForObjectProcessResult;

    public static ListAccessPointsForObjectProcessResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListAccessPointsForObjectProcessResponseBody self = new ListAccessPointsForObjectProcessResponseBody();
        return TeaModel.build(map, self);
    }

    public ListAccessPointsForObjectProcessResponseBody setListAccessPointsForObjectProcessResult(ListAccessPointsForObjectProcessResult listAccessPointsForObjectProcessResult) {
        this.listAccessPointsForObjectProcessResult = listAccessPointsForObjectProcessResult;
        return this;
    }
    public ListAccessPointsForObjectProcessResult getListAccessPointsForObjectProcessResult() {
        return this.listAccessPointsForObjectProcessResult;
    }

    public static class AccessPointForObjectProcess extends TeaModel {
        @NameInMap("AccessPointForObjectProcessAlias")
        public String accessPointForObjectProcessAlias;

        @NameInMap("AccessPointName")
        public String accessPointName;

        @NameInMap("AccessPointNameForObjectProcess")
        public String accessPointNameForObjectProcess;

        @NameInMap("Status")
        public String status;

        public static AccessPointForObjectProcess build(java.util.Map<String, ?> map) throws Exception {
            AccessPointForObjectProcess self = new AccessPointForObjectProcess();
            return TeaModel.build(map, self);
        }

        public AccessPointForObjectProcess setAccessPointForObjectProcessAlias(String accessPointForObjectProcessAlias) {
            this.accessPointForObjectProcessAlias = accessPointForObjectProcessAlias;
            return this;
        }
        public String getAccessPointForObjectProcessAlias() {
            return this.accessPointForObjectProcessAlias;
        }

        public AccessPointForObjectProcess setAccessPointName(String accessPointName) {
            this.accessPointName = accessPointName;
            return this;
        }
        public String getAccessPointName() {
            return this.accessPointName;
        }

        public AccessPointForObjectProcess setAccessPointNameForObjectProcess(String accessPointNameForObjectProcess) {
            this.accessPointNameForObjectProcess = accessPointNameForObjectProcess;
            return this;
        }
        public String getAccessPointNameForObjectProcess() {
            return this.accessPointNameForObjectProcess;
        }

        public AccessPointForObjectProcess setStatus(String status) {
            this.status = status;
            return this;
        }
        public String getStatus() {
            return this.status;
        }

    }

    public static class AccessPointsForObjectProcess extends TeaModel {
        @NameInMap("AccessPointForObjectProcess")
        public java.util.List<AccessPointForObjectProcess> accessPointForObjectProcess;

        public static AccessPointsForObjectProcess build(java.util.Map<String, ?> map) throws Exception {
            AccessPointsForObjectProcess self = new AccessPointsForObjectProcess();
            return TeaModel.build(map, self);
        }

        public AccessPointsForObjectProcess setAccessPointForObjectProcess(java.util.List<AccessPointForObjectProcess> accessPointForObjectProcess) {
            this.accessPointForObjectProcess = accessPointForObjectProcess;
            return this;
        }
        public java.util.List<AccessPointForObjectProcess> getAccessPointForObjectProcess() {
            return this.accessPointForObjectProcess;
        }

    }

    public static class ListAccessPointsForObjectProcessResult extends TeaModel {
        @NameInMap("AccessPointsForObjectProcess")
        public AccessPointsForObjectProcess accessPointsForObjectProcess;

        @NameInMap("AccountId")
        public String accountId;

        @NameInMap("IsTruncated")
        public Boolean isTruncated;

        @NameInMap("NextContinuationToken")
        public String nextContinuationToken;

        public static ListAccessPointsForObjectProcessResult build(java.util.Map<String, ?> map) throws Exception {
            ListAccessPointsForObjectProcessResult self = new ListAccessPointsForObjectProcessResult();
            return TeaModel.build(map, self);
        }

        public ListAccessPointsForObjectProcessResult setAccessPointsForObjectProcess(AccessPointsForObjectProcess accessPointsForObjectProcess) {
            this.accessPointsForObjectProcess = accessPointsForObjectProcess;
            return this;
        }
        public AccessPointsForObjectProcess getAccessPointsForObjectProcess() {
            return this.accessPointsForObjectProcess;
        }

        public ListAccessPointsForObjectProcessResult setAccountId(String accountId) {
            this.accountId = accountId;
            return this;
        }
        public String getAccountId() {
            return this.accountId;
        }

        public ListAccessPointsForObjectProcessResult setIsTruncated(Boolean isTruncated) {
            this.isTruncated = isTruncated;
            return this;
        }
        public Boolean getIsTruncated() {
            return this.isTruncated;
        }

        public ListAccessPointsForObjectProcessResult setNextContinuationToken(String nextContinuationToken) {
            this.nextContinuationToken = nextContinuationToken;
            return this;
        }
        public String getNextContinuationToken() {
            return this.nextContinuationToken;
        }

    }

}
