// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketVersioningRequest extends TeaModel {
    /**
     * <p>The request parameters.</p>
     */
    @NameInMap("body")
    public VersioningConfiguration body;

    public static PutBucketVersioningRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketVersioningRequest self = new PutBucketVersioningRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketVersioningRequest setBody(VersioningConfiguration body) {
        this.body = body;
        return this;
    }
    public VersioningConfiguration getBody() {
        return this.body;
    }

}
