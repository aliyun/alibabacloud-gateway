// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BatchRestoreObject extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>5</p>
     */
    @NameInMap("Days")
    public Integer days;

    /**
     * <strong>example:</strong>
     * <p>Standard</p>
     */
    @NameInMap("Tier")
    public String tier;

    public static BatchRestoreObject build(java.util.Map<String, ?> map) throws Exception {
        BatchRestoreObject self = new BatchRestoreObject();
        return TeaModel.build(map, self);
    }

    public BatchRestoreObject setDays(Integer days) {
        this.days = days;
        return this;
    }
    public Integer getDays() {
        return this.days;
    }

    public BatchRestoreObject setTier(String tier) {
        this.tier = tier;
        return this;
    }
    public String getTier() {
        return this.tier;
    }

}
