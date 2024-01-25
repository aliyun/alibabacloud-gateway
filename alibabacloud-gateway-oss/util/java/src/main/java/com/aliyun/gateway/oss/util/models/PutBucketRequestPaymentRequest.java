// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketRequestPaymentRequest extends TeaModel {
    /**
     * <p>The container that stores pay-by-requester configurations.</p>
     */
    @NameInMap("body")
    public RequestPaymentConfiguration body;

    public static PutBucketRequestPaymentRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketRequestPaymentRequest self = new PutBucketRequestPaymentRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketRequestPaymentRequest setBody(RequestPaymentConfiguration body) {
        this.body = body;
        return this;
    }
    public RequestPaymentConfiguration getBody() {
        return this.body;
    }

}
