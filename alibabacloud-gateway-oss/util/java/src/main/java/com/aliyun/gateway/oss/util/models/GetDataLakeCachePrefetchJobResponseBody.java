// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetDataLakeCachePrefetchJobResponseBody extends TeaModel {
    @NameInMap("DataLakeCachePrefetchJob")
    public DataLakeCachePrefetchJob dataLakeCachePrefetchJob;

    public static GetDataLakeCachePrefetchJobResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetDataLakeCachePrefetchJobResponseBody self = new GetDataLakeCachePrefetchJobResponseBody();
        return TeaModel.build(map, self);
    }

    public GetDataLakeCachePrefetchJobResponseBody setDataLakeCachePrefetchJob(DataLakeCachePrefetchJob dataLakeCachePrefetchJob) {
        this.dataLakeCachePrefetchJob = dataLakeCachePrefetchJob;
        return this;
    }
    public DataLakeCachePrefetchJob getDataLakeCachePrefetchJob() {
        return this.dataLakeCachePrefetchJob;
    }

}
