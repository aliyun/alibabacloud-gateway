// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class GetBucketRefererResponseBody extends TeaModel {
    /**
     * <p>The container that stores the hotlink protection configurations.</p>
     */
    @NameInMap("RefererConfiguration")
    public RefererConfiguration refererConfiguration;

    public static GetBucketRefererResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketRefererResponseBody self = new GetBucketRefererResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketRefererResponseBody setRefererConfiguration(RefererConfiguration refererConfiguration) {
        this.refererConfiguration = refererConfiguration;
        return this;
    }
    public RefererConfiguration getRefererConfiguration() {
        return this.refererConfiguration;
    }

}
