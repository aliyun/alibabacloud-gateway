// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListReservedCapacityResponseBody extends TeaModel {
    @NameInMap("ReservedCapacityRecordList")
    public ReservedCapacityRecordList reservedCapacityRecordList;

    public static ListReservedCapacityResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListReservedCapacityResponseBody self = new ListReservedCapacityResponseBody();
        return TeaModel.build(map, self);
    }

    public ListReservedCapacityResponseBody setReservedCapacityRecordList(ReservedCapacityRecordList reservedCapacityRecordList) {
        this.reservedCapacityRecordList = reservedCapacityRecordList;
        return this;
    }
    public ReservedCapacityRecordList getReservedCapacityRecordList() {
        return this.reservedCapacityRecordList;
    }

}
