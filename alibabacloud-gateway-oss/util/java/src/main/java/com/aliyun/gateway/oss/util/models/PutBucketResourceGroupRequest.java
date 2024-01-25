// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketResourceGroupRequest extends TeaModel {
    /**
     * <p>The container that stores the ID of the resource group.</p>
     */
    @NameInMap("body")
    public BucketResourceGroupConfiguration body;

    public static PutBucketResourceGroupRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketResourceGroupRequest self = new PutBucketResourceGroupRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketResourceGroupRequest setBody(BucketResourceGroupConfiguration body) {
        this.body = body;
        return this;
    }
    public BucketResourceGroupConfiguration getBody() {
        return this.body;
    }

}
