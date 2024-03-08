// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class LiveChannelPublishUrls extends TeaModel {
    @NameInMap("Url")
    public String url;

    public static LiveChannelPublishUrls build(java.util.Map<String, ?> map) throws Exception {
        LiveChannelPublishUrls self = new LiveChannelPublishUrls();
        return TeaModel.build(map, self);
    }

    public LiveChannelPublishUrls setUrl(String url) {
        this.url = url;
        return this;
    }
    public String getUrl() {
        return this.url;
    }

}
