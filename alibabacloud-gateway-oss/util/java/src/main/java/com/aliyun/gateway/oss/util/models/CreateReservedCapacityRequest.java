// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateReservedCapacityRequest extends TeaModel {
    @NameInMap("ReservedCapacityCreateConfiguration")
    public ReservedCapacityCreateConfiguration reservedCapacityCreateConfiguration;

    public static CreateReservedCapacityRequest build(java.util.Map<String, ?> map) throws Exception {
        CreateReservedCapacityRequest self = new CreateReservedCapacityRequest();
        return TeaModel.build(map, self);
    }

    public CreateReservedCapacityRequest setReservedCapacityCreateConfiguration(ReservedCapacityCreateConfiguration reservedCapacityCreateConfiguration) {
        this.reservedCapacityCreateConfiguration = reservedCapacityCreateConfiguration;
        return this;
    }
    public ReservedCapacityCreateConfiguration getReservedCapacityCreateConfiguration() {
        return this.reservedCapacityCreateConfiguration;
    }

}
