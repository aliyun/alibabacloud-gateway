// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class DeleteAccessPointHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The name of the access point.</p>
     */
    @NameInMap("x-oss-access-point-name")
    public String xOssAccessPointName;

    public static DeleteAccessPointHeaders build(java.util.Map<String, ?> map) throws Exception {
        DeleteAccessPointHeaders self = new DeleteAccessPointHeaders();
        return TeaModel.build(map, self);
    }

    public DeleteAccessPointHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public DeleteAccessPointHeaders setXOssAccessPointName(String xOssAccessPointName) {
        this.xOssAccessPointName = xOssAccessPointName;
        return this;
    }
    public String getXOssAccessPointName() {
        return this.xOssAccessPointName;
    }

}
