// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListResourcePoolBucketGroupsResponseBody extends TeaModel {
    @NameInMap("ListResourcePoolBucketGroupsResult")
    public ListResourcePoolBucketGroupsResult listResourcePoolBucketGroupsResult;

    public static ListResourcePoolBucketGroupsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListResourcePoolBucketGroupsResponseBody self = new ListResourcePoolBucketGroupsResponseBody();
        return TeaModel.build(map, self);
    }

    public ListResourcePoolBucketGroupsResponseBody setListResourcePoolBucketGroupsResult(ListResourcePoolBucketGroupsResult listResourcePoolBucketGroupsResult) {
        this.listResourcePoolBucketGroupsResult = listResourcePoolBucketGroupsResult;
        return this;
    }
    public ListResourcePoolBucketGroupsResult getListResourcePoolBucketGroupsResult() {
        return this.listResourcePoolBucketGroupsResult;
    }

}
