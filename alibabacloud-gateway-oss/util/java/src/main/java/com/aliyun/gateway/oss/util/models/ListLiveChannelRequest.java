// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListLiveChannelRequest extends TeaModel {
    /**
     * <p>The name of the LiveChannel from which the list operation starts. LiveChannels whose names are alphabetically after the value of the marker parameter are returned.</p>
     * 
     * <strong>example:</strong>
     * <p>channel-1</p>
     */
    @NameInMap("marker")
    public String marker;

    /**
     * <p>The maximum number of LiveChannels that can be returned for the current request. The value of max-keys cannot exceed 1000. 
     * Default value: 100.</p>
     * 
     * <strong>example:</strong>
     * <p>100</p>
     */
    @NameInMap("max-keys")
    public Long maxKeys;

    /**
     * <p>The prefix that the names of the LiveChannels that you want to return must contain. If you specify a prefix in the request, the specified prefix is included in the response.</p>
     * 
     * <strong>example:</strong>
     * <p>fun/</p>
     */
    @NameInMap("prefix")
    public String prefix;

    public static ListLiveChannelRequest build(java.util.Map<String, ?> map) throws Exception {
        ListLiveChannelRequest self = new ListLiveChannelRequest();
        return TeaModel.build(map, self);
    }

    public ListLiveChannelRequest setMarker(String marker) {
        this.marker = marker;
        return this;
    }
    public String getMarker() {
        return this.marker;
    }

    public ListLiveChannelRequest setMaxKeys(Long maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Long getMaxKeys() {
        return this.maxKeys;
    }

    public ListLiveChannelRequest setPrefix(String prefix) {
        this.prefix = prefix;
        return this;
    }
    public String getPrefix() {
        return this.prefix;
    }

}
