// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BatchOperationManifestGenerator extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>/</p>
     */
    @NameInMap("Prefix")
    public String prefix;

    /**
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>bucketname</p>
     */
    @NameInMap("SourceBucket")
    public String sourceBucket;

    public static BatchOperationManifestGenerator build(java.util.Map<String, ?> map) throws Exception {
        BatchOperationManifestGenerator self = new BatchOperationManifestGenerator();
        return TeaModel.build(map, self);
    }

    public BatchOperationManifestGenerator setPrefix(String prefix) {
        this.prefix = prefix;
        return this;
    }
    public String getPrefix() {
        return this.prefix;
    }

    public BatchOperationManifestGenerator setSourceBucket(String sourceBucket) {
        this.sourceBucket = sourceBucket;
        return this;
    }
    public String getSourceBucket() {
        return this.sourceBucket;
    }

}
