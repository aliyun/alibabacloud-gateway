// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketRefererRequest extends TeaModel {
    /**
     * <p>The container that stores the Referer configurations.</p>
     */
    @NameInMap("body")
    public RefererConfiguration body;

    public static PutBucketRefererRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketRefererRequest self = new PutBucketRefererRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketRefererRequest setBody(RefererConfiguration body) {
        this.body = body;
        return this;
    }
    public RefererConfiguration getBody() {
        return this.body;
    }

}
