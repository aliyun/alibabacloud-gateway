// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

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
        @NameInMap("ObjectProcessConfiguration")
        public ObjectProcessConfiguration objectProcessConfiguration;

        public static GetAccessPointConfigForObjectProcessResult build(java.util.Map<String, ?> map) throws Exception {
            GetAccessPointConfigForObjectProcessResult self = new GetAccessPointConfigForObjectProcessResult();
            return TeaModel.build(map, self);
        }

        public GetAccessPointConfigForObjectProcessResult setObjectProcessConfiguration(ObjectProcessConfiguration objectProcessConfiguration) {
            this.objectProcessConfiguration = objectProcessConfiguration;
            return this;
        }
        public ObjectProcessConfiguration getObjectProcessConfiguration() {
            return this.objectProcessConfiguration;
        }

    }

}
