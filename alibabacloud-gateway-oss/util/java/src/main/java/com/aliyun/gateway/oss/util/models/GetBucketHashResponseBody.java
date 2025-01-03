// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketHashResponseBody extends TeaModel {
    @NameInMap("ObjectHashConfiguration")
    public ObjectHashConfiguration objectHashConfiguration;

    public static GetBucketHashResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketHashResponseBody self = new GetBucketHashResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketHashResponseBody setObjectHashConfiguration(ObjectHashConfiguration objectHashConfiguration) {
        this.objectHashConfiguration = objectHashConfiguration;
        return this;
    }
    public ObjectHashConfiguration getObjectHashConfiguration() {
        return this.objectHashConfiguration;
    }

}
