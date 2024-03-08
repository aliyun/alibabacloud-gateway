// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetAccessPointForObjectProcessResponseBody extends TeaModel {
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
    public GetAccessPointForObjectProcessResponseBodyEndpoints endpoints;

    @NameInMap("Status")
    public String status;

    public static GetAccessPointForObjectProcessResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetAccessPointForObjectProcessResponseBody self = new GetAccessPointForObjectProcessResponseBody();
        return TeaModel.build(map, self);
    }

    public GetAccessPointForObjectProcessResponseBody setAccessPointForObjectProcessAlias(String accessPointForObjectProcessAlias) {
        this.accessPointForObjectProcessAlias = accessPointForObjectProcessAlias;
        return this;
    }
    public String getAccessPointForObjectProcessAlias() {
        return this.accessPointForObjectProcessAlias;
    }

    public GetAccessPointForObjectProcessResponseBody setAccessPointForObjectProcessArn(String accessPointForObjectProcessArn) {
        this.accessPointForObjectProcessArn = accessPointForObjectProcessArn;
        return this;
    }
    public String getAccessPointForObjectProcessArn() {
        return this.accessPointForObjectProcessArn;
    }

    public GetAccessPointForObjectProcessResponseBody setAccessPointName(String accessPointName) {
        this.accessPointName = accessPointName;
        return this;
    }
    public String getAccessPointName() {
        return this.accessPointName;
    }

    public GetAccessPointForObjectProcessResponseBody setAccessPointNameForObjectProcess(String accessPointNameForObjectProcess) {
        this.accessPointNameForObjectProcess = accessPointNameForObjectProcess;
        return this;
    }
    public String getAccessPointNameForObjectProcess() {
        return this.accessPointNameForObjectProcess;
    }

    public GetAccessPointForObjectProcessResponseBody setAccountId(String accountId) {
        this.accountId = accountId;
        return this;
    }
    public String getAccountId() {
        return this.accountId;
    }

    public GetAccessPointForObjectProcessResponseBody setCreationDate(String creationDate) {
        this.creationDate = creationDate;
        return this;
    }
    public String getCreationDate() {
        return this.creationDate;
    }

    public GetAccessPointForObjectProcessResponseBody setEndpoints(GetAccessPointForObjectProcessResponseBodyEndpoints endpoints) {
        this.endpoints = endpoints;
        return this;
    }
    public GetAccessPointForObjectProcessResponseBodyEndpoints getEndpoints() {
        return this.endpoints;
    }

    public GetAccessPointForObjectProcessResponseBody setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

    public static class GetAccessPointForObjectProcessResponseBodyEndpoints extends TeaModel {
        @NameInMap("InternalEndpoint")
        public String internalEndpoint;

        @NameInMap("PublicEndpoint")
        public String publicEndpoint;

        public static GetAccessPointForObjectProcessResponseBodyEndpoints build(java.util.Map<String, ?> map) throws Exception {
            GetAccessPointForObjectProcessResponseBodyEndpoints self = new GetAccessPointForObjectProcessResponseBodyEndpoints();
            return TeaModel.build(map, self);
        }

        public GetAccessPointForObjectProcessResponseBodyEndpoints setInternalEndpoint(String internalEndpoint) {
            this.internalEndpoint = internalEndpoint;
            return this;
        }
        public String getInternalEndpoint() {
            return this.internalEndpoint;
        }

        public GetAccessPointForObjectProcessResponseBodyEndpoints setPublicEndpoint(String publicEndpoint) {
            this.publicEndpoint = publicEndpoint;
            return this;
        }
        public String getPublicEndpoint() {
            return this.publicEndpoint;
        }

    }

}
