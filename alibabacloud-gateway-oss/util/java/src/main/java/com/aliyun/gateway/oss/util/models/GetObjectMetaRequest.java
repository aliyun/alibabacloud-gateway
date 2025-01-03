// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetObjectMetaRequest extends TeaModel {
    /**
     * <p>The versionID of the object.</p>
     * 
     * <strong>example:</strong>
     * <p>CAEQNRiBgIDMh4mD0BYiIDUzNDA4OGNmZjBjYTQ0YmI4Y2I4ZmVlYzJlNGVk****</p>
     */
    @NameInMap("versionId")
    public String versionId;

    public static GetObjectMetaRequest build(java.util.Map<String, ?> map) throws Exception {
        GetObjectMetaRequest self = new GetObjectMetaRequest();
        return TeaModel.build(map, self);
    }

    public GetObjectMetaRequest setVersionId(String versionId) {
        this.versionId = versionId;
        return this;
    }
    public String getVersionId() {
        return this.versionId;
    }

}
