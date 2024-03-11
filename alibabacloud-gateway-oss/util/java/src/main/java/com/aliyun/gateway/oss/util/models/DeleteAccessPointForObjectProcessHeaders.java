// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class DeleteAccessPointForObjectProcessHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    @NameInMap("x-oss-access-point-for-object-process-name")
    public String xOssAccessPointForObjectProcessName;

    public static DeleteAccessPointForObjectProcessHeaders build(java.util.Map<String, ?> map) throws Exception {
        DeleteAccessPointForObjectProcessHeaders self = new DeleteAccessPointForObjectProcessHeaders();
        return TeaModel.build(map, self);
    }

    public DeleteAccessPointForObjectProcessHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public DeleteAccessPointForObjectProcessHeaders setXOssAccessPointForObjectProcessName(String xOssAccessPointForObjectProcessName) {
        this.xOssAccessPointForObjectProcessName = xOssAccessPointForObjectProcessName;
        return this;
    }
    public String getXOssAccessPointForObjectProcessName() {
        return this.xOssAccessPointForObjectProcessName;
    }

}
