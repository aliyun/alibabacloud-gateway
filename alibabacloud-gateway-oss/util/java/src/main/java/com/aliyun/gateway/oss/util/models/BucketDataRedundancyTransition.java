// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class BucketDataRedundancyTransition extends TeaModel {
    @NameInMap("Bucket")
    public String bucket;

    @NameInMap("CreateTime")
    public String createTime;

    @NameInMap("EndTime")
    public String endTime;

    @NameInMap("EstimatedRemainingTime")
    public String estimatedRemainingTime;

    @NameInMap("ProcessPercentage")
    public Integer processPercentage;

    @NameInMap("StartTime")
    public String startTime;

    @NameInMap("Status")
    public String status;

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
