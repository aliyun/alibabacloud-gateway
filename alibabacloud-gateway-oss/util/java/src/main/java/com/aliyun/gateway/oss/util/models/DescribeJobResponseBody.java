// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DescribeJobResponseBody extends TeaModel {
    @NameInMap("DescribeJobResult")
    public DescribeJobResult describeJobResult;

    public static DescribeJobResponseBody build(java.util.Map<String, ?> map) throws Exception {
        DescribeJobResponseBody self = new DescribeJobResponseBody();
        return TeaModel.build(map, self);
    }

    public DescribeJobResponseBody setDescribeJobResult(DescribeJobResult describeJobResult) {
        this.describeJobResult = describeJobResult;
        return this;
    }
    public DescribeJobResult getDescribeJobResult() {
        return this.describeJobResult;
    }

}
