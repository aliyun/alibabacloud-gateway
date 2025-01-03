// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DataAcceleratorConfiguration extends TeaModel {
    @NameInMap("AcceleratePaths")
    public AcceleratePaths acceleratePaths;

    /**
     * <strong>example:</strong>
     * <p>102400</p>
     */
    @NameInMap("Quota")
    public String quota;

    public static DataAcceleratorConfiguration build(java.util.Map<String, ?> map) throws Exception {
        DataAcceleratorConfiguration self = new DataAcceleratorConfiguration();
        return TeaModel.build(map, self);
    }

    public DataAcceleratorConfiguration setAcceleratePaths(AcceleratePaths acceleratePaths) {
        this.acceleratePaths = acceleratePaths;
        return this;
    }
    public AcceleratePaths getAcceleratePaths() {
        return this.acceleratePaths;
    }

    public DataAcceleratorConfiguration setQuota(String quota) {
        this.quota = quota;
        return this;
    }
    public String getQuota() {
        return this.quota;
    }

}
