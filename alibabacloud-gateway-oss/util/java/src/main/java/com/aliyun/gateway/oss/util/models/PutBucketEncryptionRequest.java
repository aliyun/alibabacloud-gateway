// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class PutBucketEncryptionRequest extends TeaModel {
    /**
     * <p>The container that stores server-side encryption rules.</p>
     */
    @NameInMap("ServerSideEncryptionRule")
    public ServerSideEncryptionRule serverSideEncryptionRule;

    public static PutBucketEncryptionRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketEncryptionRequest self = new PutBucketEncryptionRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketEncryptionRequest setServerSideEncryptionRule(ServerSideEncryptionRule serverSideEncryptionRule) {
        this.serverSideEncryptionRule = serverSideEncryptionRule;
        return this;
    }
    public ServerSideEncryptionRule getServerSideEncryptionRule() {
        return this.serverSideEncryptionRule;
    }

}
