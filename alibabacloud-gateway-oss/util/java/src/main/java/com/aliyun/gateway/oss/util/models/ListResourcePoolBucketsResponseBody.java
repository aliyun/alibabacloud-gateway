// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListResourcePoolBucketsResponseBody extends TeaModel {
    @NameInMap("ListResourcePoolBucketsResult")
    public ListResourcePoolBucketsResult listResourcePoolBucketsResult;

    public static ListResourcePoolBucketsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListResourcePoolBucketsResponseBody self = new ListResourcePoolBucketsResponseBody();
        return TeaModel.build(map, self);
    }

    public ListResourcePoolBucketsResponseBody setListResourcePoolBucketsResult(ListResourcePoolBucketsResult listResourcePoolBucketsResult) {
        this.listResourcePoolBucketsResult = listResourcePoolBucketsResult;
        return this;
    }
    public ListResourcePoolBucketsResult getListResourcePoolBucketsResult() {
        return this.listResourcePoolBucketsResult;
    }

}
