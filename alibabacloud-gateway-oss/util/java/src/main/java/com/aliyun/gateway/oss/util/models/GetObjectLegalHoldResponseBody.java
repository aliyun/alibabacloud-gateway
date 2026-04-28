// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetObjectLegalHoldResponseBody extends TeaModel {
    @NameInMap("LegalHold")
    public LegalHold legalHold;

    public static GetObjectLegalHoldResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetObjectLegalHoldResponseBody self = new GetObjectLegalHoldResponseBody();
        return TeaModel.build(map, self);
    }

    public GetObjectLegalHoldResponseBody setLegalHold(LegalHold legalHold) {
        this.legalHold = legalHold;
        return this;
    }
    public LegalHold getLegalHold() {
        return this.legalHold;
    }

}
