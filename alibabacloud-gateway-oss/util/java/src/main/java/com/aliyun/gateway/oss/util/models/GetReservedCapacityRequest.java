// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetReservedCapacityRequest extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("x-oss-reserved-capacity-id")
    public String xOssReservedCapacityId;

    public static GetReservedCapacityRequest build(java.util.Map<String, ?> map) throws Exception {
        GetReservedCapacityRequest self = new GetReservedCapacityRequest();
        return TeaModel.build(map, self);
    }

    public GetReservedCapacityRequest setXOssReservedCapacityId(String xOssReservedCapacityId) {
        this.xOssReservedCapacityId = xOssReservedCapacityId;
        return this;
    }
    public String getXOssReservedCapacityId() {
        return this.xOssReservedCapacityId;
    }

}
