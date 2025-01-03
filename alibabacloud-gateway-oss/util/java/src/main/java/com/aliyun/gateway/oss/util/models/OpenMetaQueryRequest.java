// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class OpenMetaQueryRequest extends TeaModel {
    @NameInMap("mode")
    public String mode;

    public static OpenMetaQueryRequest build(java.util.Map<String, ?> map) throws Exception {
        OpenMetaQueryRequest self = new OpenMetaQueryRequest();
        return TeaModel.build(map, self);
    }

    public OpenMetaQueryRequest setMode(String mode) {
        this.mode = mode;
        return this;
    }
    public String getMode() {
        return this.mode;
    }

}
