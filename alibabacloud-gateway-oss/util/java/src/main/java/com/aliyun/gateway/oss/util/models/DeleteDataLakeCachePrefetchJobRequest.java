// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DeleteDataLakeCachePrefetchJobRequest extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("x-oss-datalake-job-id")
    public String xOssDatalakeJobId;

    public static DeleteDataLakeCachePrefetchJobRequest build(java.util.Map<String, ?> map) throws Exception {
        DeleteDataLakeCachePrefetchJobRequest self = new DeleteDataLakeCachePrefetchJobRequest();
        return TeaModel.build(map, self);
    }

    public DeleteDataLakeCachePrefetchJobRequest setXOssDatalakeJobId(String xOssDatalakeJobId) {
        this.xOssDatalakeJobId = xOssDatalakeJobId;
        return this;
    }
    public String getXOssDatalakeJobId() {
        return this.xOssDatalakeJobId;
    }

}
