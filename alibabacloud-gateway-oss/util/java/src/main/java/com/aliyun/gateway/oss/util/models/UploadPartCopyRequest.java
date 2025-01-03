// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class UploadPartCopyRequest extends TeaModel {
    /**
     * <p>The number of parts.</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>10000</p>
     */
    @NameInMap("partNumber")
    public Long partNumber;

    /**
     * <p>The ID that identifies the object to which the parts to upload belong.</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>0004B999EF5A239BB9138C6227D69F95</p>
     */
    @NameInMap("uploadId")
    public String uploadId;

    public static UploadPartCopyRequest build(java.util.Map<String, ?> map) throws Exception {
        UploadPartCopyRequest self = new UploadPartCopyRequest();
        return TeaModel.build(map, self);
    }

    public UploadPartCopyRequest setPartNumber(Long partNumber) {
        this.partNumber = partNumber;
        return this;
    }
    public Long getPartNumber() {
        return this.partNumber;
    }

    public UploadPartCopyRequest setUploadId(String uploadId) {
        this.uploadId = uploadId;
        return this;
    }
    public String getUploadId() {
        return this.uploadId;
    }

}
