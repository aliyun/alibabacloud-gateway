// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetAccessPointPolicyHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The name of the access point.</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>ap-01</p>
     */
    @NameInMap("x-oss-access-point-name")
    public String xOssAccessPointName;

    public static GetAccessPointPolicyHeaders build(java.util.Map<String, ?> map) throws Exception {
        GetAccessPointPolicyHeaders self = new GetAccessPointPolicyHeaders();
        return TeaModel.build(map, self);
    }

    public GetAccessPointPolicyHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public GetAccessPointPolicyHeaders setXOssAccessPointName(String xOssAccessPointName) {
        this.xOssAccessPointName = xOssAccessPointName;
        return this;
    }
    public String getXOssAccessPointName() {
        return this.xOssAccessPointName;
    }

}
