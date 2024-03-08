// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class GetBucketWebsiteResponseBody extends TeaModel {
    /**
     * <p>The containers of the website configuration.</p>
     */
    @NameInMap("WebsiteConfiguration")
    public WebsiteConfiguration websiteConfiguration;

    public static GetBucketWebsiteResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketWebsiteResponseBody self = new GetBucketWebsiteResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketWebsiteResponseBody setWebsiteConfiguration(WebsiteConfiguration websiteConfiguration) {
        this.websiteConfiguration = websiteConfiguration;
        return this;
    }
    public WebsiteConfiguration getWebsiteConfiguration() {
        return this.websiteConfiguration;
    }

}
