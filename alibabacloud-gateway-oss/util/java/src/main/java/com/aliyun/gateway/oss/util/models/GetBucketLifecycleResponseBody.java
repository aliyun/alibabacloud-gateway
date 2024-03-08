// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketLifecycleResponseBody extends TeaModel {
    /**
     * <p>The container that stores the lifecycle rules.</p>
     */
    @NameInMap("Rule")
    public java.util.List<LifecycleRule> rule;

    public static GetBucketLifecycleResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketLifecycleResponseBody self = new GetBucketLifecycleResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketLifecycleResponseBody setRule(java.util.List<LifecycleRule> rule) {
        this.rule = rule;
        return this;
    }
    public java.util.List<LifecycleRule> getRule() {
        return this.rule;
    }

}
