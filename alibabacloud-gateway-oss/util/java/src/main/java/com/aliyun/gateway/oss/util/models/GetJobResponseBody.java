// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetJobResponseBody extends TeaModel {
    @NameInMap("ImportJob")
    public GetJobResp importJob;

    public static GetJobResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetJobResponseBody self = new GetJobResponseBody();
        return TeaModel.build(map, self);
    }

    public GetJobResponseBody setImportJob(GetJobResp importJob) {
        this.importJob = importJob;
        return this;
    }
    public GetJobResp getImportJob() {
        return this.importJob;
    }

}
