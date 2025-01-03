// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateBucketDataRedundancyTransitionRequest extends TeaModel {
    /**
     * <p>The redundancy type to which you want to convert the bucket. You can only convert the redundancy type of a bucket from LRS to ZRS.</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>ZRS</p>
     */
    @NameInMap("x-oss-target-redundancy-type")
    public String xOssTargetRedundancyType;

    public static CreateBucketDataRedundancyTransitionRequest build(java.util.Map<String, ?> map) throws Exception {
        CreateBucketDataRedundancyTransitionRequest self = new CreateBucketDataRedundancyTransitionRequest();
        return TeaModel.build(map, self);
    }

    public CreateBucketDataRedundancyTransitionRequest setXOssTargetRedundancyType(String xOssTargetRedundancyType) {
        this.xOssTargetRedundancyType = xOssTargetRedundancyType;
        return this;
    }
    public String getXOssTargetRedundancyType() {
        return this.xOssTargetRedundancyType;
    }

}
