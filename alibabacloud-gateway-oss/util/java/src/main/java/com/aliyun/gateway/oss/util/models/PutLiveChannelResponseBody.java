// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutLiveChannelResponseBody extends TeaModel {
    /**
     * <p>The container that stores the URL used to play the streams ingested to the LiveChannel.</p>
     */
    @NameInMap("PlayUrls")
    public LiveChannelPlayUrls playUrls;

    /**
     * <p>The container that stores the URL used to ingest streams to the LiveChannel.</p>
     */
    @NameInMap("PublishUrls")
    public LiveChannelPublishUrls publishUrls;

    public static PutLiveChannelResponseBody build(java.util.Map<String, ?> map) throws Exception {
        PutLiveChannelResponseBody self = new PutLiveChannelResponseBody();
        return TeaModel.build(map, self);
    }

    public PutLiveChannelResponseBody setPlayUrls(LiveChannelPlayUrls playUrls) {
        this.playUrls = playUrls;
        return this;
    }
    public LiveChannelPlayUrls getPlayUrls() {
        return this.playUrls;
    }

    public PutLiveChannelResponseBody setPublishUrls(LiveChannelPublishUrls publishUrls) {
        this.publishUrls = publishUrls;
        return this;
    }
    public LiveChannelPublishUrls getPublishUrls() {
        return this.publishUrls;
    }

}
