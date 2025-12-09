// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class OpenMetaQueryRequest extends TeaModel {
    @NameInMap("MetaQuery")
    public MetaQueryOpenRequest metaQuery;

    @NameInMap("mode")
    public String mode;

    @NameInMap("role")
    public String role;

    public static OpenMetaQueryRequest build(java.util.Map<String, ?> map) throws Exception {
        OpenMetaQueryRequest self = new OpenMetaQueryRequest();
        return TeaModel.build(map, self);
    }

    public OpenMetaQueryRequest setMetaQuery(MetaQueryOpenRequest metaQuery) {
        this.metaQuery = metaQuery;
        return this;
    }
    public MetaQueryOpenRequest getMetaQuery() {
        return this.metaQuery;
    }

    public OpenMetaQueryRequest setMode(String mode) {
        this.mode = mode;
        return this;
    }
    public String getMode() {
        return this.mode;
    }

    public OpenMetaQueryRequest setRole(String role) {
        this.role = role;
        return this;
    }
    public String getRole() {
        return this.role;
    }

}
