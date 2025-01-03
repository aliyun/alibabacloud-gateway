// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetAccessPointForObjectProcessHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The name of the Object FC Access Point. The name of an Object FC Access Point must meet the following requirements:</p>
     * <p>The name cannot exceed 63 characters in length.</p>
     * <p>The name can contain only lowercase letters, digits, and hyphens (-) and cannot start or end with a hyphen (-).</p>
     * <p>The name must be unique in the current region.</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>fc-ap-01</p>
     */
    @NameInMap("x-oss-access-point-for-object-process-name")
    public String xOssAccessPointForObjectProcessName;

    public static GetAccessPointForObjectProcessHeaders build(java.util.Map<String, ?> map) throws Exception {
        GetAccessPointForObjectProcessHeaders self = new GetAccessPointForObjectProcessHeaders();
        return TeaModel.build(map, self);
    }

    public GetAccessPointForObjectProcessHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public GetAccessPointForObjectProcessHeaders setXOssAccessPointForObjectProcessName(String xOssAccessPointForObjectProcessName) {
        this.xOssAccessPointForObjectProcessName = xOssAccessPointForObjectProcessName;
        return this;
    }
    public String getXOssAccessPointForObjectProcessName() {
        return this.xOssAccessPointForObjectProcessName;
    }

}
