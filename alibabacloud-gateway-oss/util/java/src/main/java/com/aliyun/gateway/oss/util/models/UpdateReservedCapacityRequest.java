// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class UpdateReservedCapacityRequest extends TeaModel {
    @NameInMap("ReservedCapacityUpdateConfiguration")
    public ReservedCapacityUpdateConfiguration reservedCapacityUpdateConfiguration;

    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("x-oss-reserved-capacity-id")
    public String xOssReservedCapacityId;

    public static UpdateReservedCapacityRequest build(java.util.Map<String, ?> map) throws Exception {
        UpdateReservedCapacityRequest self = new UpdateReservedCapacityRequest();
        return TeaModel.build(map, self);
    }

    public UpdateReservedCapacityRequest setReservedCapacityUpdateConfiguration(ReservedCapacityUpdateConfiguration reservedCapacityUpdateConfiguration) {
        this.reservedCapacityUpdateConfiguration = reservedCapacityUpdateConfiguration;
        return this;
    }
    public ReservedCapacityUpdateConfiguration getReservedCapacityUpdateConfiguration() {
        return this.reservedCapacityUpdateConfiguration;
    }

    public UpdateReservedCapacityRequest setXOssReservedCapacityId(String xOssReservedCapacityId) {
        this.xOssReservedCapacityId = xOssReservedCapacityId;
        return this;
    }
    public String getXOssReservedCapacityId() {
        return this.xOssReservedCapacityId;
    }

}
