// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListDataLakeCachePrefetchJobResponseBody extends TeaModel {
    @NameInMap("DataLakeCachePrefetchJobs")
    public DataLakeCachePrefetchJobs dataLakeCachePrefetchJobs;

    public static ListDataLakeCachePrefetchJobResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListDataLakeCachePrefetchJobResponseBody self = new ListDataLakeCachePrefetchJobResponseBody();
        return TeaModel.build(map, self);
    }

    public ListDataLakeCachePrefetchJobResponseBody setDataLakeCachePrefetchJobs(DataLakeCachePrefetchJobs dataLakeCachePrefetchJobs) {
        this.dataLakeCachePrefetchJobs = dataLakeCachePrefetchJobs;
        return this;
    }
    public DataLakeCachePrefetchJobs getDataLakeCachePrefetchJobs() {
        return this.dataLakeCachePrefetchJobs;
    }

    public static class DataLakeCachePrefetchJobs extends TeaModel {
        @NameInMap("DataLakeCachePrefetchJob")
        public DataLakeCachePrefetchJob dataLakeCachePrefetchJob;

        @NameInMap("NextMarkerBucket")
        public String nextMarkerBucket;

        @NameInMap("NextMarkerJobId")
        public String nextMarkerJobId;

        @NameInMap("Truncated")
        public Boolean truncated;

        public static DataLakeCachePrefetchJobs build(java.util.Map<String, ?> map) throws Exception {
            DataLakeCachePrefetchJobs self = new DataLakeCachePrefetchJobs();
            return TeaModel.build(map, self);
        }

        public DataLakeCachePrefetchJobs setDataLakeCachePrefetchJob(DataLakeCachePrefetchJob dataLakeCachePrefetchJob) {
            this.dataLakeCachePrefetchJob = dataLakeCachePrefetchJob;
            return this;
        }
        public DataLakeCachePrefetchJob getDataLakeCachePrefetchJob() {
            return this.dataLakeCachePrefetchJob;
        }

        public DataLakeCachePrefetchJobs setNextMarkerBucket(String nextMarkerBucket) {
            this.nextMarkerBucket = nextMarkerBucket;
            return this;
        }
        public String getNextMarkerBucket() {
            return this.nextMarkerBucket;
        }

        public DataLakeCachePrefetchJobs setNextMarkerJobId(String nextMarkerJobId) {
            this.nextMarkerJobId = nextMarkerJobId;
            return this;
        }
        public String getNextMarkerJobId() {
            return this.nextMarkerJobId;
        }

        public DataLakeCachePrefetchJobs setTruncated(Boolean truncated) {
            this.truncated = truncated;
            return this;
        }
        public Boolean getTruncated() {
            return this.truncated;
        }

    }

}
