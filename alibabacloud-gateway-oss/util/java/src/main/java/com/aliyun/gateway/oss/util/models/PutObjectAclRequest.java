// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutObjectAclRequest extends TeaModel {
    /**
     * <p>The version id of the object.</p>
     * 
     * <strong>example:</strong>
     * <p>CAEQMhiBgIC3rpSD0BYiIDBjYTk5MmIzN2JlNjQxZTFiNGIzM2E3OTliODA0****</p>
     */
    @NameInMap("versionId")
    public String versionId;

    public static PutObjectAclRequest build(java.util.Map<String, ?> map) throws Exception {
        PutObjectAclRequest self = new PutObjectAclRequest();
        return TeaModel.build(map, self);
    }

    public PutObjectAclRequest setVersionId(String versionId) {
        this.versionId = versionId;
        return this;
    }
    public String getVersionId() {
        return this.versionId;
    }

}
