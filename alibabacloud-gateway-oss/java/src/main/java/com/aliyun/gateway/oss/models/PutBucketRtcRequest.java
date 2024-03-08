// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class PutBucketRtcRequest extends TeaModel {
    /**
     * <p>The container that stores the RTC configurations.</p>
     */
    @NameInMap("ReplicationRule")
    public RtcConfiguration replicationRule;

    public static PutBucketRtcRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketRtcRequest self = new PutBucketRtcRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketRtcRequest setReplicationRule(RtcConfiguration replicationRule) {
        this.replicationRule = replicationRule;
        return this;
    }
    public RtcConfiguration getReplicationRule() {
        return this.replicationRule;
    }

}
