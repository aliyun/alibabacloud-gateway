// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateObjectLinkResult extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>test-bucket</p>
     */
    @NameInMap("Bucket")
    public String bucket;

    /**
     * <strong>example:</strong>
     * <p>&quot;5D3D4F3D1E6C690977E79E413C5F951D&quot;</p>
     */
    @NameInMap("ETag")
    public String ETag;

    /**
     * <strong>example:</strong>
     * <p>ink-object</p>
     */
    @NameInMap("Key")
    public String key;

    public static CreateObjectLinkResult build(java.util.Map<String, ?> map) throws Exception {
        CreateObjectLinkResult self = new CreateObjectLinkResult();
        return TeaModel.build(map, self);
    }

    public CreateObjectLinkResult setBucket(String bucket) {
        this.bucket = bucket;
        return this;
    }
    public String getBucket() {
        return this.bucket;
    }

    public CreateObjectLinkResult setETag(String ETag) {
        this.ETag = ETag;
        return this;
    }
    public String getETag() {
        return this.ETag;
    }

    public CreateObjectLinkResult setKey(String key) {
        this.key = key;
        return this;
    }
    public String getKey() {
        return this.key;
    }

}
