// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BatchPutObjectTagging extends TeaModel {
    @NameInMap("tagging")
    public TagSet tagging;

    public static BatchPutObjectTagging build(java.util.Map<String, ?> map) throws Exception {
        BatchPutObjectTagging self = new BatchPutObjectTagging();
        return TeaModel.build(map, self);
    }

    public BatchPutObjectTagging setTagging(TagSet tagging) {
        this.tagging = tagging;
        return this;
    }
    public TagSet getTagging() {
        return this.tagging;
    }

}
