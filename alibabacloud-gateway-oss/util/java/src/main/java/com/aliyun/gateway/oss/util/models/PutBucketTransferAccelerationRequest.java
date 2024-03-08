// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketTransferAccelerationRequest extends TeaModel {
    /**
     * <p>The container that stores the transfer acceleration configurations.</p>
     */
    @NameInMap("body")
    public TransferAccelerationConfiguration body;

    public static PutBucketTransferAccelerationRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketTransferAccelerationRequest self = new PutBucketTransferAccelerationRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketTransferAccelerationRequest setBody(TransferAccelerationConfiguration body) {
        this.body = body;
        return this;
    }
    public TransferAccelerationConfiguration getBody() {
        return this.body;
    }

}
