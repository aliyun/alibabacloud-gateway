// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListAccessPointsForObjectProcessResponseBody extends TeaModel {
    /**
     * <p>The container that stores information about the Object FC Access Points that are returned.</p>
     */
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
        /**
         * <p>The alias of the Object FC Access Point.</p>
         * 
         * <strong>example:</strong>
         * <p>fc-ap-01-3b00521f653d2b3223680ec39dbbe2****-opapalias</p>
         */
        @NameInMap("AccessPointForObjectProcessAlias")
        public String accessPointForObjectProcessAlias;

        /**
         * <p>The name of the access point.</p>
         * 
         * <strong>example:</strong>
         * <p>fc-01</p>
         */
        @NameInMap("AccessPointName")
        public String accessPointName;

        /**
         * <p>The name of the Object FC Access Point.</p>
         * 
         * <strong>example:</strong>
         * <p>fc-ap-01</p>
         */
        @NameInMap("AccessPointNameForObjectProcess")
        public String accessPointNameForObjectProcess;

        /**
         * <p>Whether allow anonymous user access this FC Access Point.</p>
         * 
         * <strong>example:</strong>
         * <p>false</p>
         */
        @NameInMap("AllowAnonymousAccessForObjectProcess")
        public String allowAnonymousAccessForObjectProcess;

        /**
         * <p>The status of the Object FC Access Point. Valid values:</p>
         * <p>enable: The Object FC Access Point is created.</p>
         * <p>disable: The Object FC Access Point is disabled.</p>
         * <p>creating: The Object FC Access Point is being created.</p>
         * <p>deleting: The Object FC Access Point is deleted.</p>
         * 
         * <strong>example:</strong>
         * <p>enable</p>
         */
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

        public AccessPointForObjectProcess setAllowAnonymousAccessForObjectProcess(String allowAnonymousAccessForObjectProcess) {
            this.allowAnonymousAccessForObjectProcess = allowAnonymousAccessForObjectProcess;
            return this;
        }
        public String getAllowAnonymousAccessForObjectProcess() {
            return this.allowAnonymousAccessForObjectProcess;
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
        /**
         * <p>The container that stores information about a single Object FC Access Point.</p>
         */
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
        /**
         * <p>The container that stores information about all Object FC Access Points.</p>
         */
        @NameInMap("AccessPointsForObjectProcess")
        public AccessPointsForObjectProcess accessPointsForObjectProcess;

        /**
         * <p>The UID of the Alibaba Cloud account to which the Object FC Access Points belong.</p>
         * 
         * <strong>example:</strong>
         * <p>111933544165****</p>
         */
        @NameInMap("AccountId")
        public String accountId;

        /**
         * <p>Indicates whether the returned results are truncated. Valid values:</p>
         * <p>true: indicates that not all results are returned for the request.</p>
         * <p>false: indicates that all results are returned for the request.</p>
         * 
         * <strong>example:</strong>
         * <p>true</p>
         */
        @NameInMap("IsTruncated")
        public Boolean isTruncated;

        /**
         * <p>Indicates that this ListAccessPointsForObjectProcess request contains subsequent results. You need to set the NextContinuationToken element to continuation-token for subsequent results.</p>
         * 
         * <strong>example:</strong>
         * <p>abc</p>
         */
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
