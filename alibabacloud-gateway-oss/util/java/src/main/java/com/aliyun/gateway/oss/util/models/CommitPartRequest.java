// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CommitPartRequest extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("partUploadId")
    public String partUploadId;

    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("uploadId")
    public String uploadId;

    public static CommitPartRequest build(java.util.Map<String, ?> map) throws Exception {
        CommitPartRequest self = new CommitPartRequest();
        return TeaModel.build(map, self);
    }

    public CommitPartRequest setPartUploadId(String partUploadId) {
        this.partUploadId = partUploadId;
        return this;
    }
    public String getPartUploadId() {
        return this.partUploadId;
    }

    public CommitPartRequest setUploadId(String uploadId) {
        this.uploadId = uploadId;
        return this;
    }
    public String getUploadId() {
        return this.uploadId;
    }

}
