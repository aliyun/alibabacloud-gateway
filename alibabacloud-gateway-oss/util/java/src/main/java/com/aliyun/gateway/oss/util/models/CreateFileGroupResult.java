// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateFileGroupResult extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>test-bucket</p>
     */
    @NameInMap("Bucket")
    public String bucket;

    /**
     * <strong>example:</strong>
     * <p>&quot;06A4DDABDD5F65868B8C5919E76487D6&quot;</p>
     */
    @NameInMap("ETag")
    public String ETag;

    /**
     * <strong>example:</strong>
     * <p>test.txt</p>
     */
    @NameInMap("Key")
    public String key;

    /**
     * <strong>example:</strong>
     * <p>384</p>
     */
    @NameInMap("Size")
    public Long size;

    public static CreateFileGroupResult build(java.util.Map<String, ?> map) throws Exception {
        CreateFileGroupResult self = new CreateFileGroupResult();
        return TeaModel.build(map, self);
    }

    public CreateFileGroupResult setBucket(String bucket) {
        this.bucket = bucket;
        return this;
    }
    public String getBucket() {
        return this.bucket;
    }

    public CreateFileGroupResult setETag(String ETag) {
        this.ETag = ETag;
        return this;
    }
    public String getETag() {
        return this.ETag;
    }

    public CreateFileGroupResult setKey(String key) {
        this.key = key;
        return this;
    }
    public String getKey() {
        return this.key;
    }

    public CreateFileGroupResult setSize(Long size) {
        this.size = size;
        return this;
    }
    public Long getSize() {
        return this.size;
    }

}
