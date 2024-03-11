// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class PutBucketWebsiteRequest extends TeaModel {
    /**
     * <p>The container that stores the website configuration.</p>
     */
    @NameInMap("WebsiteConfiguration")
    public WebsiteConfiguration websiteConfiguration;

    public static PutBucketWebsiteRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketWebsiteRequest self = new PutBucketWebsiteRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketWebsiteRequest setWebsiteConfiguration(WebsiteConfiguration websiteConfiguration) {
        this.websiteConfiguration = websiteConfiguration;
        return this;
    }
    public WebsiteConfiguration getWebsiteConfiguration() {
        return this.websiteConfiguration;
    }

}
