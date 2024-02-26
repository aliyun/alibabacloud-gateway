// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class PutBucketHttpsConfigRequest extends TeaModel {
    /**
     * <p>The container that stores HTTPS configurations.</p>
     */
    @NameInMap("HttpsConfiguration")
    public HttpsConfiguration httpsConfiguration;

    public static PutBucketHttpsConfigRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketHttpsConfigRequest self = new PutBucketHttpsConfigRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketHttpsConfigRequest setHttpsConfiguration(HttpsConfiguration httpsConfiguration) {
        this.httpsConfiguration = httpsConfiguration;
        return this;
    }
    public HttpsConfiguration getHttpsConfiguration() {
        return this.httpsConfiguration;
    }

}
