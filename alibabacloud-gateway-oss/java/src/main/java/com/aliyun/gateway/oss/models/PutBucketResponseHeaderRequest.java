// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class PutBucketResponseHeaderRequest extends TeaModel {
    @NameInMap("ResponseHeaderConfiguration")
    public ResponseHeaderConfiguration responseHeaderConfiguration;

    public static PutBucketResponseHeaderRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketResponseHeaderRequest self = new PutBucketResponseHeaderRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketResponseHeaderRequest setResponseHeaderConfiguration(ResponseHeaderConfiguration responseHeaderConfiguration) {
        this.responseHeaderConfiguration = responseHeaderConfiguration;
        return this;
    }
    public ResponseHeaderConfiguration getResponseHeaderConfiguration() {
        return this.responseHeaderConfiguration;
    }

}
