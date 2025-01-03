// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutChannelConfiguration extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>False</p>
     */
    @NameInMap("AutoSetContentType")
    public Boolean autoSetContentType;

    /**
     * <strong>example:</strong>
     * <p>404.jpg</p>
     */
    @NameInMap("Default404Pic")
    public String default404Pic;

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

    public static PutChannelConfiguration build(java.util.Map<String, ?> map) throws Exception {
        PutChannelConfiguration self = new PutChannelConfiguration();
        return TeaModel.build(map, self);
    }

    public PutChannelConfiguration setAutoSetContentType(Boolean autoSetContentType) {
        this.autoSetContentType = autoSetContentType;
        return this;
    }
    public Boolean getAutoSetContentType() {
        return this.autoSetContentType;
    }

    public PutChannelConfiguration setDefault404Pic(String default404Pic) {
        this.default404Pic = default404Pic;
        return this;
    }
    public String getDefault404Pic() {
        return this.default404Pic;
    }

    public PutChannelConfiguration setOrigPicForbidden(Boolean origPicForbidden) {
        this.origPicForbidden = origPicForbidden;
        return this;
    }
    public Boolean getOrigPicForbidden() {
        return this.origPicForbidden;
    }

    public PutChannelConfiguration setSetAttachName(Boolean setAttachName) {
        this.setAttachName = setAttachName;
        return this;
    }
    public Boolean getSetAttachName() {
        return this.setAttachName;
    }

    public PutChannelConfiguration setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

    public PutChannelConfiguration setStyleDelimiters(String styleDelimiters) {
        this.styleDelimiters = styleDelimiters;
        return this;
    }
    public String getStyleDelimiters() {
        return this.styleDelimiters;
    }

    public PutChannelConfiguration setUseSrcFormat(Boolean useSrcFormat) {
        this.useSrcFormat = useSrcFormat;
        return this;
    }
    public Boolean getUseSrcFormat() {
        return this.useSrcFormat;
    }

    public PutChannelConfiguration setUseStyleOnly(Boolean useStyleOnly) {
        this.useStyleOnly = useStyleOnly;
        return this;
    }
    public Boolean getUseStyleOnly() {
        return this.useStyleOnly;
    }

}
