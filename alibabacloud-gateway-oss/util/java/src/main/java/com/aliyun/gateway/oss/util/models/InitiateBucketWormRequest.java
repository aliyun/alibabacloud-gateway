// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class InitiateBucketWormRequest extends TeaModel {
    /**
     * <p>The request parameters.</p>
     */
    @NameInMap("body")
    public InitiateWormConfiguration body;

    public static InitiateBucketWormRequest build(java.util.Map<String, ?> map) throws Exception {
        InitiateBucketWormRequest self = new InitiateBucketWormRequest();
        return TeaModel.build(map, self);
    }

    public InitiateBucketWormRequest setBody(InitiateWormConfiguration body) {
        this.body = body;
        return this;
    }
    public InitiateWormConfiguration getBody() {
        return this.body;
    }

}
