// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetAccessPointConfigForObjectProcessResponseBody extends TeaModel {
    @NameInMap("GetAccessPointConfigForObjectProcessResult")
    public GetAccessPointConfigForObjectProcessResult getAccessPointConfigForObjectProcessResult;

    public static GetAccessPointConfigForObjectProcessResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetAccessPointConfigForObjectProcessResponseBody self = new GetAccessPointConfigForObjectProcessResponseBody();
        return TeaModel.build(map, self);
    }

    public GetAccessPointConfigForObjectProcessResponseBody setGetAccessPointConfigForObjectProcessResult(GetAccessPointConfigForObjectProcessResult getAccessPointConfigForObjectProcessResult) {
        this.getAccessPointConfigForObjectProcessResult = getAccessPointConfigForObjectProcessResult;
        return this;
    }
    public GetAccessPointConfigForObjectProcessResult getGetAccessPointConfigForObjectProcessResult() {
        return this.getAccessPointConfigForObjectProcessResult;
    }

    public static class GetAccessPointConfigForObjectProcessResult extends TeaModel {
        @NameInMap("AllowAnonymousAccessForObjectProcess")
        public String allowAnonymousAccessForObjectProcess;

        @NameInMap("ObjectProcessConfiguration")
        public ObjectProcessConfiguration objectProcessConfiguration;

        @NameInMap("PublicAccessBlockConfiguration")
        public PublicAccessBlockConfiguration publicAccessBlockConfiguration;

        public static GetAccessPointConfigForObjectProcessResult build(java.util.Map<String, ?> map) throws Exception {
            GetAccessPointConfigForObjectProcessResult self = new GetAccessPointConfigForObjectProcessResult();
            return TeaModel.build(map, self);
        }

        public GetAccessPointConfigForObjectProcessResult setAllowAnonymousAccessForObjectProcess(String allowAnonymousAccessForObjectProcess) {
            this.allowAnonymousAccessForObjectProcess = allowAnonymousAccessForObjectProcess;
            return this;
        }
        public String getAllowAnonymousAccessForObjectProcess() {
            return this.allowAnonymousAccessForObjectProcess;
        }

        public GetAccessPointConfigForObjectProcessResult setObjectProcessConfiguration(ObjectProcessConfiguration objectProcessConfiguration) {
            this.objectProcessConfiguration = objectProcessConfiguration;
            return this;
        }
        public ObjectProcessConfiguration getObjectProcessConfiguration() {
            return this.objectProcessConfiguration;
        }

        public GetAccessPointConfigForObjectProcessResult setPublicAccessBlockConfiguration(PublicAccessBlockConfiguration publicAccessBlockConfiguration) {
            this.publicAccessBlockConfiguration = publicAccessBlockConfiguration;
            return this;
        }
        public PublicAccessBlockConfiguration getPublicAccessBlockConfiguration() {
            return this.publicAccessBlockConfiguration;
        }

    }

}
