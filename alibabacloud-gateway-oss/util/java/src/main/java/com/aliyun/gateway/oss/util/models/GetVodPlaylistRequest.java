// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetVodPlaylistRequest extends TeaModel {
    /**
     * <p>The end time of the time range during which the TS files that you want to query are generated in the Unix timestamp format. </p>
     * <blockquote>
     * <p>The value of EndTime must be greater than the value of StartTime. The duration between EndTime and StartTime must be less than one day.</p>
     * </blockquote>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>1636618271</p>
     */
    @NameInMap("endTime")
    public String endTime;

    /**
     * <p>The start time of the time range during which the TS files that you want to query are generated in the Unix timestamp format.</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>1636600271</p>
     */
    @NameInMap("startTime")
    public String startTime;

    public static GetVodPlaylistRequest build(java.util.Map<String, ?> map) throws Exception {
        GetVodPlaylistRequest self = new GetVodPlaylistRequest();
        return TeaModel.build(map, self);
    }

    public GetVodPlaylistRequest setEndTime(String endTime) {
        this.endTime = endTime;
        return this;
    }
    public String getEndTime() {
        return this.endTime;
    }

    public GetVodPlaylistRequest setStartTime(String startTime) {
        this.startTime = startTime;
        return this;
    }
    public String getStartTime() {
        return this.startTime;
    }

}
