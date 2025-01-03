// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DeleteReservedCapacityRequest extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("x-oss-reserved-capacity-id")
    public String xOssReservedCapacityId;

    public static DeleteReservedCapacityRequest build(java.util.Map<String, ?> map) throws Exception {
        DeleteReservedCapacityRequest self = new DeleteReservedCapacityRequest();
        return TeaModel.build(map, self);
    }

    public DeleteReservedCapacityRequest setXOssReservedCapacityId(String xOssReservedCapacityId) {
        this.xOssReservedCapacityId = xOssReservedCapacityId;
        return this;
    }
    public String getXOssReservedCapacityId() {
        return this.xOssReservedCapacityId;
    }

}
