// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class UpdateBucketAntiDDosInfoRequest extends TeaModel {
    /**
     * <p>The container that stores the configurations of the Anti-DDoS instance.</p>
     */
    @NameInMap("body")
    public BucketAntiDDOSConfiguration body;

    public static UpdateBucketAntiDDosInfoRequest build(java.util.Map<String, ?> map) throws Exception {
        UpdateBucketAntiDDosInfoRequest self = new UpdateBucketAntiDDosInfoRequest();
        return TeaModel.build(map, self);
    }

    public UpdateBucketAntiDDosInfoRequest setBody(BucketAntiDDOSConfiguration body) {
        this.body = body;
        return this;
    }
    public BucketAntiDDOSConfiguration getBody() {
        return this.body;
    }

}
