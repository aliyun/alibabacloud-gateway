// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CopyObjectsResultFailedObject extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>NoSuchKey</p>
     */
    @NameInMap("ErrorStatus")
    public String errorStatus;

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

    public static CopyObjectsResultFailedObject build(java.util.Map<String, ?> map) throws Exception {
        CopyObjectsResultFailedObject self = new CopyObjectsResultFailedObject();
        return TeaModel.build(map, self);
    }

    public CopyObjectsResultFailedObject setErrorStatus(String errorStatus) {
        this.errorStatus = errorStatus;
        return this;
    }
    public String getErrorStatus() {
        return this.errorStatus;
    }

    public CopyObjectsResultFailedObject setSourceKey(String sourceKey) {
        this.sourceKey = sourceKey;
        return this;
    }
    public String getSourceKey() {
        return this.sourceKey;
    }

    public CopyObjectsResultFailedObject setTargetKey(String targetKey) {
        this.targetKey = targetKey;
        return this;
    }
    public String getTargetKey() {
        return this.targetKey;
    }

}
