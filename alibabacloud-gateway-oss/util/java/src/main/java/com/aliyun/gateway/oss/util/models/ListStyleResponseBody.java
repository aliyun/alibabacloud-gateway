// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ListStyleResponseBody extends TeaModel {
    /**
     * <p>The container that was used to query the information about the specified image style.</p>
     */
    @NameInMap("Style")
    public java.util.List<StyleInfo> style;

    public static ListStyleResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListStyleResponseBody self = new ListStyleResponseBody();
        return TeaModel.build(map, self);
    }

    public ListStyleResponseBody setStyle(java.util.List<StyleInfo> style) {
        this.style = style;
        return this;
    }
    public java.util.List<StyleInfo> getStyle() {
        return this.style;
    }

}
