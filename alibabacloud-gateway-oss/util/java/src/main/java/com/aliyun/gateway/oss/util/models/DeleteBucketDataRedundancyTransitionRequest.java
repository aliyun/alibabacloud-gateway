// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class DeleteBucketDataRedundancyTransitionRequest extends TeaModel {
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
