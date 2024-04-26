// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketDataRedundancyTransitionRequest extends TeaModel {
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
