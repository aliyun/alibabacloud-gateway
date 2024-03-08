// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class GetAccessPointHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The name of the access point.</p>
     */
    @NameInMap("x-oss-access-point-name")
    public String xOssAccessPointName;

    public static GetAccessPointHeaders build(java.util.Map<String, ?> map) throws Exception {
        GetAccessPointHeaders self = new GetAccessPointHeaders();
        return TeaModel.build(map, self);
    }

    public GetAccessPointHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public GetAccessPointHeaders setXOssAccessPointName(String xOssAccessPointName) {
        this.xOssAccessPointName = xOssAccessPointName;
        return this;
    }
    public String getXOssAccessPointName() {
        return this.xOssAccessPointName;
    }

}
