// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class TransferAccelerationConfiguration extends TeaModel {
    @NameInMap("Enabled")
    public Boolean enabled;

    public static TransferAccelerationConfiguration build(java.util.Map<String, ?> map) throws Exception {
        TransferAccelerationConfiguration self = new TransferAccelerationConfiguration();
        return TeaModel.build(map, self);
    }

    public TransferAccelerationConfiguration setEnabled(Boolean enabled) {
        this.enabled = enabled;
        return this;
    }
    public Boolean getEnabled() {
        return this.enabled;
    }

}
