// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListJobResponseBody extends TeaModel {
    @NameInMap("JobList")
    public ListJobResp jobList;

    public static ListJobResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListJobResponseBody self = new ListJobResponseBody();
        return TeaModel.build(map, self);
    }

    public ListJobResponseBody setJobList(ListJobResp jobList) {
        this.jobList = jobList;
        return this;
    }
    public ListJobResp getJobList() {
        return this.jobList;
    }

}
