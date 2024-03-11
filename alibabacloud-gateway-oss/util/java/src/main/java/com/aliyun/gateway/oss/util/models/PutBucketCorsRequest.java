// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class PutBucketCorsRequest extends TeaModel {
    /**
     * <p>The container that stores CORS rules.</p>
     * <br>
     * <p>You can configure up to 10 CORS rules for a bucket. The XML message body in a request can be up to 16 KB in size.</p>
     */
    @NameInMap("CORSConfiguration")
    public CORSConfiguration CORSConfiguration;

    public static PutBucketCorsRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketCorsRequest self = new PutBucketCorsRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketCorsRequest setCORSConfiguration(CORSConfiguration CORSConfiguration) {
        this.CORSConfiguration = CORSConfiguration;
        return this;
    }
    public CORSConfiguration getCORSConfiguration() {
        return this.CORSConfiguration;
    }

}
