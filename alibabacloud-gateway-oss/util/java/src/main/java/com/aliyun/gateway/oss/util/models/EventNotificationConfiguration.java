// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class EventNotificationConfiguration extends TeaModel {
    @NameInMap("FunctionComputeConfiguration")
    public java.util.List<FunctionComputeConfiguration> functionComputeConfiguration;

    public static EventNotificationConfiguration build(java.util.Map<String, ?> map) throws Exception {
        EventNotificationConfiguration self = new EventNotificationConfiguration();
        return TeaModel.build(map, self);
    }

    public EventNotificationConfiguration setFunctionComputeConfiguration(java.util.List<FunctionComputeConfiguration> functionComputeConfiguration) {
        this.functionComputeConfiguration = functionComputeConfiguration;
        return this;
    }
    public java.util.List<FunctionComputeConfiguration> getFunctionComputeConfiguration() {
        return this.functionComputeConfiguration;
    }

}
