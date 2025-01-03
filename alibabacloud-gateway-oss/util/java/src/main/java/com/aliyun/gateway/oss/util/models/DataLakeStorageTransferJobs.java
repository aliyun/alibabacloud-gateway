// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DataLakeStorageTransferJobs extends TeaModel {
    @NameInMap("DataLakeStorageTransferJob")
    public java.util.List<DataLakeStorageTransferJob> dataLakeStorageTransferJob;

    /**
     * <strong>example:</strong>
     * <p>test-bucket</p>
     */
    @NameInMap("NextMarkerBucket")
    public String nextMarkerBucket;

    /**
     * <strong>example:</strong>
     * <p>abcdef03370a4b32ac6bc69eb1123456</p>
     */
    @NameInMap("NextMarkerJobId")
    public String nextMarkerJobId;

    /**
     * <strong>example:</strong>
     * <p>false</p>
     */
    @NameInMap("Truncated")
    public String truncated;

    public static DataLakeStorageTransferJobs build(java.util.Map<String, ?> map) throws Exception {
        DataLakeStorageTransferJobs self = new DataLakeStorageTransferJobs();
        return TeaModel.build(map, self);
    }

    public DataLakeStorageTransferJobs setDataLakeStorageTransferJob(java.util.List<DataLakeStorageTransferJob> dataLakeStorageTransferJob) {
        this.dataLakeStorageTransferJob = dataLakeStorageTransferJob;
        return this;
    }
    public java.util.List<DataLakeStorageTransferJob> getDataLakeStorageTransferJob() {
        return this.dataLakeStorageTransferJob;
    }

    public DataLakeStorageTransferJobs setNextMarkerBucket(String nextMarkerBucket) {
        this.nextMarkerBucket = nextMarkerBucket;
        return this;
    }
    public String getNextMarkerBucket() {
        return this.nextMarkerBucket;
    }

    public DataLakeStorageTransferJobs setNextMarkerJobId(String nextMarkerJobId) {
        this.nextMarkerJobId = nextMarkerJobId;
        return this;
    }
    public String getNextMarkerJobId() {
        return this.nextMarkerJobId;
    }

    public DataLakeStorageTransferJobs setTruncated(String truncated) {
        this.truncated = truncated;
        return this;
    }
    public String getTruncated() {
        return this.truncated;
    }

}
