// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class AsyncFetchTaskInfo extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>FileNotFound</p>
     */
    @NameInMap("ErrorMsg")
    public String errorMsg;

    /**
     * <strong>example:</strong>
     * <p>Success</p>
     */
    @NameInMap("State")
    public String state;

    /**
     * <strong>example:</strong>
     * <p>abc</p>
     */
    @NameInMap("TaskId")
    public String taskId;

    @NameInMap("TaskInfo")
    public AsyncFetchTaskConfiguration taskInfo;

    public static AsyncFetchTaskInfo build(java.util.Map<String, ?> map) throws Exception {
        AsyncFetchTaskInfo self = new AsyncFetchTaskInfo();
        return TeaModel.build(map, self);
    }

    public AsyncFetchTaskInfo setErrorMsg(String errorMsg) {
        this.errorMsg = errorMsg;
        return this;
    }
    public String getErrorMsg() {
        return this.errorMsg;
    }

    public AsyncFetchTaskInfo setState(String state) {
        this.state = state;
        return this;
    }
    public String getState() {
        return this.state;
    }

    public AsyncFetchTaskInfo setTaskId(String taskId) {
        this.taskId = taskId;
        return this;
    }
    public String getTaskId() {
        return this.taskId;
    }

    public AsyncFetchTaskInfo setTaskInfo(AsyncFetchTaskConfiguration taskInfo) {
        this.taskInfo = taskInfo;
        return this;
    }
    public AsyncFetchTaskConfiguration getTaskInfo() {
        return this.taskInfo;
    }

}
