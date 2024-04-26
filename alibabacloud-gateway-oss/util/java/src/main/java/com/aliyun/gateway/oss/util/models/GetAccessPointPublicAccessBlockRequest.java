// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetAccessPointPublicAccessBlockRequest extends TeaModel {
    @NameInMap("x-oss-access-point-name")
    public String xOssAccessPointName;

    public static GetAccessPointPublicAccessBlockRequest build(java.util.Map<String, ?> map) throws Exception {
        GetAccessPointPublicAccessBlockRequest self = new GetAccessPointPublicAccessBlockRequest();
        return TeaModel.build(map, self);
    }

    public GetAccessPointPublicAccessBlockRequest setXOssAccessPointName(String xOssAccessPointName) {
        this.xOssAccessPointName = xOssAccessPointName;
        return this;
    }
    public String getXOssAccessPointName() {
        return this.xOssAccessPointName;
    }

}
