// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetChannelResponseBody extends TeaModel {
    @NameInMap("channel")
    public GetChannelResult channel;

    public static GetChannelResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetChannelResponseBody self = new GetChannelResponseBody();
        return TeaModel.build(map, self);
    }

    public GetChannelResponseBody setChannel(GetChannelResult channel) {
        this.channel = channel;
        return this;
    }
    public GetChannelResult getChannel() {
        return this.channel;
    }

}
