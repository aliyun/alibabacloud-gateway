// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListCacheResponseBody extends TeaModel {
    @NameInMap("ListAllMyCacheResult")
    public ListAllMyCacheResult listAllMyCacheResult;

    public static ListCacheResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListCacheResponseBody self = new ListCacheResponseBody();
        return TeaModel.build(map, self);
    }

    public ListCacheResponseBody setListAllMyCacheResult(ListAllMyCacheResult listAllMyCacheResult) {
        this.listAllMyCacheResult = listAllMyCacheResult;
        return this;
    }
    public ListAllMyCacheResult getListAllMyCacheResult() {
        return this.listAllMyCacheResult;
    }

}
