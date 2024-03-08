// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketAccessMonitorRequest extends TeaModel {
    /**
     * <p>The access tracking configurations of the bucket.</p>
     */
    @NameInMap("body")
    public AccessMonitorConfiguration body;

    public static PutBucketAccessMonitorRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketAccessMonitorRequest self = new PutBucketAccessMonitorRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketAccessMonitorRequest setBody(AccessMonitorConfiguration body) {
        this.body = body;
        return this;
    }
    public AccessMonitorConfiguration getBody() {
        return this.body;
    }

}
