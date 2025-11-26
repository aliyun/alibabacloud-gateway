// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.hcs_mgw_models;

import com.aliyun.tea.*;

public class TimeFilter extends TeaModel {
    @NameInMap("EndTime")
    public String endTime;

    @NameInMap("StartTime")
    public String startTime;

    public static TimeFilter build(java.util.Map<String, ?> map) throws Exception {
        TimeFilter self = new TimeFilter();
        return TeaModel.build(map, self);
    }

    public TimeFilter setEndTime(String endTime) {
        this.endTime = endTime;
        return this;
    }
    public String getEndTime() {
        return this.endTime;
    }

    public TimeFilter setStartTime(String startTime) {
        this.startTime = startTime;
        return this;
    }
    public String getStartTime() {
        return this.startTime;
    }

}
