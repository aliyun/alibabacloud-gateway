// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetObjectInfoResponseBody extends TeaModel {
    @NameInMap("GetObjectInfoResult")
    public GetObjectInfoResult getObjectInfoResult;

    public static GetObjectInfoResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetObjectInfoResponseBody self = new GetObjectInfoResponseBody();
        return TeaModel.build(map, self);
    }

    public GetObjectInfoResponseBody setGetObjectInfoResult(GetObjectInfoResult getObjectInfoResult) {
        this.getObjectInfoResult = getObjectInfoResult;
        return this;
    }
    public GetObjectInfoResult getGetObjectInfoResult() {
        return this.getObjectInfoResult;
    }

}
