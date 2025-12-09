// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BatchOperationJobProgressSummary extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>3600</p>
     */
    @NameInMap("ElapsedTimeInActiveSeconds")
    public Long elapsedTimeInActiveSeconds;

    /**
     * <strong>example:</strong>
     * <p>0</p>
     */
    @NameInMap("NumberOfTasksFailed")
    public Long numberOfTasksFailed;

    /**
     * <strong>example:</strong>
     * <p>100</p>
     */
    @NameInMap("NumberOfTasksSucceeded")
    public Long numberOfTasksSucceeded;

    /**
     * <strong>example:</strong>
     * <p>100</p>
     */
    @NameInMap("TotalNumberOfTasks")
    public Long totalNumberOfTasks;

    public static BatchOperationJobProgressSummary build(java.util.Map<String, ?> map) throws Exception {
        BatchOperationJobProgressSummary self = new BatchOperationJobProgressSummary();
        return TeaModel.build(map, self);
    }

    public BatchOperationJobProgressSummary setElapsedTimeInActiveSeconds(Long elapsedTimeInActiveSeconds) {
        this.elapsedTimeInActiveSeconds = elapsedTimeInActiveSeconds;
        return this;
    }
    public Long getElapsedTimeInActiveSeconds() {
        return this.elapsedTimeInActiveSeconds;
    }

    public BatchOperationJobProgressSummary setNumberOfTasksFailed(Long numberOfTasksFailed) {
        this.numberOfTasksFailed = numberOfTasksFailed;
        return this;
    }
    public Long getNumberOfTasksFailed() {
        return this.numberOfTasksFailed;
    }

    public BatchOperationJobProgressSummary setNumberOfTasksSucceeded(Long numberOfTasksSucceeded) {
        this.numberOfTasksSucceeded = numberOfTasksSucceeded;
        return this;
    }
    public Long getNumberOfTasksSucceeded() {
        return this.numberOfTasksSucceeded;
    }

    public BatchOperationJobProgressSummary setTotalNumberOfTasks(Long totalNumberOfTasks) {
        this.totalNumberOfTasks = totalNumberOfTasks;
        return this;
    }
    public Long getTotalNumberOfTasks() {
        return this.totalNumberOfTasks;
    }

}
