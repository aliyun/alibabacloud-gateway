// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListJobHistoryResponseBody extends TeaModel {
    @NameInMap("JobHistoryList")
    public ListJobHistoryResp jobHistoryList;

    public static ListJobHistoryResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListJobHistoryResponseBody self = new ListJobHistoryResponseBody();
        return TeaModel.build(map, self);
    }

    public ListJobHistoryResponseBody setJobHistoryList(ListJobHistoryResp jobHistoryList) {
        this.jobHistoryList = jobHistoryList;
        return this;
    }
    public ListJobHistoryResp getJobHistoryList() {
        return this.jobHistoryList;
    }

}
