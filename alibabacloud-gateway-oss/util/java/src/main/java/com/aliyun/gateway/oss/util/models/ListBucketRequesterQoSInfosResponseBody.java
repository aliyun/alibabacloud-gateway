// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListBucketRequesterQoSInfosResponseBody extends TeaModel {
    @NameInMap("ListBucketRequesterQoSInfosResult")
    public ListBucketRequesterQoSInfosResult listBucketRequesterQoSInfosResult;

    public static ListBucketRequesterQoSInfosResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListBucketRequesterQoSInfosResponseBody self = new ListBucketRequesterQoSInfosResponseBody();
        return TeaModel.build(map, self);
    }

    public ListBucketRequesterQoSInfosResponseBody setListBucketRequesterQoSInfosResult(ListBucketRequesterQoSInfosResult listBucketRequesterQoSInfosResult) {
        this.listBucketRequesterQoSInfosResult = listBucketRequesterQoSInfosResult;
        return this;
    }
    public ListBucketRequesterQoSInfosResult getListBucketRequesterQoSInfosResult() {
        return this.listBucketRequesterQoSInfosResult;
    }

}
