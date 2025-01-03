// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListResourcePoolsResponseBody extends TeaModel {
    @NameInMap("ListResourcePoolsResult")
    public ListResourcePoolsResult listResourcePoolsResult;

    public static ListResourcePoolsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListResourcePoolsResponseBody self = new ListResourcePoolsResponseBody();
        return TeaModel.build(map, self);
    }

    public ListResourcePoolsResponseBody setListResourcePoolsResult(ListResourcePoolsResult listResourcePoolsResult) {
        this.listResourcePoolsResult = listResourcePoolsResult;
        return this;
    }
    public ListResourcePoolsResult getListResourcePoolsResult() {
        return this.listResourcePoolsResult;
    }

}
