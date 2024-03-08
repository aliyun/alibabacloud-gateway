// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketReplicationRequest extends TeaModel {
    /**
     * <p>The configurations for data replication.</p>
     */
    @NameInMap("body")
    public ReplicationConfiguration body;

    public static PutBucketReplicationRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketReplicationRequest self = new PutBucketReplicationRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketReplicationRequest setBody(ReplicationConfiguration body) {
        this.body = body;
        return this;
    }
    public ReplicationConfiguration getBody() {
        return this.body;
    }

}
