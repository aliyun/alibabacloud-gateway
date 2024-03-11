// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class GetBucketHttpsConfigResponseBody extends TeaModel {
    /**
     * <p>The container that stores HTTPS configurations.</p>
     */
    @NameInMap("HttpsConfiguration")
    public HttpsConfiguration httpsConfiguration;

    public static GetBucketHttpsConfigResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketHttpsConfigResponseBody self = new GetBucketHttpsConfigResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketHttpsConfigResponseBody setHttpsConfiguration(HttpsConfiguration httpsConfiguration) {
        this.httpsConfiguration = httpsConfiguration;
        return this;
    }
    public HttpsConfiguration getHttpsConfiguration() {
        return this.httpsConfiguration;
    }

}
