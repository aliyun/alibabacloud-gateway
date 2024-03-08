// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class InventoryEncryption extends TeaModel {
    @NameInMap("SSE-KMS")
    public SSEKMS SSE-KMS;

    @NameInMap("SSE-OSS")
    public SSEOSS SSE-OSS;

    public static InventoryEncryption build(java.util.Map<String, ?> map) throws Exception {
        InventoryEncryption self = new InventoryEncryption();
        return TeaModel.build(map, self);
    }

    public InventoryEncryption setSSE-KMS(SSEKMS SSE-KMS) {
        this.SSE-KMS = SSE-KMS;
        return this;
    }
    public SSEKMS getSSE-KMS() {
        return this.SSE-KMS;
    }

    public InventoryEncryption setSSE-OSS(SSEOSS SSE-OSS) {
        this.SSE-OSS = SSE-OSS;
        return this;
    }
    public SSEOSS getSSE-OSS() {
        return this.SSE-OSS;
    }

}
