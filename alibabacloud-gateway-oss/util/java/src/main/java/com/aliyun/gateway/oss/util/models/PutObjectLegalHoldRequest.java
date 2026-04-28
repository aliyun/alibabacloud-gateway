// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutObjectLegalHoldRequest extends TeaModel {
    @NameInMap("body")
    public Body body;

    public static PutObjectLegalHoldRequest build(java.util.Map<String, ?> map) throws Exception {
        PutObjectLegalHoldRequest self = new PutObjectLegalHoldRequest();
        return TeaModel.build(map, self);
    }

    public PutObjectLegalHoldRequest setBody(Body body) {
        this.body = body;
        return this;
    }
    public Body getBody() {
        return this.body;
    }

    public static class Body extends TeaModel {
        @NameInMap("LegalHold")
        public LegalHold legalHold;

        public static Body build(java.util.Map<String, ?> map) throws Exception {
            Body self = new Body();
            return TeaModel.build(map, self);
        }

        public Body setLegalHold(LegalHold legalHold) {
            this.legalHold = legalHold;
            return this;
        }
        public LegalHold getLegalHold() {
            return this.legalHold;
        }

    }

}
