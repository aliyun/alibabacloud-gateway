// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketRequestPaymentResponseBody extends TeaModel {
    /**
     * <p>The payer of the request and traffic fees. Valid values: **BucketOwner** and **Requester**.</p>
     */
    @NameInMap("Payer")
    public String payer;

    public static GetBucketRequestPaymentResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketRequestPaymentResponseBody self = new GetBucketRequestPaymentResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketRequestPaymentResponseBody setPayer(String payer) {
        this.payer = payer;
        return this;
    }
    public String getPayer() {
        return this.payer;
    }

}
