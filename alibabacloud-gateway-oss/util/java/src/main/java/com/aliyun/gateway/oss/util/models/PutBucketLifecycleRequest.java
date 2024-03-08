// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketLifecycleRequest extends TeaModel {
    /**
     * <p>The container that stores the lifecycle configurations. The container can contain up to 1,000 lifecycle rules.</p>
     */
    @NameInMap("body")
    public LifecycleConfiguration body;

    public static PutBucketLifecycleRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketLifecycleRequest self = new PutBucketLifecycleRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketLifecycleRequest setBody(LifecycleConfiguration body) {
        this.body = body;
        return this;
    }
    public LifecycleConfiguration getBody() {
        return this.body;
    }

}
