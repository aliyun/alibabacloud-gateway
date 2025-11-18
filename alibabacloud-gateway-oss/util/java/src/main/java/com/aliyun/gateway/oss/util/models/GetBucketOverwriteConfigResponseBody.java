// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketOverwriteConfigResponseBody extends TeaModel {
    @NameInMap("OverwriteConfiguration")
    public OverwriteConfiguration overwriteConfiguration;

    public static GetBucketOverwriteConfigResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketOverwriteConfigResponseBody self = new GetBucketOverwriteConfigResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketOverwriteConfigResponseBody setOverwriteConfiguration(OverwriteConfiguration overwriteConfiguration) {
        this.overwriteConfiguration = overwriteConfiguration;
        return this;
    }
    public OverwriteConfiguration getOverwriteConfiguration() {
        return this.overwriteConfiguration;
    }

}
