// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BatchOperationFailureReasons extends TeaModel {
    @NameInMap("FailureCode")
    public String failureCode;

    @NameInMap("FailureReason")
    public String failureReason;

    public static BatchOperationFailureReasons build(java.util.Map<String, ?> map) throws Exception {
        BatchOperationFailureReasons self = new BatchOperationFailureReasons();
        return TeaModel.build(map, self);
    }

    public BatchOperationFailureReasons setFailureCode(String failureCode) {
        this.failureCode = failureCode;
        return this;
    }
    public String getFailureCode() {
        return this.failureCode;
    }

    public BatchOperationFailureReasons setFailureReason(String failureReason) {
        this.failureReason = failureReason;
        return this;
    }
    public String getFailureReason() {
        return this.failureReason;
    }

}
