// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutObjectRetentionRequest extends TeaModel {
    @NameInMap("body")
    public Body body;

    public static PutObjectRetentionRequest build(java.util.Map<String, ?> map) throws Exception {
        PutObjectRetentionRequest self = new PutObjectRetentionRequest();
        return TeaModel.build(map, self);
    }

    public PutObjectRetentionRequest setBody(Body body) {
        this.body = body;
        return this;
    }
    public Body getBody() {
        return this.body;
    }

    public static class Body extends TeaModel {
        @NameInMap("Retention")
        public Retention retention;

        public static Body build(java.util.Map<String, ?> map) throws Exception {
            Body self = new Body();
            return TeaModel.build(map, self);
        }

        public Body setRetention(Retention retention) {
            this.retention = retention;
            return this;
        }
        public Retention getRetention() {
            return this.retention;
        }

    }

}
