// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutLiveChannelRequest extends TeaModel {
    /**
     * <p>The container that stores the configurations of the LiveChannel.</p>
     */
    @NameInMap("body")
    public LiveChannelConfiguration body;

    public static PutLiveChannelRequest build(java.util.Map<String, ?> map) throws Exception {
        PutLiveChannelRequest self = new PutLiveChannelRequest();
        return TeaModel.build(map, self);
    }

    public PutLiveChannelRequest setBody(LiveChannelConfiguration body) {
        this.body = body;
        return this;
    }
    public LiveChannelConfiguration getBody() {
        return this.body;
    }

}
