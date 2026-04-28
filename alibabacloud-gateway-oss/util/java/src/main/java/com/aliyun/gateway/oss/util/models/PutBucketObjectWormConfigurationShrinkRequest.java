// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketObjectWormConfigurationShrinkRequest extends TeaModel {
    @NameInMap("body")
    public String bodyShrink;

    public static PutBucketObjectWormConfigurationShrinkRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketObjectWormConfigurationShrinkRequest self = new PutBucketObjectWormConfigurationShrinkRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketObjectWormConfigurationShrinkRequest setBodyShrink(String bodyShrink) {
        this.bodyShrink = bodyShrink;
        return this;
    }
    public String getBodyShrink() {
        return this.bodyShrink;
    }

}
