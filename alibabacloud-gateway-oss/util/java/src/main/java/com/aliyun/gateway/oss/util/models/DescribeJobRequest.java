// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DescribeJobRequest extends TeaModel {
    @NameInMap("batchJobId")
    public String batchJobId;

    public static DescribeJobRequest build(java.util.Map<String, ?> map) throws Exception {
        DescribeJobRequest self = new DescribeJobRequest();
        return TeaModel.build(map, self);
    }

    public DescribeJobRequest setBatchJobId(String batchJobId) {
        this.batchJobId = batchJobId;
        return this;
    }
    public String getBatchJobId() {
        return this.batchJobId;
    }

}
