// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ReservedCapacityRecordList extends TeaModel {
    @NameInMap("ReservedCapacityRecord")
    public java.util.List<ReservedCapacityRecord> reservedCapacityRecord;

    public static ReservedCapacityRecordList build(java.util.Map<String, ?> map) throws Exception {
        ReservedCapacityRecordList self = new ReservedCapacityRecordList();
        return TeaModel.build(map, self);
    }

    public ReservedCapacityRecordList setReservedCapacityRecord(java.util.List<ReservedCapacityRecord> reservedCapacityRecord) {
        this.reservedCapacityRecord = reservedCapacityRecord;
        return this;
    }
    public java.util.List<ReservedCapacityRecord> getReservedCapacityRecord() {
        return this.reservedCapacityRecord;
    }

}
