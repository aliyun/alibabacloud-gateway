// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutStyleRequest extends TeaModel {
    /**
     * <p>The style content.</p>
     */
    @NameInMap("Style")
    public Style style;

    /**
     * <p>The category of the style.</p>
     * 
     * <strong>example:</strong>
     * <p>image</p>
     */
    @NameInMap("category")
    public String category;

    /**
     * <p>The name of the image style.</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>imagestyle</p>
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

    public PutStyleRequest setCategory(String category) {
        this.category = category;
        return this;
    }
    public String getCategory() {
        return this.category;
    }

    public PutStyleRequest setStyleName(String styleName) {
        this.styleName = styleName;
        return this;
    }
    public String getStyleName() {
        return this.styleName;
    }

}
