// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class UpdateBucketAntiDDosInfoHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The Anti-DDoS instance ID.</p>
     */
    @NameInMap("x-oss-defender-instance")
    public String xOssDefenderInstance;

    /**
     * <p>The new status of the Anti-DDoS instance. Valid values:</p>
     * <br>
     * <p>*   Init: You must specify the custom domain name that you want to protect.</p>
     * <p>*   Defending: You can select whether to specify the custom domain name that you want to protect.</p>
     * <p>*   HaltDefending: You do not need to specify the custom domain name that you want to protect.</p>
     */
    @NameInMap("x-oss-defender-status")
    public String xOssDefenderStatus;

    public static UpdateBucketAntiDDosInfoHeaders build(java.util.Map<String, ?> map) throws Exception {
        UpdateBucketAntiDDosInfoHeaders self = new UpdateBucketAntiDDosInfoHeaders();
        return TeaModel.build(map, self);
    }

    public UpdateBucketAntiDDosInfoHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public UpdateBucketAntiDDosInfoHeaders setXOssDefenderInstance(String xOssDefenderInstance) {
        this.xOssDefenderInstance = xOssDefenderInstance;
        return this;
    }
    public String getXOssDefenderInstance() {
        return this.xOssDefenderInstance;
    }

    public UpdateBucketAntiDDosInfoHeaders setXOssDefenderStatus(String xOssDefenderStatus) {
        this.xOssDefenderStatus = xOssDefenderStatus;
        return this;
    }
    public String getXOssDefenderStatus() {
        return this.xOssDefenderStatus;
    }

}
