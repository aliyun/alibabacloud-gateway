// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutLiveChannelStatusRequest extends TeaModel {
    /**
     * <p>The status of the LiveChannel. </p>
     * <p>Valid values:</p>
     * <p>- enabled: enables the LiveChannel.</p>
     * <p>- disabled: disables the LiveChannel.</p>
     */
    @NameInMap("status")
    public String status;

    public static PutLiveChannelStatusRequest build(java.util.Map<String, ?> map) throws Exception {
        PutLiveChannelStatusRequest self = new PutLiveChannelStatusRequest();
        return TeaModel.build(map, self);
    }

    public PutLiveChannelStatusRequest setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

}
