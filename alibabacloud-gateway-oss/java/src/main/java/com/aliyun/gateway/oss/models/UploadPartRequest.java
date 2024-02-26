// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class UploadPartRequest extends TeaModel {
    /**
     * <p>The request body.</p>
     */
    @NameInMap("body")
    public java.io.InputStream body;

    /**
     * <p>The number that identifies a part. </p>
     * <br>
     * <p>Valid values: 1 to 10000.</p>
     * <br>
     * <p>The size of a part ranges from 100 KB to 5 GB. </p>
     * <p>> In multipart upload, each part except the last part must be larger than or equal to 100 KB in size. When you call the UploadPart operation, the size of each part is not verified because not all parts have been uploaded and OSS does not know which part is the last part. The size of each part is verified only when you call CompleteMultipartUpload.</p>
     */
    @NameInMap("partNumber")
    public Long partNumber;

    /**
     * <p>The ID that identifies the object to which the part that you want to upload belongs.</p>
     */
    @NameInMap("uploadId")
    public String uploadId;

    public static UploadPartRequest build(java.util.Map<String, ?> map) throws Exception {
        UploadPartRequest self = new UploadPartRequest();
        return TeaModel.build(map, self);
    }

    public UploadPartRequest setBody(java.io.InputStream body) {
        this.body = body;
        return this;
    }
    public java.io.InputStream getBody() {
        return this.body;
    }

    public UploadPartRequest setPartNumber(Long partNumber) {
        this.partNumber = partNumber;
        return this;
    }
    public Long getPartNumber() {
        return this.partNumber;
    }

    public UploadPartRequest setUploadId(String uploadId) {
        this.uploadId = uploadId;
        return this;
    }
    public String getUploadId() {
        return this.uploadId;
    }

}
