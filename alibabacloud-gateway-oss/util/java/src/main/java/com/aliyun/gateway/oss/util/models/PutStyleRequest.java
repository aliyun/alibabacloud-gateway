// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class PutStyleRequest extends TeaModel {
    /**
     * <p>The style content.</p>
     */
    @NameInMap("Style")
    public Style style;

    /**
     * <p>The name of the image style.</p>
     */
    @NameInMap("styleName")
    public String styleName;

    public static PutStyleRequest build(java.util.Map<String, ?> map) throws Exception {
        PutStyleRequest self = new PutStyleRequest();
        return TeaModel.build(map, self);
    }

    public PutStyleRequest setStyle(Style style) {
        this.style = style;
        return this;
    }
    public Style getStyle() {
        return this.style;
    }

    public PutStyleRequest setStyleName(String styleName) {
        this.styleName = styleName;
        return this;
    }
    public String getStyleName() {
        return this.styleName;
    }

}
