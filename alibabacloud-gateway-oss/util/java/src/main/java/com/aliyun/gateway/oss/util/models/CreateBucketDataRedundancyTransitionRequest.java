// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class CreateBucketDataRedundancyTransitionRequest extends TeaModel {
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
