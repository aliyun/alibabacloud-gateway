// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutCnameRequest extends TeaModel {
    /**
     * <p>The container that stores the CNAME record.</p>
     */
    @NameInMap("body")
    public BucketCnameConfiguration body;

    public static PutCnameRequest build(java.util.Map<String, ?> map) throws Exception {
        PutCnameRequest self = new PutCnameRequest();
        return TeaModel.build(map, self);
    }

    public PutCnameRequest setBody(BucketCnameConfiguration body) {
        this.body = body;
        return this;
    }
    public BucketCnameConfiguration getBody() {
        return this.body;
    }

}
