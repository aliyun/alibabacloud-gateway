// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListUserRegionsResponseBody extends TeaModel {
    @NameInMap("ListUserRegionsResult")
    public ListUserRegionsResult listUserRegionsResult;

    public static ListUserRegionsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListUserRegionsResponseBody self = new ListUserRegionsResponseBody();
        return TeaModel.build(map, self);
    }

    public ListUserRegionsResponseBody setListUserRegionsResult(ListUserRegionsResult listUserRegionsResult) {
        this.listUserRegionsResult = listUserRegionsResult;
        return this;
    }
    public ListUserRegionsResult getListUserRegionsResult() {
        return this.listUserRegionsResult;
    }

}
