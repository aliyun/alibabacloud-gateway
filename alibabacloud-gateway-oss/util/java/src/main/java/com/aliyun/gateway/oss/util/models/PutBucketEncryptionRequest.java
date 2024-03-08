// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketEncryptionRequest extends TeaModel {
    /**
     * <p>The container that stores server-side encryption rules.</p>
     */
    @NameInMap("body")
    public ServerSideEncryptionRule body;

    public static PutBucketEncryptionRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketEncryptionRequest self = new PutBucketEncryptionRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketEncryptionRequest setBody(ServerSideEncryptionRule body) {
        this.body = body;
        return this;
    }
    public ServerSideEncryptionRule getBody() {
        return this.body;
    }

}
