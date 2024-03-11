// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class DeleteStyleRequest extends TeaModel {
    /**
     * <p>The name of the image style.</p>
     */
    @NameInMap("styleName")
    public String styleName;

    public static DeleteStyleRequest build(java.util.Map<String, ?> map) throws Exception {
        DeleteStyleRequest self = new DeleteStyleRequest();
        return TeaModel.build(map, self);
    }

    public DeleteStyleRequest setStyleName(String styleName) {
        this.styleName = styleName;
        return this;
    }
    public String getStyleName() {
        return this.styleName;
    }

}
