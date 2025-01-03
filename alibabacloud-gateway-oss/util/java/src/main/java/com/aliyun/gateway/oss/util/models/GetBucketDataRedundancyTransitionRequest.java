// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketDataRedundancyTransitionRequest extends TeaModel {
    /**
     * <p>The ID of the redundancy change task.</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>751f5243f8ac4ae89f34726534d1****</p>
     */
    @NameInMap("x-oss-redundancy-transition-taskid")
    public String xOssRedundancyTransitionTaskid;

    public static GetBucketDataRedundancyTransitionRequest build(java.util.Map<String, ?> map) throws Exception {
        GetBucketDataRedundancyTransitionRequest self = new GetBucketDataRedundancyTransitionRequest();
        return TeaModel.build(map, self);
    }

    public GetBucketDataRedundancyTransitionRequest setXOssRedundancyTransitionTaskid(String xOssRedundancyTransitionTaskid) {
        this.xOssRedundancyTransitionTaskid = xOssRedundancyTransitionTaskid;
        return this;
    }
    public String getXOssRedundancyTransitionTaskid() {
        return this.xOssRedundancyTransitionTaskid;
    }

}
