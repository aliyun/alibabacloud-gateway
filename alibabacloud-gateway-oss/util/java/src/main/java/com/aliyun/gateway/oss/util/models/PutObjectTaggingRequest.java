// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class PutObjectTaggingRequest extends TeaModel {
    /**
     * <p>The container of the tag set.</p>
     */
    @NameInMap("Tagging")
    public Tagging tagging;

    /**
     * <p>The version id of the target object.</p>
     */
    @NameInMap("versionId")
    public String versionId;

    public static PutObjectTaggingRequest build(java.util.Map<String, ?> map) throws Exception {
        PutObjectTaggingRequest self = new PutObjectTaggingRequest();
        return TeaModel.build(map, self);
    }

    public PutObjectTaggingRequest setTagging(Tagging tagging) {
        this.tagging = tagging;
        return this;
    }
    public Tagging getTagging() {
        return this.tagging;
    }

    public PutObjectTaggingRequest setVersionId(String versionId) {
        this.versionId = versionId;
        return this;
    }
    public String getVersionId() {
        return this.versionId;
    }

}
