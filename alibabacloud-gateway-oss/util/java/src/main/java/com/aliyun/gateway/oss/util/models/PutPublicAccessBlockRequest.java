// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutPublicAccessBlockRequest extends TeaModel {
    @NameInMap("body")
    public PublicAccessBlockConfiguration body;

    public static PutPublicAccessBlockRequest build(java.util.Map<String, ?> map) throws Exception {
        PutPublicAccessBlockRequest self = new PutPublicAccessBlockRequest();
        return TeaModel.build(map, self);
    }

    public PutPublicAccessBlockRequest setBody(PublicAccessBlockConfiguration body) {
        this.body = body;
        return this;
    }
    public PublicAccessBlockConfiguration getBody() {
        return this.body;
    }

}
