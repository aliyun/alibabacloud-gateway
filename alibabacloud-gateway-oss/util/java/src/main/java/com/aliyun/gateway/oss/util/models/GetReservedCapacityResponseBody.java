// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetReservedCapacityResponseBody extends TeaModel {
    @NameInMap("ReservedCapacityRecord")
    public ReservedCapacityRecord reservedCapacityRecord;

    public static GetReservedCapacityResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetReservedCapacityResponseBody self = new GetReservedCapacityResponseBody();
        return TeaModel.build(map, self);
    }

    public GetReservedCapacityResponseBody setReservedCapacityRecord(ReservedCapacityRecord reservedCapacityRecord) {
        this.reservedCapacityRecord = reservedCapacityRecord;
        return this;
    }
    public ReservedCapacityRecord getReservedCapacityRecord() {
        return this.reservedCapacityRecord;
    }

}
