// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketPolicyRequest extends TeaModel {
    /**
     * <p>The request parameters.</p>
     */
    @NameInMap("body")
    public String body;

    public static PutBucketPolicyRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketPolicyRequest self = new PutBucketPolicyRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketPolicyRequest setBody(String body) {
        this.body = body;
        return this;
    }
    public String getBody() {
        return this.body;
    }

}
