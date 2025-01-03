// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetDataLakeCachePrefetchJobRequest extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("x-oss-datalake-job-id")
    public String xOssDatalakeJobId;

    public static GetDataLakeCachePrefetchJobRequest build(java.util.Map<String, ?> map) throws Exception {
        GetDataLakeCachePrefetchJobRequest self = new GetDataLakeCachePrefetchJobRequest();
        return TeaModel.build(map, self);
    }

    public GetDataLakeCachePrefetchJobRequest setXOssDatalakeJobId(String xOssDatalakeJobId) {
        this.xOssDatalakeJobId = xOssDatalakeJobId;
        return this;
    }
    public String getXOssDatalakeJobId() {
        return this.xOssDatalakeJobId;
    }

}
