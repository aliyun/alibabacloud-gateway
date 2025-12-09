// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BatchOperationManifestLocation extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>bucketname</p>
     */
    @NameInMap("Bucket")
    public String bucket;

    /**
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>objectname</p>
     */
    @NameInMap("Object")
    public String object;

    /**
     * <strong>example:</strong>
     * <p><strong><strong><strong><strong>gICx0puvmxkiIDBkODc3M2RlZjgyNjQ0NDViZGVlYmRmNzI2</strong></strong></strong></strong></p>
     */
    @NameInMap("VersionId")
    public String versionId;

    public static BatchOperationManifestLocation build(java.util.Map<String, ?> map) throws Exception {
        BatchOperationManifestLocation self = new BatchOperationManifestLocation();
        return TeaModel.build(map, self);
    }

    public BatchOperationManifestLocation setBucket(String bucket) {
        this.bucket = bucket;
        return this;
    }
    public String getBucket() {
        return this.bucket;
    }

    public BatchOperationManifestLocation setObject(String object) {
        this.object = object;
        return this;
    }
    public String getObject() {
        return this.object;
    }

    public BatchOperationManifestLocation setVersionId(String versionId) {
        this.versionId = versionId;
        return this;
    }
    public String getVersionId() {
        return this.versionId;
    }

}
