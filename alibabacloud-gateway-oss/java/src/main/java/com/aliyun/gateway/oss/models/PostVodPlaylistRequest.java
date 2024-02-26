// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class PostVodPlaylistRequest extends TeaModel {
    /**
     * <p>The end time of the time range during which the TS files that you want to query are generated, </p>
     * <p>which is a Unix timestamp.</p>
     * <p>> The value of EndTime must be later than the value of StartTime. The duration between EndTime and StartTime must be shorter than one day.</p>
     */
    @NameInMap("endTime")
    public String endTime;

    /**
     * <p>The start time of the time range during which the TS files that you want to query are generated, which is a Unix timestamp.</p>
     */
    @NameInMap("startTime")
    public String startTime;

    public static PostVodPlaylistRequest build(java.util.Map<String, ?> map) throws Exception {
        PostVodPlaylistRequest self = new PostVodPlaylistRequest();
        return TeaModel.build(map, self);
    }

    public PostVodPlaylistRequest setEndTime(String endTime) {
        this.endTime = endTime;
        return this;
    }
    public String getEndTime() {
        return this.endTime;
    }

    public PostVodPlaylistRequest setStartTime(String startTime) {
        this.startTime = startTime;
        return this;
    }
    public String getStartTime() {
        return this.startTime;
    }

}
