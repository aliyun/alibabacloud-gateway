// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetUserQoSInfoHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <strong>example:</strong>
     * <p>false</p>
     */
    @NameInMap("x-oss-return-default")
    public Boolean xOssReturnDefault;

    public static GetUserQoSInfoHeaders build(java.util.Map<String, ?> map) throws Exception {
        GetUserQoSInfoHeaders self = new GetUserQoSInfoHeaders();
        return TeaModel.build(map, self);
    }

    public GetUserQoSInfoHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public GetUserQoSInfoHeaders setXOssReturnDefault(Boolean xOssReturnDefault) {
        this.xOssReturnDefault = xOssReturnDefault;
        return this;
    }
    public Boolean getXOssReturnDefault() {
        return this.xOssReturnDefault;
    }

}
