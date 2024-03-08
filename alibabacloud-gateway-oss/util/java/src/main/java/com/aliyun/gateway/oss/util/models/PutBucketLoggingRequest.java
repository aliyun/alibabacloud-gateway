// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketLoggingRequest extends TeaModel {
    /**
     * <p>The container that stores the information about the logging status.</p>
     */
    @NameInMap("body")
    public BucketLoggingStatus body;

    public static PutBucketLoggingRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketLoggingRequest self = new PutBucketLoggingRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketLoggingRequest setBody(BucketLoggingStatus body) {
        this.body = body;
        return this;
    }
    public BucketLoggingStatus getBody() {
        return this.body;
    }

}
