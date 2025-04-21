// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutChannelRequest extends TeaModel {
    /**
     * <p>Container for storing image processing channel configuration</p>
     */
    @NameInMap("Channel")
    public Channel channel;

    public static PutChannelRequest build(java.util.Map<String, ?> map) throws Exception {
        PutChannelRequest self = new PutChannelRequest();
        return TeaModel.build(map, self);
    }

    public PutChannelRequest setChannel(Channel channel) {
        this.channel = channel;
        return this;
    }
    public Channel getChannel() {
        return this.channel;
    }

}
