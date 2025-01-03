// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BucketDataRedundancyTransition extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>example-bucket</p>
     */
    @NameInMap("Bucket")
    public String bucket;

    /**
     * <strong>example:</strong>
     * <p>2023-11-17T09:14:39.000Z</p>
     */
    @NameInMap("CreateTime")
    public String createTime;

    /**
     * <strong>example:</strong>
     * <p>2023-11-17T09:14:39.000Z</p>
     */
    @NameInMap("EndTime")
    public String endTime;

    /**
     * <strong>example:</strong>
     * <p>10</p>
     */
    @NameInMap("EstimatedRemainingTime")
    public String estimatedRemainingTime;

    /**
     * <strong>example:</strong>
     * <p>88</p>
     */
    @NameInMap("ProcessPercentage")
    public Integer processPercentage;

    /**
     * <strong>example:</strong>
     * <p>2023-11-17T09:14:39.000Z</p>
     */
    @NameInMap("StartTime")
    public String startTime;

    /**
     * <strong>example:</strong>
     * <p>Queueing</p>
     */
    @NameInMap("Status")
    public String status;

    /**
     * <strong>example:</strong>
     * <p>751f5243f8ac4ae89f34726534d1****</p>
     */
    @NameInMap("TaskId")
    public String taskId;

    public static BucketDataRedundancyTransition build(java.util.Map<String, ?> map) throws Exception {
        BucketDataRedundancyTransition self = new BucketDataRedundancyTransition();
        return TeaModel.build(map, self);
    }

    public BucketDataRedundancyTransition setBucket(String bucket) {
        this.bucket = bucket;
        return this;
    }
    public String getBucket() {
        return this.bucket;
    }

    public BucketDataRedundancyTransition setCreateTime(String createTime) {
        this.createTime = createTime;
        return this;
    }
    public String getCreateTime() {
        return this.createTime;
    }

    public BucketDataRedundancyTransition setEndTime(String endTime) {
        this.endTime = endTime;
        return this;
    }
    public String getEndTime() {
        return this.endTime;
    }

    public BucketDataRedundancyTransition setEstimatedRemainingTime(String estimatedRemainingTime) {
        this.estimatedRemainingTime = estimatedRemainingTime;
        return this;
    }
    public String getEstimatedRemainingTime() {
        return this.estimatedRemainingTime;
    }

    public BucketDataRedundancyTransition setProcessPercentage(Integer processPercentage) {
        this.processPercentage = processPercentage;
        return this;
    }
    public Integer getProcessPercentage() {
        return this.processPercentage;
    }

    public BucketDataRedundancyTransition setStartTime(String startTime) {
        this.startTime = startTime;
        return this;
    }
    public String getStartTime() {
        return this.startTime;
    }

    public BucketDataRedundancyTransition setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

    public BucketDataRedundancyTransition setTaskId(String taskId) {
        this.taskId = taskId;
        return this;
    }
    public String getTaskId() {
        return this.taskId;
    }

}
