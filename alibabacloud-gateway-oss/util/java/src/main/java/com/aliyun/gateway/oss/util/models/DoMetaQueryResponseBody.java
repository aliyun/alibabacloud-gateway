// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DoMetaQueryResponseBody extends TeaModel {
    /**
     * <p>The container for the query result.</p>
     */
    @NameInMap("MetaQuery")
    public MetaQueryResp metaQuery;

    public static DoMetaQueryResponseBody build(java.util.Map<String, ?> map) throws Exception {
        DoMetaQueryResponseBody self = new DoMetaQueryResponseBody();
        return TeaModel.build(map, self);
    }

    public DoMetaQueryResponseBody setMetaQuery(MetaQueryResp metaQuery) {
        this.metaQuery = metaQuery;
        return this;
    }
    public MetaQueryResp getMetaQuery() {
        return this.metaQuery;
    }

}
