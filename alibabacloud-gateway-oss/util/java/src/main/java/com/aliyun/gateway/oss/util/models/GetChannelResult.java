// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetChannelResult extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>false</p>
     */
    @NameInMap("AutoSetContentType")
    public Boolean autoSetContentType;

    /**
     * <strong>example:</strong>
     * <p>Fri, 15 Nov 2024 08:11:56 GMT</p>
     */
    @NameInMap("CreateTime")
    public String createTime;

    /**
     * <strong>example:</strong>
     * <p>404.jpg</p>
     */
    @NameInMap("Default404Pic")
    public String default404Pic;

    /**
     * <strong>example:</strong>
     * <p>Fri, 15 Nov 2024 08:11:56 GMT</p>
     */
    @NameInMap("LastModifyTime")
    public String lastModifyTime;

    /**
     * <strong>example:</strong>
     * <p>test-bucket</p>
     */
    @NameInMap("Name")
    public String name;

    /**
     * <strong>example:</strong>
     * <p>false</p>
     */
    @NameInMap("OrigPicForbidden")
    public Boolean origPicForbidden;

    /**
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("SetAttachName")
    public Boolean setAttachName;

    /**
     * <strong>example:</strong>
     * <p>enble</p>
     */
    @NameInMap("Status")
    public String status;

    /**
     * <strong>example:</strong>
     * <p>-,|</p>
     */
    @NameInMap("StyleDelimiters")
    public String styleDelimiters;

    /**
     * <strong>example:</strong>
     * <p>false</p>
     */
    @NameInMap("UseSrcFormat")
    public Boolean useSrcFormat;

    /**
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("UseStyleOnly")
    public Boolean useStyleOnly;

    public static GetChannelResult build(java.util.Map<String, ?> map) throws Exception {
        GetChannelResult self = new GetChannelResult();
        return TeaModel.build(map, self);
    }

    public GetChannelResult setAutoSetContentType(Boolean autoSetContentType) {
        this.autoSetContentType = autoSetContentType;
        return this;
    }
    public Boolean getAutoSetContentType() {
        return this.autoSetContentType;
    }

    public GetChannelResult setCreateTime(String createTime) {
        this.createTime = createTime;
        return this;
    }
    public String getCreateTime() {
        return this.createTime;
    }

    public GetChannelResult setDefault404Pic(String default404Pic) {
        this.default404Pic = default404Pic;
        return this;
    }
    public String getDefault404Pic() {
        return this.default404Pic;
    }

    public GetChannelResult setLastModifyTime(String lastModifyTime) {
        this.lastModifyTime = lastModifyTime;
        return this;
    }
    public String getLastModifyTime() {
        return this.lastModifyTime;
    }

    public GetChannelResult setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return this.name;
    }

    public GetChannelResult setOrigPicForbidden(Boolean origPicForbidden) {
        this.origPicForbidden = origPicForbidden;
        return this;
    }
    public Boolean getOrigPicForbidden() {
        return this.origPicForbidden;
    }

    public GetChannelResult setSetAttachName(Boolean setAttachName) {
        this.setAttachName = setAttachName;
        return this;
    }
    public Boolean getSetAttachName() {
        return this.setAttachName;
    }

    public GetChannelResult setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

    public GetChannelResult setStyleDelimiters(String styleDelimiters) {
        this.styleDelimiters = styleDelimiters;
        return this;
    }
    public String getStyleDelimiters() {
        return this.styleDelimiters;
    }

    public GetChannelResult setUseSrcFormat(Boolean useSrcFormat) {
        this.useSrcFormat = useSrcFormat;
        return this;
    }
    public Boolean getUseSrcFormat() {
        return this.useSrcFormat;
    }

    public GetChannelResult setUseStyleOnly(Boolean useStyleOnly) {
        this.useStyleOnly = useStyleOnly;
        return this;
    }
    public Boolean getUseStyleOnly() {
        return this.useStyleOnly;
    }

}
