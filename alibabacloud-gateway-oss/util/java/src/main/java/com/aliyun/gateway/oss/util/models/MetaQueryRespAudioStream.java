// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class MetaQueryRespAudioStream extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>320087</p>
     */
    @NameInMap("Bitrate")
    public Long bitrate;

    /**
     * <strong>example:</strong>
     * <p>2</p>
     */
    @NameInMap("Channels")
    public Long channels;

    /**
     * <strong>example:</strong>
     * <p>aac</p>
     */
    @NameInMap("CodecName")
    public String codecName;

    /**
     * <strong>example:</strong>
     * <p>3.690667</p>
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
     * <p>48000</p>
     */
    @NameInMap("SampleRate")
    public Long sampleRate;

    /**
     * <strong>example:</strong>
     * <p>0.0235</p>
     */
    @NameInMap("StartTime")
    public Double startTime;

    public static MetaQueryRespAudioStream build(java.util.Map<String, ?> map) throws Exception {
        MetaQueryRespAudioStream self = new MetaQueryRespAudioStream();
        return TeaModel.build(map, self);
    }

    public MetaQueryRespAudioStream setBitrate(Long bitrate) {
        this.bitrate = bitrate;
        return this;
    }
    public Long getBitrate() {
        return this.bitrate;
    }

    public MetaQueryRespAudioStream setChannels(Long channels) {
        this.channels = channels;
        return this;
    }
    public Long getChannels() {
        return this.channels;
    }

    public MetaQueryRespAudioStream setCodecName(String codecName) {
        this.codecName = codecName;
        return this;
    }
    public String getCodecName() {
        return this.codecName;
    }

    public MetaQueryRespAudioStream setDuration(Double duration) {
        this.duration = duration;
        return this;
    }
    public Double getDuration() {
        return this.duration;
    }

    public MetaQueryRespAudioStream setLanguage(String language) {
        this.language = language;
        return this;
    }
    public String getLanguage() {
        return this.language;
    }

    public MetaQueryRespAudioStream setSampleRate(Long sampleRate) {
        this.sampleRate = sampleRate;
        return this;
    }
    public Long getSampleRate() {
        return this.sampleRate;
    }

    public MetaQueryRespAudioStream setStartTime(Double startTime) {
        this.startTime = startTime;
        return this;
    }
    public Double getStartTime() {
        return this.startTime;
    }

}
