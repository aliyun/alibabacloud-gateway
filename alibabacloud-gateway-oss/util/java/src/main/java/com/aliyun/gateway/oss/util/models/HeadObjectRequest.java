// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class HeadObjectRequest extends TeaModel {
    /**
     * <p>The version ID of the object for which you want to query metadata.</p>
     */
    @NameInMap("versionId")
    public String versionId;

    public static HeadObjectRequest build(java.util.Map<String, ?> map) throws Exception {
        HeadObjectRequest self = new HeadObjectRequest();
        return TeaModel.build(map, self);
    }

    public HeadObjectRequest setVersionId(String versionId) {
        this.versionId = versionId;
        return this;
    }
    public String getVersionId() {
        return this.versionId;
    }

}
