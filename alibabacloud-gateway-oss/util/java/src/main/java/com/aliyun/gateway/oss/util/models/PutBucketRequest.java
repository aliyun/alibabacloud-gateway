// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketRequest extends TeaModel {
    /**
     * <p>The container that stores the information about the bucket to be created.</p>
     */
    @NameInMap("body")
    public CreateBucketConfiguration body;

    public static PutBucketRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketRequest self = new PutBucketRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketRequest setBody(CreateBucketConfiguration body) {
        this.body = body;
        return this;
    }
    public CreateBucketConfiguration getBody() {
        return this.body;
    }

}
