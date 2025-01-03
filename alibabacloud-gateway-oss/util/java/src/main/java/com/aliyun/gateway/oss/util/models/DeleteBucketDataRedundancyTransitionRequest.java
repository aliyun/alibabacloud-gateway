// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DeleteBucketDataRedundancyTransitionRequest extends TeaModel {
    /**
     * <p>The ID of the redundancy type change task.</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>4be5beb0f74f490186311b268bf6****</p>
     */
    @NameInMap("x-oss-redundancy-transition-taskid")
    public String xOssRedundancyTransitionTaskid;

    public static DeleteBucketDataRedundancyTransitionRequest build(java.util.Map<String, ?> map) throws Exception {
        DeleteBucketDataRedundancyTransitionRequest self = new DeleteBucketDataRedundancyTransitionRequest();
        return TeaModel.build(map, self);
    }

    public DeleteBucketDataRedundancyTransitionRequest setXOssRedundancyTransitionTaskid(String xOssRedundancyTransitionTaskid) {
        this.xOssRedundancyTransitionTaskid = xOssRedundancyTransitionTaskid;
        return this;
    }
    public String getXOssRedundancyTransitionTaskid() {
        return this.xOssRedundancyTransitionTaskid;
    }

}
