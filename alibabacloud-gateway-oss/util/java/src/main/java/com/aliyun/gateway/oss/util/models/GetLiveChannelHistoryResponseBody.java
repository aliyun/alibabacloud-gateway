// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetLiveChannelHistoryResponseBody extends TeaModel {
    /**
     * <p>Specifies the container that stores a stream pushing record.</p>
     */
    @NameInMap("LiveRecord")
    public java.util.List<LiveRecord> liveRecord;

    public static GetLiveChannelHistoryResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetLiveChannelHistoryResponseBody self = new GetLiveChannelHistoryResponseBody();
        return TeaModel.build(map, self);
    }

    public GetLiveChannelHistoryResponseBody setLiveRecord(java.util.List<LiveRecord> liveRecord) {
        this.liveRecord = liveRecord;
        return this;
    }
    public java.util.List<LiveRecord> getLiveRecord() {
        return this.liveRecord;
    }

}
