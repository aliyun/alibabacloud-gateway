// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetUserQoSInfoResponseBody extends TeaModel {
    @NameInMap("QoSConfiguration")
    public UserQosConfiguration qoSConfiguration;

    public static GetUserQoSInfoResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetUserQoSInfoResponseBody self = new GetUserQoSInfoResponseBody();
        return TeaModel.build(map, self);
    }

    public GetUserQoSInfoResponseBody setQoSConfiguration(UserQosConfiguration qoSConfiguration) {
        this.qoSConfiguration = qoSConfiguration;
        return this;
    }
    public UserQosConfiguration getQoSConfiguration() {
        return this.qoSConfiguration;
    }

}
