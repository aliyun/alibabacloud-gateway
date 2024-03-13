// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketTagsRequest extends TeaModel {
    /**
     * <p>The container used to store TagSet.</p>
     */
    @NameInMap("Tagging")
    public Tagging tagging;

    public static PutBucketTagsRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketTagsRequest self = new PutBucketTagsRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketTagsRequest setTagging(Tagging tagging) {
        this.tagging = tagging;
        return this;
    }
    public Tagging getTagging() {
        return this.tagging;
    }

}
