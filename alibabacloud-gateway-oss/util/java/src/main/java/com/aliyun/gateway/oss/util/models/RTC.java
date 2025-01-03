// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class RTC extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>Enabled</p>
     */
    @NameInMap("Status")
    public String status;

    public static RTC build(java.util.Map<String, ?> map) throws Exception {
        RTC self = new RTC();
        return TeaModel.build(map, self);
    }

    public RTC setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

}
