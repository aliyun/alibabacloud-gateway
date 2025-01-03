// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DataLakeCachePrefetchJob extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>bucket-example</p>
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
     * <p>84dcbfa10wdp488211dc6cb287f4d804</p>
     */
    @NameInMap("HistoryId")
    public String historyId;

    /**
     * <strong>example:</strong>
     * <p>84dcbfa10wdp488211dc6cb287f4d804</p>
     */
    @NameInMap("Id")
    public String id;

    /**
     * <strong>example:</strong>
     * <p>1727164655</p>
     */
    @NameInMap("LastModifyTime")
    public Long lastModifyTime;

    @NameInMap("Rule")
    public DataLakeCachePrefetchJobRule rule;

    /**
     * <strong>example:</strong>
     * <p>REPLICATION_JOB_IDLE</p>
     */
    @NameInMap("Status")
    public String status;

    /**
     * <strong>example:</strong>
     * <p>0</p>
     */
    @NameInMap("Type")
    public Integer type;

    public static DataLakeCachePrefetchJob build(java.util.Map<String, ?> map) throws Exception {
        DataLakeCachePrefetchJob self = new DataLakeCachePrefetchJob();
        return TeaModel.build(map, self);
    }

    public DataLakeCachePrefetchJob setBucket(String bucket) {
        this.bucket = bucket;
        return this;
    }
    public String getBucket() {
        return this.bucket;
    }

    public DataLakeCachePrefetchJob setCreateTime(Long createTime) {
        this.createTime = createTime;
        return this;
    }
    public Long getCreateTime() {
        return this.createTime;
    }

    public DataLakeCachePrefetchJob setHistoryId(String historyId) {
        this.historyId = historyId;
        return this;
    }
    public String getHistoryId() {
        return this.historyId;
    }

    public DataLakeCachePrefetchJob setId(String id) {
        this.id = id;
        return this;
    }
    public String getId() {
        return this.id;
    }

    public DataLakeCachePrefetchJob setLastModifyTime(Long lastModifyTime) {
        this.lastModifyTime = lastModifyTime;
        return this;
    }
    public Long getLastModifyTime() {
        return this.lastModifyTime;
    }

    public DataLakeCachePrefetchJob setRule(DataLakeCachePrefetchJobRule rule) {
        this.rule = rule;
        return this;
    }
    public DataLakeCachePrefetchJobRule getRule() {
        return this.rule;
    }

    public DataLakeCachePrefetchJob setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

    public DataLakeCachePrefetchJob setType(Integer type) {
        this.type = type;
        return this;
    }
    public Integer getType() {
        return this.type;
    }

}
