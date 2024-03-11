// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class GetLiveChannelStatResponseBody extends TeaModel {
    /**
     * <p>The container that stores the returned results of the GetLiveChannelStat request.</p>
     */
    @NameInMap("LiveChannelStat")
    public LiveChannelStat liveChannelStat;

    public static GetLiveChannelStatResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetLiveChannelStatResponseBody self = new GetLiveChannelStatResponseBody();
        return TeaModel.build(map, self);
    }

    public GetLiveChannelStatResponseBody setLiveChannelStat(LiveChannelStat liveChannelStat) {
        this.liveChannelStat = liveChannelStat;
        return this;
    }
    public LiveChannelStat getLiveChannelStat() {
        return this.liveChannelStat;
    }

    public static class LiveChannelStat extends TeaModel {
        /**
         * <p>The container that stores audio stream information if Status is set to Live.</p>
         * <p>>Video and audio containers can be returned only if Status is set to Live. However, these two containers may not necessarily be returned if Status is set to Live. For example, if the client has connected to the LiveChannel but no audio or video stream is sent, these two containers are not returned.</p>
         */
        @NameInMap("Audio")
        public LiveChannelAudio audio;

        /**
         * <p>If Status is set to Live, this element indicates the time when the current client starts to ingest streams. The value of the element is in the ISO 8601 format.</p>
         */
        @NameInMap("ConnectedTime")
        public String connectedTime;

        /**
         * <p>If Status is set to Live, this element indicates the IP address of the current client that ingests streams.</p>
         */
        @NameInMap("RemoteAddr")
        public String remoteAddr;

        /**
         * <p>The current stream ingestion status of the LiveChannel. Valid value: Disabled、Live、Idle。</p>
         */
        @NameInMap("Status")
        public String status;

        /**
         * <p>The container that stores video stream information if Status is set to Live.</p>
         * <br>
         * <p>>Video and audio containers can be returned only if Status is set to Live. However, these two containers may not necessarily be returned if Status is set to Live. For example, if the client has connected to the LiveChannel but no audio or video stream is sent, these two containers are not returned.</p>
         */
        @NameInMap("Video")
        public LiveChannelVideo video;

        public static LiveChannelStat build(java.util.Map<String, ?> map) throws Exception {
            LiveChannelStat self = new LiveChannelStat();
            return TeaModel.build(map, self);
        }

        public LiveChannelStat setAudio(LiveChannelAudio audio) {
            this.audio = audio;
            return this;
        }
        public LiveChannelAudio getAudio() {
            return this.audio;
        }

        public LiveChannelStat setConnectedTime(String connectedTime) {
            this.connectedTime = connectedTime;
            return this;
        }
        public String getConnectedTime() {
            return this.connectedTime;
        }

        public LiveChannelStat setRemoteAddr(String remoteAddr) {
            this.remoteAddr = remoteAddr;
            return this;
        }
        public String getRemoteAddr() {
            return this.remoteAddr;
        }

        public LiveChannelStat setStatus(String status) {
            this.status = status;
            return this;
        }
        public String getStatus() {
            return this.status;
        }

        public LiveChannelStat setVideo(LiveChannelVideo video) {
            this.video = video;
            return this;
        }
        public LiveChannelVideo getVideo() {
            return this.video;
        }

    }

}
