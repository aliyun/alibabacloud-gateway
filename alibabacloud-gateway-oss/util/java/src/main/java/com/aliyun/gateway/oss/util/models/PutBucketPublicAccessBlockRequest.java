// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketPublicAccessBlockRequest extends TeaModel {
    @NameInMap("body")
    public PublicAccessBlockConfiguration body;

    public static PutBucketPublicAccessBlockRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketPublicAccessBlockRequest self = new PutBucketPublicAccessBlockRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketPublicAccessBlockRequest setBody(PublicAccessBlockConfiguration body) {
        this.body = body;
        return this;
    }
    public PublicAccessBlockConfiguration getBody() {
        return this.body;
    }

}
