// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class AsyncFetchTaskResult extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>abc</p>
     */
    @NameInMap("TaskId")
    public String taskId;

    public static AsyncFetchTaskResult build(java.util.Map<String, ?> map) throws Exception {
        AsyncFetchTaskResult self = new AsyncFetchTaskResult();
        return TeaModel.build(map, self);
    }

    public AsyncFetchTaskResult setTaskId(String taskId) {
        this.taskId = taskId;
        return this;
    }
    public String getTaskId() {
        return this.taskId;
    }

}
