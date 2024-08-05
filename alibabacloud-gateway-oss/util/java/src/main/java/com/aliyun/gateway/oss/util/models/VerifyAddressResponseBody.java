// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class VerifyAddressResponseBody extends TeaModel {
    /**
     * <p>1</p>
     */
    @NameInMap("VerifyAddressResponse")
    public VerifyAddressResp verifyAddressResponse;

    public static VerifyAddressResponseBody build(java.util.Map<String, ?> map) throws Exception {
        VerifyAddressResponseBody self = new VerifyAddressResponseBody();
        return TeaModel.build(map, self);
    }

    public VerifyAddressResponseBody setVerifyAddressResponse(VerifyAddressResp verifyAddressResponse) {
        this.verifyAddressResponse = verifyAddressResponse;
        return this;
    }
    public VerifyAddressResp getVerifyAddressResponse() {
        return this.verifyAddressResponse;
    }

}
