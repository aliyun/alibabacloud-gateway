// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class GetLiveChannelHistoryResponseBody extends TeaModel {
    /**
     * <p>The container that stores the returned results of the GetLiveChannelHistory request.</p>
     */
    @NameInMap("LiveChannelHistory")
    public LiveChannelHistory liveChannelHistory;

    public static GetLiveChannelHistoryResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetLiveChannelHistoryResponseBody self = new GetLiveChannelHistoryResponseBody();
        return TeaModel.build(map, self);
    }

    public GetLiveChannelHistoryResponseBody setLiveChannelHistory(LiveChannelHistory liveChannelHistory) {
        this.liveChannelHistory = liveChannelHistory;
        return this;
    }
    public LiveChannelHistory getLiveChannelHistory() {
        return this.liveChannelHistory;
    }

    public static class LiveChannelHistory extends TeaModel {
        /**
         * <p>The container that stores a list of stream pushing records.</p>
         */
        @NameInMap("LiveRecord")
        public java.util.List<LiveRecord> liveRecord;

        public static LiveChannelHistory build(java.util.Map<String, ?> map) throws Exception {
            LiveChannelHistory self = new LiveChannelHistory();
            return TeaModel.build(map, self);
        }

        public LiveChannelHistory setLiveRecord(java.util.List<LiveRecord> liveRecord) {
            this.liveRecord = liveRecord;
            return this;
        }
        public java.util.List<LiveRecord> getLiveRecord() {
            return this.liveRecord;
        }

    }

}
