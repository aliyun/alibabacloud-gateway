// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class LiveChannelAudio extends TeaModel {
    @NameInMap("Bandwidth")
    public Long bandwidth;

    @NameInMap("Codec")
    public String codec;

    @NameInMap("SampleRate")
    public Long sampleRate;

    public static LiveChannelAudio build(java.util.Map<String, ?> map) throws Exception {
        LiveChannelAudio self = new LiveChannelAudio();
        return TeaModel.build(map, self);
    }

    public LiveChannelAudio setBandwidth(Long bandwidth) {
        this.bandwidth = bandwidth;
        return this;
    }
    public Long getBandwidth() {
        return this.bandwidth;
    }

    public LiveChannelAudio setCodec(String codec) {
        this.codec = codec;
        return this;
    }
    public String getCodec() {
        return this.codec;
    }

    public LiveChannelAudio setSampleRate(Long sampleRate) {
        this.sampleRate = sampleRate;
        return this;
    }
    public Long getSampleRate() {
        return this.sampleRate;
    }

}
