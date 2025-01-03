// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetProcessConfigurationResponseBody extends TeaModel {
    @NameInMap("BucketProcessConfiguration")
    public GetBucketProcessConfiguration bucketProcessConfiguration;

    public static GetProcessConfigurationResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetProcessConfigurationResponseBody self = new GetProcessConfigurationResponseBody();
        return TeaModel.build(map, self);
    }

    public GetProcessConfigurationResponseBody setBucketProcessConfiguration(GetBucketProcessConfiguration bucketProcessConfiguration) {
        this.bucketProcessConfiguration = bucketProcessConfiguration;
        return this;
    }
    public GetBucketProcessConfiguration getBucketProcessConfiguration() {
        return this.bucketProcessConfiguration;
    }

}
