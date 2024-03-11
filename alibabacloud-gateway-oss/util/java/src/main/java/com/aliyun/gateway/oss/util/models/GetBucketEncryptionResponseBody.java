// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class GetBucketEncryptionResponseBody extends TeaModel {
    /**
     * <p>The container that stores server-side encryption rules.</p>
     */
    @NameInMap("ServerSideEncryptionRule")
    public ServerSideEncryptionRule serverSideEncryptionRule;

    public static GetBucketEncryptionResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketEncryptionResponseBody self = new GetBucketEncryptionResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketEncryptionResponseBody setServerSideEncryptionRule(ServerSideEncryptionRule serverSideEncryptionRule) {
        this.serverSideEncryptionRule = serverSideEncryptionRule;
        return this;
    }
    public ServerSideEncryptionRule getServerSideEncryptionRule() {
        return this.serverSideEncryptionRule;
    }

    public static class ServerSideEncryptionRule extends TeaModel {
        /**
         * <p>The container that stores the default server-side encryption method.</p>
         */
        @NameInMap("ApplyServerSideEncryptionByDefault")
        public ApplyServerSideEncryptionByDefault applyServerSideEncryptionByDefault;

        public static ServerSideEncryptionRule build(java.util.Map<String, ?> map) throws Exception {
            ServerSideEncryptionRule self = new ServerSideEncryptionRule();
            return TeaModel.build(map, self);
        }

        public ServerSideEncryptionRule setApplyServerSideEncryptionByDefault(ApplyServerSideEncryptionByDefault applyServerSideEncryptionByDefault) {
            this.applyServerSideEncryptionByDefault = applyServerSideEncryptionByDefault;
            return this;
        }
        public ApplyServerSideEncryptionByDefault getApplyServerSideEncryptionByDefault() {
            return this.applyServerSideEncryptionByDefault;
        }

    }

}
