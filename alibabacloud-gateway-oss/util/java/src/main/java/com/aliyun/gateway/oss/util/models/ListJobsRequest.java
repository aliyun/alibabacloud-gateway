// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListJobsRequest extends TeaModel {
    @NameInMap("batchJobStatuses")
    public String batchJobStatuses;

    @NameInMap("continuation-token")
    public String continuationToken;

    @NameInMap("max-keys")
    public Integer maxKeys;

    public static ListJobsRequest build(java.util.Map<String, ?> map) throws Exception {
        ListJobsRequest self = new ListJobsRequest();
        return TeaModel.build(map, self);
    }

    public ListJobsRequest setBatchJobStatuses(String batchJobStatuses) {
        this.batchJobStatuses = batchJobStatuses;
        return this;
    }
    public String getBatchJobStatuses() {
        return this.batchJobStatuses;
    }

    public ListJobsRequest setContinuationToken(String continuationToken) {
        this.continuationToken = continuationToken;
        return this;
    }
    public String getContinuationToken() {
        return this.continuationToken;
    }

    public ListJobsRequest setMaxKeys(Integer maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Integer getMaxKeys() {
        return this.maxKeys;
    }

}
