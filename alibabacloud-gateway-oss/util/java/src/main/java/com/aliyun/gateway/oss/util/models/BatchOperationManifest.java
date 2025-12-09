// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BatchOperationManifest extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("Location")
    public BatchOperationManifestLocation location;

    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("Spec")
    public BatchOperationJobSpec spec;

    public static BatchOperationManifest build(java.util.Map<String, ?> map) throws Exception {
        BatchOperationManifest self = new BatchOperationManifest();
        return TeaModel.build(map, self);
    }

    public BatchOperationManifest setLocation(BatchOperationManifestLocation location) {
        this.location = location;
        return this;
    }
    public BatchOperationManifestLocation getLocation() {
        return this.location;
    }

    public BatchOperationManifest setSpec(BatchOperationJobSpec spec) {
        this.spec = spec;
        return this;
    }
    public BatchOperationJobSpec getSpec() {
        return this.spec;
    }

}
