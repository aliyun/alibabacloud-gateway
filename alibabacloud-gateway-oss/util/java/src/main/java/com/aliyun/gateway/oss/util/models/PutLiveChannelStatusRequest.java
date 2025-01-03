// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutLiveChannelStatusRequest extends TeaModel {
    /**
     * <p>The status of the LiveChannel. 
     * Valid values:</p>
     * <ul>
     * <li>enabled: enables the LiveChannel.</li>
     * <li>disabled: disables the LiveChannel.</li>
     * </ul>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>enabled</p>
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
