// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DataLakeStorageTransferJob extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>bucket1</p>
     */
    @NameInMap("Bucket")
    public String bucket;

    /**
     * <strong>example:</strong>
     * <p>1727162332</p>
     */
    @NameInMap("CreateTime")
    public Long createTime;

    /**
     * <strong>example:</strong>
     * <p>abcdef87e3c04af2af290ec52d123456</p>
     */
    @NameInMap("HistoryId")
    public String historyId;

    /**
     * <strong>example:</strong>
     * <p>abcdef87e3c04af2af290ec52d123456</p>
     */
    @NameInMap("Id")
    public String id;

    /**
     * <strong>example:</strong>
     * <p>1727164655</p>
     */
    @NameInMap("LastModifyTime")
    public Long lastModifyTime;

    @NameInMap("ProgressInfo")
    public ProgressInfo progressInfo;

    @NameInMap("Rule")
    public DataLakeStorageTransferJobRule rule;

    /**
     * <strong>example:</strong>
     * <p>REPLICATION_JOB_IDLE</p>
     */
    @NameInMap("Status")
    public String status;

    /**
     * <strong>example:</strong>
     * <p>1</p>
     */
    @NameInMap("Type")
    public Integer type;

    public static DataLakeStorageTransferJob build(java.util.Map<String, ?> map) throws Exception {
        DataLakeStorageTransferJob self = new DataLakeStorageTransferJob();
        return TeaModel.build(map, self);
    }

    public DataLakeStorageTransferJob setBucket(String bucket) {
        this.bucket = bucket;
        return this;
    }
    public String getBucket() {
        return this.bucket;
    }

    public DataLakeStorageTransferJob setCreateTime(Long createTime) {
        this.createTime = createTime;
        return this;
    }
    public Long getCreateTime() {
        return this.createTime;
    }

    public DataLakeStorageTransferJob setHistoryId(String historyId) {
        this.historyId = historyId;
        return this;
    }
    public String getHistoryId() {
        return this.historyId;
    }

    public DataLakeStorageTransferJob setId(String id) {
        this.id = id;
        return this;
    }
    public String getId() {
        return this.id;
    }

    public DataLakeStorageTransferJob setLastModifyTime(Long lastModifyTime) {
        this.lastModifyTime = lastModifyTime;
        return this;
    }
    public Long getLastModifyTime() {
        return this.lastModifyTime;
    }

    public DataLakeStorageTransferJob setProgressInfo(ProgressInfo progressInfo) {
        this.progressInfo = progressInfo;
        return this;
    }
    public ProgressInfo getProgressInfo() {
        return this.progressInfo;
    }

    public DataLakeStorageTransferJob setRule(DataLakeStorageTransferJobRule rule) {
        this.rule = rule;
        return this;
    }
    public DataLakeStorageTransferJobRule getRule() {
        return this.rule;
    }

    public DataLakeStorageTransferJob setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

    public DataLakeStorageTransferJob setType(Integer type) {
        this.type = type;
        return this;
    }
    public Integer getType() {
        return this.type;
    }

    public static class ProgressInfo extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>50</p>
         */
        @NameInMap("Percent")
        public Long percent;

        public static ProgressInfo build(java.util.Map<String, ?> map) throws Exception {
            ProgressInfo self = new ProgressInfo();
            return TeaModel.build(map, self);
        }

        public ProgressInfo setPercent(Long percent) {
            this.percent = percent;
            return this;
        }
        public Long getPercent() {
            return this.percent;
        }

    }

}
