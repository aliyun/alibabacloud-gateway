// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class CreateBucketDataRedundancyTransitionResponseBody extends TeaModel {
    @NameInMap("TaskId")
    public String taskId;

    public static CreateBucketDataRedundancyTransitionResponseBody build(java.util.Map<String, ?> map) throws Exception {
        CreateBucketDataRedundancyTransitionResponseBody self = new CreateBucketDataRedundancyTransitionResponseBody();
        return TeaModel.build(map, self);
    }

    public CreateBucketDataRedundancyTransitionResponseBody setTaskId(String taskId) {
        this.taskId = taskId;
        return this;
    }
    public String getTaskId() {
        return this.taskId;
    }

}
