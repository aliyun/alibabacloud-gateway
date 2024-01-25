// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketRtcRequest extends TeaModel {
    /**
     * <p>The container that stores the RTC configurations.</p>
     */
    @NameInMap("body")
    public RtcConfiguration body;

    public static PutBucketRtcRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketRtcRequest self = new PutBucketRtcRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketRtcRequest setBody(RtcConfiguration body) {
        this.body = body;
        return this;
    }
    public RtcConfiguration getBody() {
        return this.body;
    }

}
