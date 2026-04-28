// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutObjectRetentionShrinkRequest extends TeaModel {
    @NameInMap("body")
    public String bodyShrink;

    public static PutObjectRetentionShrinkRequest build(java.util.Map<String, ?> map) throws Exception {
        PutObjectRetentionShrinkRequest self = new PutObjectRetentionShrinkRequest();
        return TeaModel.build(map, self);
    }

    public PutObjectRetentionShrinkRequest setBodyShrink(String bodyShrink) {
        this.bodyShrink = bodyShrink;
        return this;
    }
    public String getBodyShrink() {
        return this.bodyShrink;
    }

}
