// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class PutLiveChannelRequest extends TeaModel {
    /**
     * <p>The container that stores the configurations of the LiveChannel.</p>
     */
    @NameInMap("LiveChannelConfiguration")
    public LiveChannelConfiguration liveChannelConfiguration;

    public static PutLiveChannelRequest build(java.util.Map<String, ?> map) throws Exception {
        PutLiveChannelRequest self = new PutLiveChannelRequest();
        return TeaModel.build(map, self);
    }

    public PutLiveChannelRequest setLiveChannelConfiguration(LiveChannelConfiguration liveChannelConfiguration) {
        this.liveChannelConfiguration = liveChannelConfiguration;
        return this;
    }
    public LiveChannelConfiguration getLiveChannelConfiguration() {
        return this.liveChannelConfiguration;
    }

}
