// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class GetAccessPointForObjectProcessResponseBody extends TeaModel {
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
        @NameInMap("InternalEndpoint")
        public String internalEndpoint;

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
        @NameInMap("AccessPointForObjectProcessAlias")
        public String accessPointForObjectProcessAlias;

        @NameInMap("AccessPointForObjectProcessArn")
        public String accessPointForObjectProcessArn;

        @NameInMap("AccessPointName")
        public String accessPointName;

        @NameInMap("AccessPointNameForObjectProcess")
        public String accessPointNameForObjectProcess;

        @NameInMap("AccountId")
        public String accountId;

        @NameInMap("CreationDate")
        public String creationDate;

        @NameInMap("Endpoints")
        public Endpoints endpoints;

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

        public GetAccessPointForObjectProcessResult setStatus(String status) {
            this.status = status;
            return this;
        }
        public String getStatus() {
            return this.status;
        }

    }

}
