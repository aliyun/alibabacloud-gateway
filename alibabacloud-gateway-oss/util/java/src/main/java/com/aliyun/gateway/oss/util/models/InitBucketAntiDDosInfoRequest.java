// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class InitBucketAntiDDosInfoRequest extends TeaModel {
    /**
     * <p>The container that stores the configurations of the Anti-DDoS instance.</p>
     */
    @NameInMap("body")
    public BucketAntiDDOSConfiguration body;

    public static InitBucketAntiDDosInfoRequest build(java.util.Map<String, ?> map) throws Exception {
        InitBucketAntiDDosInfoRequest self = new InitBucketAntiDDosInfoRequest();
        return TeaModel.build(map, self);
    }

    public InitBucketAntiDDosInfoRequest setBody(BucketAntiDDOSConfiguration body) {
        this.body = body;
        return this;
    }
    public BucketAntiDDOSConfiguration getBody() {
        return this.body;
    }

}
