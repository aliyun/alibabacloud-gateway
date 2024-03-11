// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class GetSymlinkRequest extends TeaModel {
    /**
     * <p>The version of the object to which the symbolic link points.</p>
     */
    @NameInMap("versionId")
    public String versionId;

    public static GetSymlinkRequest build(java.util.Map<String, ?> map) throws Exception {
        GetSymlinkRequest self = new GetSymlinkRequest();
        return TeaModel.build(map, self);
    }

    public GetSymlinkRequest setVersionId(String versionId) {
        this.versionId = versionId;
        return this;
    }
    public String getVersionId() {
        return this.versionId;
    }

}
