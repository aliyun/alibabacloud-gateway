// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketOverwriteConfigRequest extends TeaModel {
    @NameInMap("OverwriteConfiguration")
    public OverwriteConfiguration overwriteConfiguration;

    public static PutBucketOverwriteConfigRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketOverwriteConfigRequest self = new PutBucketOverwriteConfigRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketOverwriteConfigRequest setOverwriteConfiguration(OverwriteConfiguration overwriteConfiguration) {
        this.overwriteConfiguration = overwriteConfiguration;
        return this;
    }
    public OverwriteConfiguration getOverwriteConfiguration() {
        return this.overwriteConfiguration;
    }

}
