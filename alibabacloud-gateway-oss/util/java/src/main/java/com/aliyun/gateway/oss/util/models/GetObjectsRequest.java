// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetObjectsRequest extends TeaModel {
    @NameInMap("GetObjectsRequest")
    public GetObjectsReq getObjectsRequest;

    public static GetObjectsRequest build(java.util.Map<String, ?> map) throws Exception {
        GetObjectsRequest self = new GetObjectsRequest();
        return TeaModel.build(map, self);
    }

    public GetObjectsRequest setGetObjectsRequest(GetObjectsReq getObjectsRequest) {
        this.getObjectsRequest = getObjectsRequest;
        return this;
    }
    public GetObjectsReq getGetObjectsRequest() {
        return this.getObjectsRequest;
    }

}
