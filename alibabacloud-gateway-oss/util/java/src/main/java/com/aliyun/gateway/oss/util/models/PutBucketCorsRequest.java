// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketCorsRequest extends TeaModel {
    /**
     * <p>The container that stores the information about CORS rules of the bucket.</p>
     */
    @NameInMap("body")
    public CORSConfiguration body;

    public static PutBucketCorsRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketCorsRequest self = new PutBucketCorsRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketCorsRequest setBody(CORSConfiguration body) {
        this.body = body;
        return this;
    }
    public CORSConfiguration getBody() {
        return this.body;
    }

}
