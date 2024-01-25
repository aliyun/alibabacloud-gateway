// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketWebsiteRequest extends TeaModel {
    /**
     * <p>The information about the root node.</p>
     */
    @NameInMap("body")
    public WebsiteConfiguration body;

    public static PutBucketWebsiteRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketWebsiteRequest self = new PutBucketWebsiteRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketWebsiteRequest setBody(WebsiteConfiguration body) {
        this.body = body;
        return this;
    }
    public WebsiteConfiguration getBody() {
        return this.body;
    }

}
