// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class InitBucketAntiDDosInfoHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The ID of the Anti-DDoS instance.</p>
     */
    @NameInMap("x-oss-defender-instance")
    public String xOssDefenderInstance;

    /**
     * <p>The type of the Anti-DDoS instance. Set the value to AntiDDosPremimum.</p>
     */
    @NameInMap("x-oss-defender-type")
    public String xOssDefenderType;

    public static InitBucketAntiDDosInfoHeaders build(java.util.Map<String, ?> map) throws Exception {
        InitBucketAntiDDosInfoHeaders self = new InitBucketAntiDDosInfoHeaders();
        return TeaModel.build(map, self);
    }

    public InitBucketAntiDDosInfoHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public InitBucketAntiDDosInfoHeaders setXOssDefenderInstance(String xOssDefenderInstance) {
        this.xOssDefenderInstance = xOssDefenderInstance;
        return this;
    }
    public String getXOssDefenderInstance() {
        return this.xOssDefenderInstance;
    }

    public InitBucketAntiDDosInfoHeaders setXOssDefenderType(String xOssDefenderType) {
        this.xOssDefenderType = xOssDefenderType;
        return this;
    }
    public String getXOssDefenderType() {
        return this.xOssDefenderType;
    }

}
