// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class UpdateJobStatusResponseBody extends TeaModel {
    @NameInMap("UpdateJobStatusResult")
    public BatchOperationUpdateJobStatusResult updateJobStatusResult;

    public static UpdateJobStatusResponseBody build(java.util.Map<String, ?> map) throws Exception {
        UpdateJobStatusResponseBody self = new UpdateJobStatusResponseBody();
        return TeaModel.build(map, self);
    }

    public UpdateJobStatusResponseBody setUpdateJobStatusResult(BatchOperationUpdateJobStatusResult updateJobStatusResult) {
        this.updateJobStatusResult = updateJobStatusResult;
        return this;
    }
    public BatchOperationUpdateJobStatusResult getUpdateJobStatusResult() {
        return this.updateJobStatusResult;
    }

}
