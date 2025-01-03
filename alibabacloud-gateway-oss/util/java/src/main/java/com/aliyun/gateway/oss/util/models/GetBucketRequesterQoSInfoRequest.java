// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketRequesterQoSInfoRequest extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>26753xxxxxxxx14340</p>
     */
    @NameInMap("qosRequester")
    public String qosRequester;

    public static GetBucketRequesterQoSInfoRequest build(java.util.Map<String, ?> map) throws Exception {
        GetBucketRequesterQoSInfoRequest self = new GetBucketRequesterQoSInfoRequest();
        return TeaModel.build(map, self);
    }

    public GetBucketRequesterQoSInfoRequest setQosRequester(String qosRequester) {
        this.qosRequester = qosRequester;
        return this;
    }
    public String getQosRequester() {
        return this.qosRequester;
    }

}
