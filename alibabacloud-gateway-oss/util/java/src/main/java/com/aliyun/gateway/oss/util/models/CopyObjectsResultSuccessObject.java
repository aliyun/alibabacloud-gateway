// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CopyObjectsResultSuccessObject extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>5EB63BBBE01EEED093CB22BB8F5ACDC3</p>
     */
    @NameInMap("ETag")
    public String ETag;

    /**
     * <strong>example:</strong>
     * <p>abc/source.txt</p>
     */
    @NameInMap("SourceKey")
    public String sourceKey;

    /**
     * <strong>example:</strong>
     * <p>abc/target.txt</p>
     */
    @NameInMap("TargetKey")
    public String targetKey;

    public static CopyObjectsResultSuccessObject build(java.util.Map<String, ?> map) throws Exception {
        CopyObjectsResultSuccessObject self = new CopyObjectsResultSuccessObject();
        return TeaModel.build(map, self);
    }

    public CopyObjectsResultSuccessObject setETag(String ETag) {
        this.ETag = ETag;
        return this;
    }
    public String getETag() {
        return this.ETag;
    }

    public CopyObjectsResultSuccessObject setSourceKey(String sourceKey) {
        this.sourceKey = sourceKey;
        return this;
    }
    public String getSourceKey() {
        return this.sourceKey;
    }

    public CopyObjectsResultSuccessObject setTargetKey(String targetKey) {
        this.targetKey = targetKey;
        return this;
    }
    public String getTargetKey() {
        return this.targetKey;
    }

}
