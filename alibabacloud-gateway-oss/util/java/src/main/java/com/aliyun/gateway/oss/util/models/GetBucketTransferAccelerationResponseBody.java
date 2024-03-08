// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketTransferAccelerationResponseBody extends TeaModel {
    /**
     * <p>Indicates whether the transfer acceleration feature is enabled. Valid values:</p>
     * <br>
     * <p>*   true</p>
     * <p>*   false</p>
     */
    @NameInMap("Enabled")
    public Boolean enabled;

    public static GetBucketTransferAccelerationResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketTransferAccelerationResponseBody self = new GetBucketTransferAccelerationResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketTransferAccelerationResponseBody setEnabled(Boolean enabled) {
        this.enabled = enabled;
        return this;
    }
    public Boolean getEnabled() {
        return this.enabled;
    }

}
