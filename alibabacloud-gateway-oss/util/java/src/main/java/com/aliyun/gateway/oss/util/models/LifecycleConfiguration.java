// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class LifecycleConfiguration extends TeaModel {
    @NameInMap("Rule")
    public java.util.List<LifecycleRule> rule;

    public static LifecycleConfiguration build(java.util.Map<String, ?> map) throws Exception {
        LifecycleConfiguration self = new LifecycleConfiguration();
        return TeaModel.build(map, self);
    }

    public LifecycleConfiguration setRule(java.util.List<LifecycleRule> rule) {
        this.rule = rule;
        return this;
    }
    public java.util.List<LifecycleRule> getRule() {
        return this.rule;
    }

}
