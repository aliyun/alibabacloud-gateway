// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketAccessMonitorResponseBody extends TeaModel {
    /**
     * <p>The access tracking status of the bucket.</p>
     */
    @NameInMap("Status")
    public String status;

    public static GetBucketAccessMonitorResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketAccessMonitorResponseBody self = new GetBucketAccessMonitorResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketAccessMonitorResponseBody setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

}
