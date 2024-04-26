// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetStyleResponseBody extends TeaModel {
    /**
     * <p>The container that stores the information about the image style.</p>
     */
    @NameInMap("Style")
    public StyleInfo style;

    public static GetStyleResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetStyleResponseBody self = new GetStyleResponseBody();
        return TeaModel.build(map, self);
    }

    public GetStyleResponseBody setStyle(StyleInfo style) {
        this.style = style;
        return this;
    }
    public StyleInfo getStyle() {
        return this.style;
    }

}
