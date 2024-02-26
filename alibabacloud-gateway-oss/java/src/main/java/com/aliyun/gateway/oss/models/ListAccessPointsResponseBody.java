// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class ListAccessPointsResponseBody extends TeaModel {
    /**
     * <p>The container that stores the information about access points.</p>
     */
    @NameInMap("ListAccessPointsResult")
    public ListAccessPointsResult listAccessPointsResult;

    public static ListAccessPointsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListAccessPointsResponseBody self = new ListAccessPointsResponseBody();
        return TeaModel.build(map, self);
    }

    public ListAccessPointsResponseBody setListAccessPointsResult(ListAccessPointsResult listAccessPointsResult) {
        this.listAccessPointsResult = listAccessPointsResult;
        return this;
    }
    public ListAccessPointsResult getListAccessPointsResult() {
        return this.listAccessPointsResult;
    }

}
