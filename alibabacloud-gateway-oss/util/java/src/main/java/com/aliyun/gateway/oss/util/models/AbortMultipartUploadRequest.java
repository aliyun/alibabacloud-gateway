// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class AbortMultipartUploadRequest extends TeaModel {
    /**
     * <p>The ID of the multipart upload task.</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>0004B9895DBBB6E****</p>
     */
    @NameInMap("uploadId")
    public String uploadId;

    public static AbortMultipartUploadRequest build(java.util.Map<String, ?> map) throws Exception {
        AbortMultipartUploadRequest self = new AbortMultipartUploadRequest();
        return TeaModel.build(map, self);
    }

    public AbortMultipartUploadRequest setUploadId(String uploadId) {
        this.uploadId = uploadId;
        return this;
    }
    public String getUploadId() {
        return this.uploadId;
    }

}
