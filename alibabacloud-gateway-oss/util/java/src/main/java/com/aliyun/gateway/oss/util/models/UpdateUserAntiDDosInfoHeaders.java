// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class UpdateUserAntiDDosInfoHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The Anti-DDoS instance ID.</p>
     */
    @NameInMap("x-oss-defender-instance")
    public String xOssDefenderInstance;

    /**
     * <p>The new status of the Anti-DDoS instance. Set the value to HaltDefending, which indicates that the Anti-DDos protection is disabled for a bucket.</p>
     */
    @NameInMap("x-oss-defender-status")
    public String xOssDefenderStatus;

    public static UpdateUserAntiDDosInfoHeaders build(java.util.Map<String, ?> map) throws Exception {
        UpdateUserAntiDDosInfoHeaders self = new UpdateUserAntiDDosInfoHeaders();
        return TeaModel.build(map, self);
    }

    public UpdateUserAntiDDosInfoHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public UpdateUserAntiDDosInfoHeaders setXOssDefenderInstance(String xOssDefenderInstance) {
        this.xOssDefenderInstance = xOssDefenderInstance;
        return this;
    }
    public String getXOssDefenderInstance() {
        return this.xOssDefenderInstance;
    }

    public UpdateUserAntiDDosInfoHeaders setXOssDefenderStatus(String xOssDefenderStatus) {
        this.xOssDefenderStatus = xOssDefenderStatus;
        return this;
    }
    public String getXOssDefenderStatus() {
        return this.xOssDefenderStatus;
    }

}
