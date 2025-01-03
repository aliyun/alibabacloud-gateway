// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DeleteBucketRequesterQoSInfoRequest extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>26753xxxxxxxx14340</p>
     */
    @NameInMap("qosRequester")
    public String qosRequester;

    public static DeleteBucketRequesterQoSInfoRequest build(java.util.Map<String, ?> map) throws Exception {
        DeleteBucketRequesterQoSInfoRequest self = new DeleteBucketRequesterQoSInfoRequest();
        return TeaModel.build(map, self);
    }

    public DeleteBucketRequesterQoSInfoRequest setQosRequester(String qosRequester) {
        this.qosRequester = qosRequester;
        return this;
    }
    public String getQosRequester() {
        return this.qosRequester;
    }

}
