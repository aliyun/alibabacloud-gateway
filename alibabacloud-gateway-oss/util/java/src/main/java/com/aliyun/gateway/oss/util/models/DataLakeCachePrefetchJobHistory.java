// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DataLakeCachePrefetchJobHistory extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>1727176823</p>
     */
    @NameInMap("EndTime")
    public Long endTime;

    /**
     * <strong>example:</strong>
     * <p>8w643e67a54c4670b5ed321511234567</p>
     */
    @NameInMap("Id")
    public String id;

    /**
     * <strong>example:</strong>
     * <p>345sdf60df1329d88482912343ea74g2</p>
     */
    @NameInMap("JobId")
    public String jobId;

    /**
     * <strong>example:</strong>
     * <p>1727176463</p>
     */
    @NameInMap("StartTime")
    public Long startTime;

    /**
     * <strong>example:</strong>
     * <p>REPLICATION_JOB_SUCCESS</p>
     */
    @NameInMap("Status")
    public String status;

    /**
     * <strong>example:</strong>
     * <p>0</p>
     */
    @NameInMap("SucceedCount")
    public Long succeedCount;

    /**
     * <strong>example:</strong>
     * <p>1</p>
     */
    @NameInMap("TotalCount")
    public Long totalCount;

    public static DataLakeCachePrefetchJobHistory build(java.util.Map<String, ?> map) throws Exception {
        DataLakeCachePrefetchJobHistory self = new DataLakeCachePrefetchJobHistory();
        return TeaModel.build(map, self);
    }

    public DataLakeCachePrefetchJobHistory setEndTime(Long endTime) {
        this.endTime = endTime;
        return this;
    }
    public Long getEndTime() {
        return this.endTime;
    }

    public DataLakeCachePrefetchJobHistory setId(String id) {
        this.id = id;
        return this;
    }
    public String getId() {
        return this.id;
    }

    public DataLakeCachePrefetchJobHistory setJobId(String jobId) {
        this.jobId = jobId;
        return this;
    }
    public String getJobId() {
        return this.jobId;
    }

    public DataLakeCachePrefetchJobHistory setStartTime(Long startTime) {
        this.startTime = startTime;
        return this;
    }
    public Long getStartTime() {
        return this.startTime;
    }

    public DataLakeCachePrefetchJobHistory setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

    public DataLakeCachePrefetchJobHistory setSucceedCount(Long succeedCount) {
        this.succeedCount = succeedCount;
        return this;
    }
    public Long getSucceedCount() {
        return this.succeedCount;
    }

    public DataLakeCachePrefetchJobHistory setTotalCount(Long totalCount) {
        this.totalCount = totalCount;
        return this;
    }
    public Long getTotalCount() {
        return this.totalCount;
    }

}
