// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class GetBucketResponseHeaderResponseBody extends TeaModel {
    @NameInMap("ResponseHeaderConfiguration")
    public ResponseHeaderConfiguration responseHeaderConfiguration;

    public static GetBucketResponseHeaderResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketResponseHeaderResponseBody self = new GetBucketResponseHeaderResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketResponseHeaderResponseBody setResponseHeaderConfiguration(ResponseHeaderConfiguration responseHeaderConfiguration) {
        this.responseHeaderConfiguration = responseHeaderConfiguration;
        return this;
    }
    public ResponseHeaderConfiguration getResponseHeaderConfiguration() {
        return this.responseHeaderConfiguration;
    }

}
