// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class SSEKMS extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>abcd</p>
     * 
     * <strong>if can be null:</strong>
     * <p>true</p>
     */
    @NameInMap("KeyId")
    public String keyId;

    public static SSEKMS build(java.util.Map<String, ?> map) throws Exception {
        SSEKMS self = new SSEKMS();
        return TeaModel.build(map, self);
    }

    public SSEKMS setKeyId(String keyId) {
        this.keyId = keyId;
        return this;
    }
    public String getKeyId() {
        return this.keyId;
    }

}
