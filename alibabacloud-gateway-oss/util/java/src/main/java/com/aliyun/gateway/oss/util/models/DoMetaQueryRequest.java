// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class DoMetaQueryRequest extends TeaModel {
    /**
     * <p>The containerfor query conditions.</p>
     */
    @NameInMap("MetaQuery")
    public MetaQuery metaQuery;

    public static DoMetaQueryRequest build(java.util.Map<String, ?> map) throws Exception {
        DoMetaQueryRequest self = new DoMetaQueryRequest();
        return TeaModel.build(map, self);
    }

    public DoMetaQueryRequest setMetaQuery(MetaQuery metaQuery) {
        this.metaQuery = metaQuery;
        return this;
    }
    public MetaQuery getMetaQuery() {
        return this.metaQuery;
    }

}
