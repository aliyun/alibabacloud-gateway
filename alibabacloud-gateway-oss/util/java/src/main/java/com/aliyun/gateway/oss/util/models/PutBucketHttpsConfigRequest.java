// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketHttpsConfigRequest extends TeaModel {
    /**
     * <p>The container that stores HTTPS configurations.</p>
     */
    @NameInMap("body")
    public HttpsConfiguration body;

    public static PutBucketHttpsConfigRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketHttpsConfigRequest self = new PutBucketHttpsConfigRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketHttpsConfigRequest setBody(HttpsConfiguration body) {
        this.body = body;
        return this;
    }
    public HttpsConfiguration getBody() {
        return this.body;
    }

}
