// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateReservedCapacityResponseBody extends TeaModel {
    @NameInMap("CreateLargeReservedCapacityResult")
    public CreateLargeReservedCapacityResult createLargeReservedCapacityResult;

    public static CreateReservedCapacityResponseBody build(java.util.Map<String, ?> map) throws Exception {
        CreateReservedCapacityResponseBody self = new CreateReservedCapacityResponseBody();
        return TeaModel.build(map, self);
    }

    public CreateReservedCapacityResponseBody setCreateLargeReservedCapacityResult(CreateLargeReservedCapacityResult createLargeReservedCapacityResult) {
        this.createLargeReservedCapacityResult = createLargeReservedCapacityResult;
        return this;
    }
    public CreateLargeReservedCapacityResult getCreateLargeReservedCapacityResult() {
        return this.createLargeReservedCapacityResult;
    }

}
