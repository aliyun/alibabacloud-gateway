// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutObjectLegalHoldShrinkRequest extends TeaModel {
    @NameInMap("body")
    public String bodyShrink;

    public static PutObjectLegalHoldShrinkRequest build(java.util.Map<String, ?> map) throws Exception {
        PutObjectLegalHoldShrinkRequest self = new PutObjectLegalHoldShrinkRequest();
        return TeaModel.build(map, self);
    }

    public PutObjectLegalHoldShrinkRequest setBodyShrink(String bodyShrink) {
        this.bodyShrink = bodyShrink;
        return this;
    }
    public String getBodyShrink() {
        return this.bodyShrink;
    }

}
