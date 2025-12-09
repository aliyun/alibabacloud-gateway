// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateJobResponseBody extends TeaModel {
    @NameInMap("CreateJobResult")
    public CreateJobResult createJobResult;

    public static CreateJobResponseBody build(java.util.Map<String, ?> map) throws Exception {
        CreateJobResponseBody self = new CreateJobResponseBody();
        return TeaModel.build(map, self);
    }

    public CreateJobResponseBody setCreateJobResult(CreateJobResult createJobResult) {
        this.createJobResult = createJobResult;
        return this;
    }
    public CreateJobResult getCreateJobResult() {
        return this.createJobResult;
    }

}
