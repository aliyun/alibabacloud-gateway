// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class InventoryEncryption extends TeaModel {
    @NameInMap("SSE-KMS")
    public SSEKMS SSEKMS;

    @NameInMap("SSE-OSS")
    public String SSEOSS;

    public static InventoryEncryption build(java.util.Map<String, ?> map) throws Exception {
        InventoryEncryption self = new InventoryEncryption();
        return TeaModel.build(map, self);
    }

    public InventoryEncryption setSSEKMS(SSEKMS SSEKMS) {
        this.SSEKMS = SSEKMS;
        return this;
    }
    public SSEKMS getSSEKMS() {
        return this.SSEKMS;
    }

    public InventoryEncryption setSSEOSS(String SSEOSS) {
        this.SSEOSS = SSEOSS;
        return this;
    }
    public String getSSEOSS() {
        return this.SSEOSS;
    }

}
