// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DataLakeStorageTransferJobHistory extends TeaModel {
    @NameInMap("DetailInfo")
    public DetailInfo detailInfo;

    /**
     * <strong>example:</strong>
     * <p>1731830653</p>
     */
    @NameInMap("EndTime")
    public Long endTime;

    /**
     * <strong>example:</strong>
     * <p>abcdef03370a4b32ac6bc69eb1123456</p>
     */
    @NameInMap("Id")
    public String id;

    /**
     * <strong>example:</strong>
     * <p>43452</p>
     */
    @NameInMap("JobId")
    public String jobId;

    /**
     * <strong>example:</strong>
     * <p>1731830653</p>
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
     * <p>12345</p>
     */
    @NameInMap("SucceedCount")
    public Long succeedCount;

    /**
     * <strong>example:</strong>
     * <p>12</p>
     */
    @NameInMap("TotalCount")
    public Long totalCount;

    public static DataLakeStorageTransferJobHistory build(java.util.Map<String, ?> map) throws Exception {
        DataLakeStorageTransferJobHistory self = new DataLakeStorageTransferJobHistory();
        return TeaModel.build(map, self);
    }

    public DataLakeStorageTransferJobHistory setDetailInfo(DetailInfo detailInfo) {
        this.detailInfo = detailInfo;
        return this;
    }
    public DetailInfo getDetailInfo() {
        return this.detailInfo;
    }

    public DataLakeStorageTransferJobHistory setEndTime(Long endTime) {
        this.endTime = endTime;
        return this;
    }
    public Long getEndTime() {
        return this.endTime;
    }

    public DataLakeStorageTransferJobHistory setId(String id) {
        this.id = id;
        return this;
    }
    public String getId() {
        return this.id;
    }

    public DataLakeStorageTransferJobHistory setJobId(String jobId) {
        this.jobId = jobId;
        return this;
    }
    public String getJobId() {
        return this.jobId;
    }

    public DataLakeStorageTransferJobHistory setStartTime(Long startTime) {
        this.startTime = startTime;
        return this;
    }
    public Long getStartTime() {
        return this.startTime;
    }

    public DataLakeStorageTransferJobHistory setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

    public DataLakeStorageTransferJobHistory setSucceedCount(Long succeedCount) {
        this.succeedCount = succeedCount;
        return this;
    }
    public Long getSucceedCount() {
        return this.succeedCount;
    }

    public DataLakeStorageTransferJobHistory setTotalCount(Long totalCount) {
        this.totalCount = totalCount;
        return this;
    }
    public Long getTotalCount() {
        return this.totalCount;
    }

    public static class DetailInfo extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>NotSupported</p>
         */
        @NameInMap("ErrorMsg")
        public String errorMsg;

        /**
         * <strong>example:</strong>
         * <p>0</p>
         */
        @NameInMap("HDFSFailedCount")
        public Long HDFSFailedCount;

        /**
         * <strong>example:</strong>
         * <p>abc/</p>
         */
        @NameInMap("HDFSTransferDataDir")
        public String HDFSTransferDataDir;

        /**
         * <strong>example:</strong>
         * <p>hdfs-error/</p>
         */
        @NameInMap("HDFSTransferErrInfoDir")
        public String HDFSTransferErrInfoDir;

        /**
         * <strong>example:</strong>
         * <p>meta-dta/</p>
         */
        @NameInMap("HDFSTransferImportMetaDir")
        public String HDFSTransferImportMetaDir;

        /**
         * <strong>example:</strong>
         * <p>abcdef03370a4b32ac6bc69eb1123456</p>
         */
        @NameInMap("HDFSTransferJobId")
        public String HDFSTransferJobId;

        /**
         * <strong>example:</strong>
         * <p>hdfs-logs/</p>
         */
        @NameInMap("LogBaseDir")
        public String logBaseDir;

        /**
         * <strong>example:</strong>
         * <p>verify-error/</p>
         */
        @NameInMap("VerifyErrInfoDir")
        public String verifyErrInfoDir;

        /**
         * <strong>example:</strong>
         * <p>VERIFY_SUCCEED</p>
         */
        @NameInMap("VerifyStatus")
        public String verifyStatus;

        /**
         * <strong>example:</strong>
         * <p>123456</p>
         */
        @NameInMap("VerifyTotalCount")
        public Long verifyTotalCount;

        public static DetailInfo build(java.util.Map<String, ?> map) throws Exception {
            DetailInfo self = new DetailInfo();
            return TeaModel.build(map, self);
        }

        public DetailInfo setErrorMsg(String errorMsg) {
            this.errorMsg = errorMsg;
            return this;
        }
        public String getErrorMsg() {
            return this.errorMsg;
        }

        public DetailInfo setHDFSFailedCount(Long HDFSFailedCount) {
            this.HDFSFailedCount = HDFSFailedCount;
            return this;
        }
        public Long getHDFSFailedCount() {
            return this.HDFSFailedCount;
        }

        public DetailInfo setHDFSTransferDataDir(String HDFSTransferDataDir) {
            this.HDFSTransferDataDir = HDFSTransferDataDir;
            return this;
        }
        public String getHDFSTransferDataDir() {
            return this.HDFSTransferDataDir;
        }

        public DetailInfo setHDFSTransferErrInfoDir(String HDFSTransferErrInfoDir) {
            this.HDFSTransferErrInfoDir = HDFSTransferErrInfoDir;
            return this;
        }
        public String getHDFSTransferErrInfoDir() {
            return this.HDFSTransferErrInfoDir;
        }

        public DetailInfo setHDFSTransferImportMetaDir(String HDFSTransferImportMetaDir) {
            this.HDFSTransferImportMetaDir = HDFSTransferImportMetaDir;
            return this;
        }
        public String getHDFSTransferImportMetaDir() {
            return this.HDFSTransferImportMetaDir;
        }

        public DetailInfo setHDFSTransferJobId(String HDFSTransferJobId) {
            this.HDFSTransferJobId = HDFSTransferJobId;
            return this;
        }
        public String getHDFSTransferJobId() {
            return this.HDFSTransferJobId;
        }

        public DetailInfo setLogBaseDir(String logBaseDir) {
            this.logBaseDir = logBaseDir;
            return this;
        }
        public String getLogBaseDir() {
            return this.logBaseDir;
        }

        public DetailInfo setVerifyErrInfoDir(String verifyErrInfoDir) {
            this.verifyErrInfoDir = verifyErrInfoDir;
            return this;
        }
        public String getVerifyErrInfoDir() {
            return this.verifyErrInfoDir;
        }

        public DetailInfo setVerifyStatus(String verifyStatus) {
            this.verifyStatus = verifyStatus;
            return this;
        }
        public String getVerifyStatus() {
            return this.verifyStatus;
        }

        public DetailInfo setVerifyTotalCount(Long verifyTotalCount) {
            this.verifyTotalCount = verifyTotalCount;
            return this;
        }
        public Long getVerifyTotalCount() {
            return this.verifyTotalCount;
        }

    }

}
