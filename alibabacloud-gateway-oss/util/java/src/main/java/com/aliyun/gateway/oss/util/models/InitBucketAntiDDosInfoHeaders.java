// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class InitBucketAntiDDosInfoHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The ID of the Anti-DDoS instance.</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>cbcac8d2-4f75-4d6d-9f2e-c3447f73****</p>
     */
    @NameInMap("x-oss-defender-instance")
    public String defenderInstance;

    /**
     * <p>The type of the Anti-DDoS instance. Set the value to AntiDDosPremimum.</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>AntiDDosPremimum</p>
     */
    @NameInMap("x-oss-defender-type")
    public String defenderType;

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

    public InitBucketAntiDDosInfoHeaders setDefenderInstance(String defenderInstance) {
        this.defenderInstance = defenderInstance;
        return this;
    }
    public String getDefenderInstance() {
        return this.defenderInstance;
    }

    public InitBucketAntiDDosInfoHeaders setDefenderType(String defenderType) {
        this.defenderType = defenderType;
        return this;
    }
    public String getDefenderType() {
        return this.defenderType;
    }

}
