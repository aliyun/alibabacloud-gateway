// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketHashRequest extends TeaModel {
    @NameInMap("ObjectHashConfiguration")
    public ObjectHashConfiguration objectHashConfiguration;

    public static PutBucketHashRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketHashRequest self = new PutBucketHashRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketHashRequest setObjectHashConfiguration(ObjectHashConfiguration objectHashConfiguration) {
        this.objectHashConfiguration = objectHashConfiguration;
        return this;
    }
    public ObjectHashConfiguration getObjectHashConfiguration() {
        return this.objectHashConfiguration;
    }

}
