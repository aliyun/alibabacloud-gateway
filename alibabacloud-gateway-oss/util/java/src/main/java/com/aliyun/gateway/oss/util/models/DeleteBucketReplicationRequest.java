// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class DeleteBucketReplicationRequest extends TeaModel {
    /**
     * <p>The container that stores the data replication rules to be deleted.</p>
     */
    @NameInMap("body")
    public ReplicationRules body;

    public static DeleteBucketReplicationRequest build(java.util.Map<String, ?> map) throws Exception {
        DeleteBucketReplicationRequest self = new DeleteBucketReplicationRequest();
        return TeaModel.build(map, self);
    }

    public DeleteBucketReplicationRequest setBody(ReplicationRules body) {
        this.body = body;
        return this;
    }
    public ReplicationRules getBody() {
        return this.body;
    }

}
