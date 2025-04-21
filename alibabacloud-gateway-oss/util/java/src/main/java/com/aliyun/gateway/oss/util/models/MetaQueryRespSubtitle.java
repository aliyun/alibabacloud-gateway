// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class MetaQueryRespSubtitle extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>mov_text</p>
     */
    @NameInMap("CodecName")
    public String codecName;

    /**
     * <strong>example:</strong>
     * <p>71.378</p>
     */
    @NameInMap("Duration")
    public Double duration;

    /**
     * <strong>example:</strong>
     * <p>en</p>
     */
    @NameInMap("Language")
    public String language;

    /**
     * <strong>example:</strong>
     * <p>0.000000</p>
     */
    @NameInMap("StartTime")
    public Double startTime;

    public static MetaQueryRespSubtitle build(java.util.Map<String, ?> map) throws Exception {
        MetaQueryRespSubtitle self = new MetaQueryRespSubtitle();
        return TeaModel.build(map, self);
    }

    public MetaQueryRespSubtitle setCodecName(String codecName) {
        this.codecName = codecName;
        return this;
    }
    public String getCodecName() {
        return this.codecName;
    }

    public MetaQueryRespSubtitle setDuration(Double duration) {
        this.duration = duration;
        return this;
    }
    public Double getDuration() {
        return this.duration;
    }

    public MetaQueryRespSubtitle setLanguage(String language) {
        this.language = language;
        return this;
    }
    public String getLanguage() {
        return this.language;
    }

    public MetaQueryRespSubtitle setStartTime(Double startTime) {
        this.startTime = startTime;
        return this;
    }
    public Double getStartTime() {
        return this.startTime;
    }

}
