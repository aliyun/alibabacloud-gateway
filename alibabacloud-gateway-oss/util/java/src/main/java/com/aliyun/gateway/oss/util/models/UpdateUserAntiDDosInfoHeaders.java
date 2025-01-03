// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class UpdateUserAntiDDosInfoHeaders extends TeaModel {
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
     * <p>The new status of the Anti-DDoS instance. Set the value to HaltDefending, which indicates that the Anti-DDos protection is disabled for a bucket.</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>HaltDefending</p>
     */
    @NameInMap("x-oss-defender-status")
    public String defenderStatus;

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

    public UpdateUserAntiDDosInfoHeaders setDefenderInstance(String defenderInstance) {
        this.defenderInstance = defenderInstance;
        return this;
    }
    public String getDefenderInstance() {
        return this.defenderInstance;
    }

    public UpdateUserAntiDDosInfoHeaders setDefenderStatus(String defenderStatus) {
        this.defenderStatus = defenderStatus;
        return this;
    }
    public String getDefenderStatus() {
        return this.defenderStatus;
    }

}
