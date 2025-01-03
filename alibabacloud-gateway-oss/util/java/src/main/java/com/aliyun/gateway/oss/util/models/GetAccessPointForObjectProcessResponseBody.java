// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetAccessPointForObjectProcessResponseBody extends TeaModel {
    /**
     * <p>The container that stores information about the Object FC Access Point.</p>
     */
    @NameInMap("GetAccessPointForObjectProcessResult")
    public GetAccessPointForObjectProcessResult getAccessPointForObjectProcessResult;

    public static GetAccessPointForObjectProcessResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetAccessPointForObjectProcessResponseBody self = new GetAccessPointForObjectProcessResponseBody();
        return TeaModel.build(map, self);
    }

    public GetAccessPointForObjectProcessResponseBody setGetAccessPointForObjectProcessResult(GetAccessPointForObjectProcessResult getAccessPointForObjectProcessResult) {
        this.getAccessPointForObjectProcessResult = getAccessPointForObjectProcessResult;
        return this;
    }
    public GetAccessPointForObjectProcessResult getGetAccessPointForObjectProcessResult() {
        return this.getAccessPointForObjectProcessResult;
    }

    public static class Endpoints extends TeaModel {
        /**
         * <p>The internal endpoint of the Object FC Access Point.</p>
         * 
         * <strong>example:</strong>
         * <p>fc-ap-01-111933544165****.oss-cn-qingdao-internal.oss-object-process.aliyuncs.com</p>
         */
        @NameInMap("InternalEndpoint")
        public String internalEndpoint;

        /**
         * <p>The public endpoint of the Object FC Access Point.</p>
         * 
         * <strong>example:</strong>
         * <p>fc-ap-01-111933544165****.oss-cn-qingdao.oss-object-process.aliyuncs.com</p>
         */
        @NameInMap("PublicEndpoint")
        public String publicEndpoint;

        public static Endpoints build(java.util.Map<String, ?> map) throws Exception {
            Endpoints self = new Endpoints();
            return TeaModel.build(map, self);
        }

        public Endpoints setInternalEndpoint(String internalEndpoint) {
            this.internalEndpoint = internalEndpoint;
            return this;
        }
        public String getInternalEndpoint() {
            return this.internalEndpoint;
        }

        public Endpoints setPublicEndpoint(String publicEndpoint) {
            this.publicEndpoint = publicEndpoint;
            return this;
        }
        public String getPublicEndpoint() {
            return this.publicEndpoint;
        }

    }

    public static class GetAccessPointForObjectProcessResult extends TeaModel {
        /**
         * <p>The alias of the Object FC Access Point.</p>
         * 
         * <strong>example:</strong>
         * <p>fc-ap-01-3b00521f653d2b3223680ec39dbbe2****-opapalias</p>
         */
        @NameInMap("AccessPointForObjectProcessAlias")
        public String accessPointForObjectProcessAlias;

        /**
         * <p>The ARN of the Object FC Access Point.</p>
         * 
         * <strong>example:</strong>
         * <p>acs:oss:cn-qingdao:119335441657143:accesspointforobjectprocess/fc-ap-01</p>
         */
        @NameInMap("AccessPointForObjectProcessArn")
        public String accessPointForObjectProcessArn;

        /**
         * <p>The name of the access point.</p>
         * 
         * <strong>example:</strong>
         * <p>ap-01</p>
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
         * <p>The UID of the Alibaba Cloud account to which the Object FC Access Point belongs.</p>
         * 
         * <strong>example:</strong>
         * <p>111933544165****</p>
         */
        @NameInMap("AccountId")
        public String accountId;

        /**
         * <p>Whether allow anonymous users to access this FC Access Point.</p>
         * 
         * <strong>example:</strong>
         * <p>false</p>
         */
        @NameInMap("AllowAnonymousAccessForObjectProcess")
        public String allowAnonymousAccessForObjectProcess;

        /**
         * <p>The time when the Object FC Access Point was created. The value is a timestamp.</p>
         * 
         * <strong>example:</strong>
         * <p>1626769503</p>
         */
        @NameInMap("CreationDate")
        public String creationDate;

        /**
         * <p>The container that stores the endpoints of the Object FC Access Point.</p>
         */
        @NameInMap("Endpoints")
        public Endpoints endpoints;

        /**
         * <p>The container in which the Block Public Access configurations are stored.</p>
         */
        @NameInMap("PublicAccessBlockConfiguration")
        public PublicAccessBlockConfiguration publicAccessBlockConfiguration;

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

        public static GetAccessPointForObjectProcessResult build(java.util.Map<String, ?> map) throws Exception {
            GetAccessPointForObjectProcessResult self = new GetAccessPointForObjectProcessResult();
            return TeaModel.build(map, self);
        }

        public GetAccessPointForObjectProcessResult setAccessPointForObjectProcessAlias(String accessPointForObjectProcessAlias) {
            this.accessPointForObjectProcessAlias = accessPointForObjectProcessAlias;
            return this;
        }
        public String getAccessPointForObjectProcessAlias() {
            return this.accessPointForObjectProcessAlias;
        }

        public GetAccessPointForObjectProcessResult setAccessPointForObjectProcessArn(String accessPointForObjectProcessArn) {
            this.accessPointForObjectProcessArn = accessPointForObjectProcessArn;
            return this;
        }
        public String getAccessPointForObjectProcessArn() {
            return this.accessPointForObjectProcessArn;
        }

        public GetAccessPointForObjectProcessResult setAccessPointName(String accessPointName) {
            this.accessPointName = accessPointName;
            return this;
        }
        public String getAccessPointName() {
            return this.accessPointName;
        }

        public GetAccessPointForObjectProcessResult setAccessPointNameForObjectProcess(String accessPointNameForObjectProcess) {
            this.accessPointNameForObjectProcess = accessPointNameForObjectProcess;
            return this;
        }
        public String getAccessPointNameForObjectProcess() {
            return this.accessPointNameForObjectProcess;
        }

        public GetAccessPointForObjectProcessResult setAccountId(String accountId) {
            this.accountId = accountId;
            return this;
        }
        public String getAccountId() {
            return this.accountId;
        }

        public GetAccessPointForObjectProcessResult setAllowAnonymousAccessForObjectProcess(String allowAnonymousAccessForObjectProcess) {
            this.allowAnonymousAccessForObjectProcess = allowAnonymousAccessForObjectProcess;
            return this;
        }
        public String getAllowAnonymousAccessForObjectProcess() {
            return this.allowAnonymousAccessForObjectProcess;
        }

        public GetAccessPointForObjectProcessResult setCreationDate(String creationDate) {
            this.creationDate = creationDate;
            return this;
        }
        public String getCreationDate() {
            return this.creationDate;
        }

        public GetAccessPointForObjectProcessResult setEndpoints(Endpoints endpoints) {
            this.endpoints = endpoints;
            return this;
        }
        public Endpoints getEndpoints() {
            return this.endpoints;
        }

        public GetAccessPointForObjectProcessResult setPublicAccessBlockConfiguration(PublicAccessBlockConfiguration publicAccessBlockConfiguration) {
            this.publicAccessBlockConfiguration = publicAccessBlockConfiguration;
            return this;
        }
        public PublicAccessBlockConfiguration getPublicAccessBlockConfiguration() {
            return this.publicAccessBlockConfiguration;
        }

        public GetAccessPointForObjectProcessResult setStatus(String status) {
            this.status = status;
            return this;
        }
        public String getStatus() {
            return this.status;
        }

    }

}
