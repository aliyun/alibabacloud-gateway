// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketResponseHeaderRequest extends TeaModel {
    @NameInMap("body")
    public ResponseHeaderConfiguration body;

    public static PutBucketResponseHeaderRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketResponseHeaderRequest self = new PutBucketResponseHeaderRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketResponseHeaderRequest setBody(ResponseHeaderConfiguration body) {
        this.body = body;
        return this;
    }
    public ResponseHeaderConfiguration getBody() {
        return this.body;
    }

}
