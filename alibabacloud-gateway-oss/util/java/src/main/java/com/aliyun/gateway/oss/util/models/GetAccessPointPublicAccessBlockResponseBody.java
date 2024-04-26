// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetAccessPointPublicAccessBlockResponseBody extends TeaModel {
    @NameInMap("PublicAccessBlockConfiguration")
    public PublicAccessBlockConfiguration publicAccessBlockConfiguration;

    public static GetAccessPointPublicAccessBlockResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetAccessPointPublicAccessBlockResponseBody self = new GetAccessPointPublicAccessBlockResponseBody();
        return TeaModel.build(map, self);
    }

    public GetAccessPointPublicAccessBlockResponseBody setPublicAccessBlockConfiguration(PublicAccessBlockConfiguration publicAccessBlockConfiguration) {
        this.publicAccessBlockConfiguration = publicAccessBlockConfiguration;
        return this;
    }
    public PublicAccessBlockConfiguration getPublicAccessBlockConfiguration() {
        return this.publicAccessBlockConfiguration;
    }

}
