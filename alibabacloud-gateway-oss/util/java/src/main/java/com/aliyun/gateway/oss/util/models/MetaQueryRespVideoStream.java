// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class MetaQueryRespVideoStream extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>8</p>
     */
    @NameInMap("BitDepth")
    public Long bitDepth;

    /**
     * <strong>example:</strong>
     * <p>5407765</p>
     */
    @NameInMap("Bitrate")
    public Long bitrate;

    /**
     * <strong>example:</strong>
     * <p>h264</p>
     */
    @NameInMap("CodecName")
    public String codecName;

    /**
     * <strong>example:</strong>
     * <p>bt709</p>
     */
    @NameInMap("ColorSpace")
    public String colorSpace;

    /**
     * <strong>example:</strong>
     * <p>22.88</p>
     */
    @NameInMap("Duration")
    public Double duration;

    /**
     * <strong>example:</strong>
     * <p>572</p>
     */
    @NameInMap("FrameCount")
    public Long frameCount;

    /**
     * <strong>example:</strong>
     * <p>25/1</p>
     */
    @NameInMap("FrameRate")
    public String frameRate;

    /**
     * <strong>example:</strong>
     * <p>720</p>
     */
    @NameInMap("Height")
    public Long height;

    /**
     * <strong>example:</strong>
     * <p>en</p>
     */
    @NameInMap("Language")
    public String language;

    /**
     * <strong>example:</strong>
     * <p>yuv420p</p>
     */
    @NameInMap("PixelFormat")
    public String pixelFormat;

    /**
     * <strong>example:</strong>
     * <p>0.000000</p>
     */
    @NameInMap("StartTime")
    public Double startTime;

    /**
     * <strong>example:</strong>
     * <p>1280</p>
     */
    @NameInMap("Width")
    public Long width;

    public static MetaQueryRespVideoStream build(java.util.Map<String, ?> map) throws Exception {
        MetaQueryRespVideoStream self = new MetaQueryRespVideoStream();
        return TeaModel.build(map, self);
    }

    public MetaQueryRespVideoStream setBitDepth(Long bitDepth) {
        this.bitDepth = bitDepth;
        return this;
    }
    public Long getBitDepth() {
        return this.bitDepth;
    }

    public MetaQueryRespVideoStream setBitrate(Long bitrate) {
        this.bitrate = bitrate;
        return this;
    }
    public Long getBitrate() {
        return this.bitrate;
    }

    public MetaQueryRespVideoStream setCodecName(String codecName) {
        this.codecName = codecName;
        return this;
    }
    public String getCodecName() {
        return this.codecName;
    }

    public MetaQueryRespVideoStream setColorSpace(String colorSpace) {
        this.colorSpace = colorSpace;
        return this;
    }
    public String getColorSpace() {
        return this.colorSpace;
    }

    public MetaQueryRespVideoStream setDuration(Double duration) {
        this.duration = duration;
        return this;
    }
    public Double getDuration() {
        return this.duration;
    }

    public MetaQueryRespVideoStream setFrameCount(Long frameCount) {
        this.frameCount = frameCount;
        return this;
    }
    public Long getFrameCount() {
        return this.frameCount;
    }

    public MetaQueryRespVideoStream setFrameRate(String frameRate) {
        this.frameRate = frameRate;
        return this;
    }
    public String getFrameRate() {
        return this.frameRate;
    }

    public MetaQueryRespVideoStream setHeight(Long height) {
        this.height = height;
        return this;
    }
    public Long getHeight() {
        return this.height;
    }

    public MetaQueryRespVideoStream setLanguage(String language) {
        this.language = language;
        return this;
    }
    public String getLanguage() {
        return this.language;
    }

    public MetaQueryRespVideoStream setPixelFormat(String pixelFormat) {
        this.pixelFormat = pixelFormat;
        return this;
    }
    public String getPixelFormat() {
        return this.pixelFormat;
    }

    public MetaQueryRespVideoStream setStartTime(Double startTime) {
        this.startTime = startTime;
        return this;
    }
    public Double getStartTime() {
        return this.startTime;
    }

    public MetaQueryRespVideoStream setWidth(Long width) {
        this.width = width;
        return this;
    }
    public Long getWidth() {
        return this.width;
    }

}
