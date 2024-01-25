// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetLiveChannelStatResponseBody extends TeaModel {
    /**
     * <p>The container that stores audio stream information when Status is set to Live. </p>
     * <p>>  The Video and Audio containers can be returned only when Status is set to Live. However, the Video and Audio containers may not be returned when Status is set to Live. For example, if the client has connected to the LiveChannel but no audio or video stream is sent, these two containers are not returned.</p>
     */
    @NameInMap("Audio")
    public LiveChannelAudio audio;

    /**
     * <p>If the value of Status is Live, this parameter indicates the time when the current client start to push streams. The value of this parameter is in the ISO8601 format.</p>
     */
    @NameInMap("ConnectedTime")
    public String connectedTime;

    /**
     * <p>If the value of Status is Live, this parameter indicates the IP address of the current client that pushes streams.</p>
     */
    @NameInMap("RemoteAddr")
    public String remoteAddr;

    /**
     * <p>Indicates the current stream pushing status of the LiveChannel.</p>
     * <p><br>Valid value: Disabled, Live, and Idle</p>
     */
    @NameInMap("Status")
    public String status;

    /**
     * <p>The container that stores video stream information when Status is set to Live. </p>
     * <p>> Video and audio containers can be returned only when Status is set to Live. However, these two containers may not be returned when Status is set to Live. For example, if the client has connected to the LiveChannel but no audio or video stream is sent, these two containers are not returned.</p>
     */
    @NameInMap("Video")
    public LiveChannelVideo video;

    public static GetLiveChannelStatResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetLiveChannelStatResponseBody self = new GetLiveChannelStatResponseBody();
        return TeaModel.build(map, self);
    }

    public GetLiveChannelStatResponseBody setAudio(LiveChannelAudio audio) {
        this.audio = audio;
        return this;
    }
    public LiveChannelAudio getAudio() {
        return this.audio;
    }

    public GetLiveChannelStatResponseBody setConnectedTime(String connectedTime) {
        this.connectedTime = connectedTime;
        return this;
    }
    public String getConnectedTime() {
        return this.connectedTime;
    }

    public GetLiveChannelStatResponseBody setRemoteAddr(String remoteAddr) {
        this.remoteAddr = remoteAddr;
        return this;
    }
    public String getRemoteAddr() {
        return this.remoteAddr;
    }

    public GetLiveChannelStatResponseBody setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

    public GetLiveChannelStatResponseBody setVideo(LiveChannelVideo video) {
        this.video = video;
        return this;
    }
    public LiveChannelVideo getVideo() {
        return this.video;
    }

}
