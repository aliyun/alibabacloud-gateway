// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketTagsRequest extends TeaModel {
    /**
     * <p>The container that stores the tags.</p>
     */
    @NameInMap("body")
    public Tagging body;

    public static PutBucketTagsRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketTagsRequest self = new PutBucketTagsRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketTagsRequest setBody(Tagging body) {
        this.body = body;
        return this;
    }
    public Tagging getBody() {
        return this.body;
    }

}
