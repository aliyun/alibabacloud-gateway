// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ListLiveChannelResponseBody extends TeaModel {
    /**
     * <p>Indicates whether the returned results are truncated.</p>
     * <br>
     * <p>- true: indicates that all results are returned.</p>
     * <p>- false: indicates that not all results are returned.</p>
     */
    @NameInMap("IsTruncated")
    public Boolean isTruncated;

    /**
     * <p>The container that stores the information about each returned LiveChannel.</p>
     */
    @NameInMap("LiveChannel")
    public java.util.List<LiveChannel> liveChannel;

    /**
     * <p>The name of the LiveChannel from which the ListLiveChannel operation begins.</p>
     */
    @NameInMap("Marker")
    public String marker;

    /**
     * <p>The maximum number of buckets returned each time.</p>
     */
    @NameInMap("MaxKeys")
    public Long maxKeys;

    /**
     * <p>If not all results are returned, the NextMarker element is included in the response to indicate the Marker value of the next request.</p>
     */
    @NameInMap("NextMarker")
    public String nextMarker;

    /**
     * <p>The prefix that the names of the returned LiveChannels contain.</p>
     */
    @NameInMap("Prefix")
    public String prefix;

    public static ListLiveChannelResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListLiveChannelResponseBody self = new ListLiveChannelResponseBody();
        return TeaModel.build(map, self);
    }

    public ListLiveChannelResponseBody setIsTruncated(Boolean isTruncated) {
        this.isTruncated = isTruncated;
        return this;
    }
    public Boolean getIsTruncated() {
        return this.isTruncated;
    }

    public ListLiveChannelResponseBody setLiveChannel(java.util.List<LiveChannel> liveChannel) {
        this.liveChannel = liveChannel;
        return this;
    }
    public java.util.List<LiveChannel> getLiveChannel() {
        return this.liveChannel;
    }

    public ListLiveChannelResponseBody setMarker(String marker) {
        this.marker = marker;
        return this;
    }
    public String getMarker() {
        return this.marker;
    }

    public ListLiveChannelResponseBody setMaxKeys(Long maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Long getMaxKeys() {
        return this.maxKeys;
    }

    public ListLiveChannelResponseBody setNextMarker(String nextMarker) {
        this.nextMarker = nextMarker;
        return this;
    }
    public String getNextMarker() {
        return this.nextMarker;
    }

    public ListLiveChannelResponseBody setPrefix(String prefix) {
        this.prefix = prefix;
        return this;
    }
    public String getPrefix() {
        return this.prefix;
    }

}
