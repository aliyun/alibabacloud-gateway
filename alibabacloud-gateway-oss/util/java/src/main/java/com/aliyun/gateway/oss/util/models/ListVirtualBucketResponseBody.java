// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListVirtualBucketResponseBody extends TeaModel {
    @NameInMap("ListVirtualBucketResult")
    public ListVirtualBucketResult listVirtualBucketResult;

    public static ListVirtualBucketResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListVirtualBucketResponseBody self = new ListVirtualBucketResponseBody();
        return TeaModel.build(map, self);
    }

    public ListVirtualBucketResponseBody setListVirtualBucketResult(ListVirtualBucketResult listVirtualBucketResult) {
        this.listVirtualBucketResult = listVirtualBucketResult;
        return this;
    }
    public ListVirtualBucketResult getListVirtualBucketResult() {
        return this.listVirtualBucketResult;
    }

}
