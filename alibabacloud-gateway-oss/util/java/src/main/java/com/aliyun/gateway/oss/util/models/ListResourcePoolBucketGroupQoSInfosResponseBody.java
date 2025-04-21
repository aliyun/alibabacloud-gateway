// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListResourcePoolBucketGroupQoSInfosResponseBody extends TeaModel {
    @NameInMap("ListResourcePoolBucketGroupQoSInfosResult")
    public ListResourcePoolBucketGroupQoSInfosResult listResourcePoolBucketGroupQoSInfosResult;

    public static ListResourcePoolBucketGroupQoSInfosResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListResourcePoolBucketGroupQoSInfosResponseBody self = new ListResourcePoolBucketGroupQoSInfosResponseBody();
        return TeaModel.build(map, self);
    }

    public ListResourcePoolBucketGroupQoSInfosResponseBody setListResourcePoolBucketGroupQoSInfosResult(ListResourcePoolBucketGroupQoSInfosResult listResourcePoolBucketGroupQoSInfosResult) {
        this.listResourcePoolBucketGroupQoSInfosResult = listResourcePoolBucketGroupQoSInfosResult;
        return this;
    }
    public ListResourcePoolBucketGroupQoSInfosResult getListResourcePoolBucketGroupQoSInfosResult() {
        return this.listResourcePoolBucketGroupQoSInfosResult;
    }

}
