// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class GetAccessPointResponseBody extends TeaModel {
    /**
     * <p>The container that stores the information about the access point.</p>
     */
    @NameInMap("GetAccessPointResult")
    public GetAccessPointResult getAccessPointResult;

    public static GetAccessPointResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetAccessPointResponseBody self = new GetAccessPointResponseBody();
        return TeaModel.build(map, self);
    }

    public GetAccessPointResponseBody setGetAccessPointResult(GetAccessPointResult getAccessPointResult) {
        this.getAccessPointResult = getAccessPointResult;
        return this;
    }
    public GetAccessPointResult getGetAccessPointResult() {
        return this.getAccessPointResult;
    }

}
