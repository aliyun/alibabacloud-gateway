// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketVersioningResponseBody extends TeaModel {
    /**
     * <p>The versioning state. Valid values:</p>
     * <br>
     * <p>*   Enabled</p>
     * <p>*   Suspended</p>
     */
    @NameInMap("Status")
    public String status;

    public static GetBucketVersioningResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketVersioningResponseBody self = new GetBucketVersioningResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketVersioningResponseBody setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

}
