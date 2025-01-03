// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class RestoreObjectRequest extends TeaModel {
    /**
     * <p>The container that stores information about the RestoreObject request.</p>
     */
    @NameInMap("RestoreRequest")
    public RestoreRequest restoreRequest;

    /**
     * <p>The version number of the object that you want to restore.</p>
     * 
     * <strong>example:</strong>
     * <p>CAEQNRiBgMClj7qD0BYiIDQ5Y2QyMjc3NGZkODRlMTU5M2VkY2U3MWRiNGRh</p>
     */
    @NameInMap("versionId")
    public String versionId;

    public static RestoreObjectRequest build(java.util.Map<String, ?> map) throws Exception {
        RestoreObjectRequest self = new RestoreObjectRequest();
        return TeaModel.build(map, self);
    }

    public RestoreObjectRequest setRestoreRequest(RestoreRequest restoreRequest) {
        this.restoreRequest = restoreRequest;
        return this;
    }
    public RestoreRequest getRestoreRequest() {
        return this.restoreRequest;
    }

    public RestoreObjectRequest setVersionId(String versionId) {
        this.versionId = versionId;
        return this;
    }
    public String getVersionId() {
        return this.versionId;
    }

}
