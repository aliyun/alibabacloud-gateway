// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ChannelInfo extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>False</p>
     */
    @NameInMap("AutoSetContentType")
    public Boolean autoSetContentType;

    /**
     * <strong>example:</strong>
     * <p>test-channel</p>
     */
    @NameInMap("Name")
    public String name;

    /**
     * <strong>example:</strong>
     * <p>True</p>
     */
    @NameInMap("OrigPicForbidden")
    public Boolean origPicForbidden;

    /**
     * <strong>example:</strong>
     * <p>True</p>
     */
    @NameInMap("SetAttachName")
    public Boolean setAttachName;

    /**
     * <strong>example:</strong>
     * <p>enable</p>
     */
    @NameInMap("Status")
    public String status;

    /**
     * <strong>example:</strong>
     * <ul>
     * <li></li>
     * </ul>
     */
    @NameInMap("StyleDelimiters")
    public String styleDelimiters;

    /**
     * <strong>example:</strong>
     * <p>True</p>
     */
    @NameInMap("UseSrcFormat")
    public Boolean useSrcFormat;

    /**
     * <strong>example:</strong>
     * <p>False</p>
     */
    @NameInMap("UseStyleOnly")
    public Boolean useStyleOnly;

    public static ChannelInfo build(java.util.Map<String, ?> map) throws Exception {
        ChannelInfo self = new ChannelInfo();
        return TeaModel.build(map, self);
    }

    public ChannelInfo setAutoSetContentType(Boolean autoSetContentType) {
        this.autoSetContentType = autoSetContentType;
        return this;
    }
    public Boolean getAutoSetContentType() {
        return this.autoSetContentType;
    }

    public ChannelInfo setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return this.name;
    }

    public ChannelInfo setOrigPicForbidden(Boolean origPicForbidden) {
        this.origPicForbidden = origPicForbidden;
        return this;
    }
    public Boolean getOrigPicForbidden() {
        return this.origPicForbidden;
    }

    public ChannelInfo setSetAttachName(Boolean setAttachName) {
        this.setAttachName = setAttachName;
        return this;
    }
    public Boolean getSetAttachName() {
        return this.setAttachName;
    }

    public ChannelInfo setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

    public ChannelInfo setStyleDelimiters(String styleDelimiters) {
        this.styleDelimiters = styleDelimiters;
        return this;
    }
    public String getStyleDelimiters() {
        return this.styleDelimiters;
    }

    public ChannelInfo setUseSrcFormat(Boolean useSrcFormat) {
        this.useSrcFormat = useSrcFormat;
        return this;
    }
    public Boolean getUseSrcFormat() {
        return this.useSrcFormat;
    }

    public ChannelInfo setUseStyleOnly(Boolean useStyleOnly) {
        this.useStyleOnly = useStyleOnly;
        return this;
    }
    public Boolean getUseStyleOnly() {
        return this.useStyleOnly;
    }

}
