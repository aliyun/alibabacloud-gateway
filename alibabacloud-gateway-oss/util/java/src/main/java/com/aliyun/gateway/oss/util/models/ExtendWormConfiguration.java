// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ExtendWormConfiguration extends TeaModel {
    @NameInMap("RetentionPeriodInDays")
    public Integer retentionPeriodInDays;

    public static ExtendWormConfiguration build(java.util.Map<String, ?> map) throws Exception {
        ExtendWormConfiguration self = new ExtendWormConfiguration();
        return TeaModel.build(map, self);
    }

    public ExtendWormConfiguration setRetentionPeriodInDays(Integer retentionPeriodInDays) {
        this.retentionPeriodInDays = retentionPeriodInDays;
        return this;
    }
    public Integer getRetentionPeriodInDays() {
        return this.retentionPeriodInDays;
    }

}
