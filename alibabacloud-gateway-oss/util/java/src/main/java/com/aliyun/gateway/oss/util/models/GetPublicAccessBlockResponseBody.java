// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetPublicAccessBlockResponseBody extends TeaModel {
    @NameInMap("PublicAccessBlockConfiguration")
    public PublicAccessBlockConfiguration publicAccessBlockConfiguration;

    public static GetPublicAccessBlockResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetPublicAccessBlockResponseBody self = new GetPublicAccessBlockResponseBody();
        return TeaModel.build(map, self);
    }

    public GetPublicAccessBlockResponseBody setPublicAccessBlockConfiguration(PublicAccessBlockConfiguration publicAccessBlockConfiguration) {
        this.publicAccessBlockConfiguration = publicAccessBlockConfiguration;
        return this;
    }
    public PublicAccessBlockConfiguration getPublicAccessBlockConfiguration() {
        return this.publicAccessBlockConfiguration;
    }

}
