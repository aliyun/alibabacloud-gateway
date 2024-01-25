// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ListAccessPointsForObjectProcessResponseBody extends TeaModel {
    @NameInMap("AccessPointsForObjectProcess")
    public ListAccessPointsForObjectProcessResponseBodyAccessPointsForObjectProcess accessPointsForObjectProcess;

    @NameInMap("AccountId")
    public String accountId;

    @NameInMap("IsTruncated")
    public Boolean isTruncated;

    @NameInMap("NextContinuationToken")
    public String nextContinuationToken;

    public static ListAccessPointsForObjectProcessResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListAccessPointsForObjectProcessResponseBody self = new ListAccessPointsForObjectProcessResponseBody();
        return TeaModel.build(map, self);
    }

    public ListAccessPointsForObjectProcessResponseBody setAccessPointsForObjectProcess(ListAccessPointsForObjectProcessResponseBodyAccessPointsForObjectProcess accessPointsForObjectProcess) {
        this.accessPointsForObjectProcess = accessPointsForObjectProcess;
        return this;
    }
    public ListAccessPointsForObjectProcessResponseBodyAccessPointsForObjectProcess getAccessPointsForObjectProcess() {
        return this.accessPointsForObjectProcess;
    }

    public ListAccessPointsForObjectProcessResponseBody setAccountId(String accountId) {
        this.accountId = accountId;
        return this;
    }
    public String getAccountId() {
        return this.accountId;
    }

    public ListAccessPointsForObjectProcessResponseBody setIsTruncated(Boolean isTruncated) {
        this.isTruncated = isTruncated;
        return this;
    }
    public Boolean getIsTruncated() {
        return this.isTruncated;
    }

    public ListAccessPointsForObjectProcessResponseBody setNextContinuationToken(String nextContinuationToken) {
        this.nextContinuationToken = nextContinuationToken;
        return this;
    }
    public String getNextContinuationToken() {
        return this.nextContinuationToken;
    }

    public static class ListAccessPointsForObjectProcessResponseBodyAccessPointsForObjectProcessAccessPointForObjectProcess extends TeaModel {
        @NameInMap("AccessPointForObjectProcessAlias")
        public String accessPointForObjectProcessAlias;

        @NameInMap("AccessPointName")
        public String accessPointName;

        @NameInMap("AccessPointNameForObjectProcess")
        public String accessPointNameForObjectProcess;

        @NameInMap("Status")
        public String status;

        public static ListAccessPointsForObjectProcessResponseBodyAccessPointsForObjectProcessAccessPointForObjectProcess build(java.util.Map<String, ?> map) throws Exception {
            ListAccessPointsForObjectProcessResponseBodyAccessPointsForObjectProcessAccessPointForObjectProcess self = new ListAccessPointsForObjectProcessResponseBodyAccessPointsForObjectProcessAccessPointForObjectProcess();
            return TeaModel.build(map, self);
        }

        public ListAccessPointsForObjectProcessResponseBodyAccessPointsForObjectProcessAccessPointForObjectProcess setAccessPointForObjectProcessAlias(String accessPointForObjectProcessAlias) {
            this.accessPointForObjectProcessAlias = accessPointForObjectProcessAlias;
            return this;
        }
        public String getAccessPointForObjectProcessAlias() {
            return this.accessPointForObjectProcessAlias;
        }

        public ListAccessPointsForObjectProcessResponseBodyAccessPointsForObjectProcessAccessPointForObjectProcess setAccessPointName(String accessPointName) {
            this.accessPointName = accessPointName;
            return this;
        }
        public String getAccessPointName() {
            return this.accessPointName;
        }

        public ListAccessPointsForObjectProcessResponseBodyAccessPointsForObjectProcessAccessPointForObjectProcess setAccessPointNameForObjectProcess(String accessPointNameForObjectProcess) {
            this.accessPointNameForObjectProcess = accessPointNameForObjectProcess;
            return this;
        }
        public String getAccessPointNameForObjectProcess() {
            return this.accessPointNameForObjectProcess;
        }

        public ListAccessPointsForObjectProcessResponseBodyAccessPointsForObjectProcessAccessPointForObjectProcess setStatus(String status) {
            this.status = status;
            return this;
        }
        public String getStatus() {
            return this.status;
        }

    }

    public static class ListAccessPointsForObjectProcessResponseBodyAccessPointsForObjectProcess extends TeaModel {
        @NameInMap("AccessPointForObjectProcess")
        public java.util.List<ListAccessPointsForObjectProcessResponseBodyAccessPointsForObjectProcessAccessPointForObjectProcess> accessPointForObjectProcess;

        public static ListAccessPointsForObjectProcessResponseBodyAccessPointsForObjectProcess build(java.util.Map<String, ?> map) throws Exception {
            ListAccessPointsForObjectProcessResponseBodyAccessPointsForObjectProcess self = new ListAccessPointsForObjectProcessResponseBodyAccessPointsForObjectProcess();
            return TeaModel.build(map, self);
        }

        public ListAccessPointsForObjectProcessResponseBodyAccessPointsForObjectProcess setAccessPointForObjectProcess(java.util.List<ListAccessPointsForObjectProcessResponseBodyAccessPointsForObjectProcessAccessPointForObjectProcess> accessPointForObjectProcess) {
            this.accessPointForObjectProcess = accessPointForObjectProcess;
            return this;
        }
        public java.util.List<ListAccessPointsForObjectProcessResponseBodyAccessPointsForObjectProcessAccessPointForObjectProcess> getAccessPointForObjectProcess() {
            return this.accessPointForObjectProcess;
        }

    }

}
