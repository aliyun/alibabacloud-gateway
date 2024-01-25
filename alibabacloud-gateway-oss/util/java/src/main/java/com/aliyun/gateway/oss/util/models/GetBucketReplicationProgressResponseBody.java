// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketReplicationProgressResponseBody extends TeaModel {
    /**
     * <p>The container that stores the progress of the data replication task corresponding to the specified data replication rule.</p>
     */
    @NameInMap("Rule")
    public ReplicationProgressRule rule;

    public static GetBucketReplicationProgressResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketReplicationProgressResponseBody self = new GetBucketReplicationProgressResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketReplicationProgressResponseBody setRule(ReplicationProgressRule rule) {
        this.rule = rule;
        return this;
    }
    public ReplicationProgressRule getRule() {
        return this.rule;
    }

}
