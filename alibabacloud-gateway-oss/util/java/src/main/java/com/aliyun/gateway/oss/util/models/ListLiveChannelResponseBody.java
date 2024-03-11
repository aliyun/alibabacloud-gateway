// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class ListLiveChannelResponseBody extends TeaModel {
    /**
     * <p>The container that stores the results of the ListLiveChannel request.</p>
     */
    @NameInMap("ListLiveChannelResult")
    public ListLiveChannelResult listLiveChannelResult;

    public static ListLiveChannelResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListLiveChannelResponseBody self = new ListLiveChannelResponseBody();
        return TeaModel.build(map, self);
    }

    public ListLiveChannelResponseBody setListLiveChannelResult(ListLiveChannelResult listLiveChannelResult) {
        this.listLiveChannelResult = listLiveChannelResult;
        return this;
    }
    public ListLiveChannelResult getListLiveChannelResult() {
        return this.listLiveChannelResult;
    }

    public static class ListLiveChannelResult extends TeaModel {
        /**
         * <p>Indicates whether all results are returned.</p>
         * <p>- true: All results are returned.</p>
         * <p>- false: Not all results are returned.</p>
         */
        @NameInMap("IsTruncated")
        public Boolean isTruncated;

        /**
         * <p>The container that stores the information about each returned LiveChannel.</p>
         */
        @NameInMap("LiveChannel")
        public java.util.List<LiveChannel> liveChannels;

        /**
         * <p>The name of the LiveChannel after which the ListLiveChannel operation starts.</p>
         */
        @NameInMap("Marker")
        public String marker;

        /**
         * <p>The maximum number of returned LiveChannels in the response.</p>
         */
        @NameInMap("MaxKeys")
        public Long maxKeys;

        /**
         * <p>If not all results are returned, the NextMarker parameter is included in the response to indicate the Marker value of the next request.</p>
         */
        @NameInMap("NextMarker")
        public String nextMarker;

        /**
         * <p>The prefix that the names of the returned LiveChannels contain.</p>
         */
        @NameInMap("Prefix")
        public String prefix;

        public static ListLiveChannelResult build(java.util.Map<String, ?> map) throws Exception {
            ListLiveChannelResult self = new ListLiveChannelResult();
            return TeaModel.build(map, self);
        }

        public ListLiveChannelResult setIsTruncated(Boolean isTruncated) {
            this.isTruncated = isTruncated;
            return this;
        }
        public Boolean getIsTruncated() {
            return this.isTruncated;
        }

        public ListLiveChannelResult setLiveChannels(java.util.List<LiveChannel> liveChannels) {
            this.liveChannels = liveChannels;
            return this;
        }
        public java.util.List<LiveChannel> getLiveChannels() {
            return this.liveChannels;
        }

        public ListLiveChannelResult setMarker(String marker) {
            this.marker = marker;
            return this;
        }
        public String getMarker() {
            return this.marker;
        }

        public ListLiveChannelResult setMaxKeys(Long maxKeys) {
            this.maxKeys = maxKeys;
            return this;
        }
        public Long getMaxKeys() {
            return this.maxKeys;
        }

        public ListLiveChannelResult setNextMarker(String nextMarker) {
            this.nextMarker = nextMarker;
            return this;
        }
        public String getNextMarker() {
            return this.nextMarker;
        }

        public ListLiveChannelResult setPrefix(String prefix) {
            this.prefix = prefix;
            return this;
        }
        public String getPrefix() {
            return this.prefix;
        }

    }

}
