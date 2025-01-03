// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetAccessPointPolicyForObjectProcessHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The name of the Object FC Access Point.</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>fc-ap-01</p>
     */
    @NameInMap("x-oss-access-point-for-object-process-name")
    public String xOssAccessPointForObjectProcessName;

    public static GetAccessPointPolicyForObjectProcessHeaders build(java.util.Map<String, ?> map) throws Exception {
        GetAccessPointPolicyForObjectProcessHeaders self = new GetAccessPointPolicyForObjectProcessHeaders();
        return TeaModel.build(map, self);
    }

    public GetAccessPointPolicyForObjectProcessHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public GetAccessPointPolicyForObjectProcessHeaders setXOssAccessPointForObjectProcessName(String xOssAccessPointForObjectProcessName) {
        this.xOssAccessPointForObjectProcessName = xOssAccessPointForObjectProcessName;
        return this;
    }
    public String getXOssAccessPointForObjectProcessName() {
        return this.xOssAccessPointForObjectProcessName;
    }

}
