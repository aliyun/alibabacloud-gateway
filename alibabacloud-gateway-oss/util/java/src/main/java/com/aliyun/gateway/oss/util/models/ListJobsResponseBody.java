// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListJobsResponseBody extends TeaModel {
    @NameInMap("ListJobsResult")
    public BatchOperationListJobsResult listJobsResult;

    public static ListJobsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListJobsResponseBody self = new ListJobsResponseBody();
        return TeaModel.build(map, self);
    }

    public ListJobsResponseBody setListJobsResult(BatchOperationListJobsResult listJobsResult) {
        this.listJobsResult = listJobsResult;
        return this;
    }
    public BatchOperationListJobsResult getListJobsResult() {
        return this.listJobsResult;
    }

}
