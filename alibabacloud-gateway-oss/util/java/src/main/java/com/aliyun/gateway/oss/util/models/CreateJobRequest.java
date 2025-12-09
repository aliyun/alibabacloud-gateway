// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateJobRequest extends TeaModel {
    @NameInMap("CreateJobRequest")
    public BatchCreateJobRequest createJobRequest;

    public static CreateJobRequest build(java.util.Map<String, ?> map) throws Exception {
        CreateJobRequest self = new CreateJobRequest();
        return TeaModel.build(map, self);
    }

    public CreateJobRequest setCreateJobRequest(BatchCreateJobRequest createJobRequest) {
        this.createJobRequest = createJobRequest;
        return this;
    }
    public BatchCreateJobRequest getCreateJobRequest() {
        return this.createJobRequest;
    }

}
