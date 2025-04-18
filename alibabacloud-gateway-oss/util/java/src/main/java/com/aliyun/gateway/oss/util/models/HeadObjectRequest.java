// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class HeadObjectRequest extends TeaModel {
    /**
     * <p>The version ID of the object for which you want to query metadata.</p>
     * 
     * <strong>example:</strong>
     * <p>CAEQMxiBgMCZov2D0BYiIDY4MDllOTc2YmY5MjQxMzdiOGI3OTlhNTU0ODIx****</p>
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
