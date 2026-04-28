// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class LegalHold extends TeaModel {
    @NameInMap("Status")
    public String status;

    public static LegalHold build(java.util.Map<String, ?> map) throws Exception {
        LegalHold self = new LegalHold();
        return TeaModel.build(map, self);
    }

    public LegalHold setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

}
