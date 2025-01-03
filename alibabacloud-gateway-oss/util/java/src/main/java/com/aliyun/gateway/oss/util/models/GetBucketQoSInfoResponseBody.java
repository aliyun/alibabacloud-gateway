// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketQoSInfoResponseBody extends TeaModel {
    @NameInMap("QoSConfiguration")
    public BucketQoSConfiguration qoSConfiguration;

    public static GetBucketQoSInfoResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketQoSInfoResponseBody self = new GetBucketQoSInfoResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketQoSInfoResponseBody setQoSConfiguration(BucketQoSConfiguration qoSConfiguration) {
        this.qoSConfiguration = qoSConfiguration;
        return this;
    }
    public BucketQoSConfiguration getQoSConfiguration() {
        return this.qoSConfiguration;
    }

}
