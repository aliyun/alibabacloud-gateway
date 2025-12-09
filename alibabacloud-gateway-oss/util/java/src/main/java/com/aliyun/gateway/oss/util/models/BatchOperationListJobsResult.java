// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BatchOperationListJobsResult extends TeaModel {
    @NameInMap("Jobs")
    public BatchOperationListJobs jobs;

    /**
     * <strong>example:</strong>
     * <p>/</p>
     */
    @NameInMap("NextToken")
    public String nextToken;

    public static BatchOperationListJobsResult build(java.util.Map<String, ?> map) throws Exception {
        BatchOperationListJobsResult self = new BatchOperationListJobsResult();
        return TeaModel.build(map, self);
    }

    public BatchOperationListJobsResult setJobs(BatchOperationListJobs jobs) {
        this.jobs = jobs;
        return this;
    }
    public BatchOperationListJobs getJobs() {
        return this.jobs;
    }

    public BatchOperationListJobsResult setNextToken(String nextToken) {
        this.nextToken = nextToken;
        return this;
    }
    public String getNextToken() {
        return this.nextToken;
    }

}
