// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class AccessMonitorConfiguration extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("AllowCopy")
    public Boolean allowCopy;

    @NameInMap("Status")
    public String status;

    public static AccessMonitorConfiguration build(java.util.Map<String, ?> map) throws Exception {
        AccessMonitorConfiguration self = new AccessMonitorConfiguration();
        return TeaModel.build(map, self);
    }

    public AccessMonitorConfiguration setAllowCopy(Boolean allowCopy) {
        this.allowCopy = allowCopy;
        return this;
    }
    public Boolean getAllowCopy() {
        return this.allowCopy;
    }

    public AccessMonitorConfiguration setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

}
