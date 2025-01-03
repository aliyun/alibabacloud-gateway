// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class UpdateBucketAntiDDosInfoHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The Anti-DDoS instance ID.</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>cbcac8d2-4f75-4d6d-9f2e-c3447f73****</p>
     */
    @NameInMap("x-oss-defender-instance")
    public String defenderInstance;

    /**
     * <p>The new status of the Anti-DDoS instance. Valid values:</p>
     * <ul>
     * <li>Init: You must specify the custom domain name that you want to protect.</li>
     * <li>Defending: You can select whether to specify the custom domain name that you want to protect.</li>
     * <li>HaltDefending: You do not need to specify the custom domain name that you want to protect.</li>
     * </ul>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>Init</p>
     */
    @NameInMap("x-oss-defender-status")
    public String defenderStatus;

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

    public UpdateBucketAntiDDosInfoHeaders setDefenderInstance(String defenderInstance) {
        this.defenderInstance = defenderInstance;
        return this;
    }
    public String getDefenderInstance() {
        return this.defenderInstance;
    }

    public UpdateBucketAntiDDosInfoHeaders setDefenderStatus(String defenderStatus) {
        this.defenderStatus = defenderStatus;
        return this;
    }
    public String getDefenderStatus() {
        return this.defenderStatus;
    }

}
