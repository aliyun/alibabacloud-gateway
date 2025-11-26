// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.hcs_mgw_models;

import com.aliyun.tea.*;

public class ScheduleRule extends TeaModel {
    @NameInMap("MaxScheduleCount")
    public Long maxScheduleCount;

    @NameInMap("StartCronExpression")
    public String startCronExpression;

    @NameInMap("SuspendCronExpression")
    public String suspendCronExpression;

    public static ScheduleRule build(java.util.Map<String, ?> map) throws Exception {
        ScheduleRule self = new ScheduleRule();
        return TeaModel.build(map, self);
    }

    public ScheduleRule setMaxScheduleCount(Long maxScheduleCount) {
        this.maxScheduleCount = maxScheduleCount;
        return this;
    }
    public Long getMaxScheduleCount() {
        return this.maxScheduleCount;
    }

    public ScheduleRule setStartCronExpression(String startCronExpression) {
        this.startCronExpression = startCronExpression;
        return this;
    }
    public String getStartCronExpression() {
        return this.startCronExpression;
    }

    public ScheduleRule setSuspendCronExpression(String suspendCronExpression) {
        this.suspendCronExpression = suspendCronExpression;
        return this;
    }
    public String getSuspendCronExpression() {
        return this.suspendCronExpression;
    }

}
