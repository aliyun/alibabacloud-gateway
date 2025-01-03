// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListResourcePoolRequesterQoSInfosResponseBody extends TeaModel {
    @NameInMap("ListResourcePoolRequesterQoSInfosResult")
    public ListResourcePoolRequesterQoSInfosResult listResourcePoolRequesterQoSInfosResult;

    public static ListResourcePoolRequesterQoSInfosResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListResourcePoolRequesterQoSInfosResponseBody self = new ListResourcePoolRequesterQoSInfosResponseBody();
        return TeaModel.build(map, self);
    }

    public ListResourcePoolRequesterQoSInfosResponseBody setListResourcePoolRequesterQoSInfosResult(ListResourcePoolRequesterQoSInfosResult listResourcePoolRequesterQoSInfosResult) {
        this.listResourcePoolRequesterQoSInfosResult = listResourcePoolRequesterQoSInfosResult;
        return this;
    }
    public ListResourcePoolRequesterQoSInfosResult getListResourcePoolRequesterQoSInfosResult() {
        return this.listResourcePoolRequesterQoSInfosResult;
    }

}
