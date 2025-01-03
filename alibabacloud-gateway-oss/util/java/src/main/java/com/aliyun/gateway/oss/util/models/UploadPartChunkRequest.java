// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class UploadPartChunkRequest extends TeaModel {
    @NameInMap("body")
    public java.io.InputStream body;

    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("chunkNumber")
    public Long chunkNumber;

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

    public static UploadPartChunkRequest build(java.util.Map<String, ?> map) throws Exception {
        UploadPartChunkRequest self = new UploadPartChunkRequest();
        return TeaModel.build(map, self);
    }

    public UploadPartChunkRequest setBody(java.io.InputStream body) {
        this.body = body;
        return this;
    }
    public java.io.InputStream getBody() {
        return this.body;
    }

    public UploadPartChunkRequest setChunkNumber(Long chunkNumber) {
        this.chunkNumber = chunkNumber;
        return this;
    }
    public Long getChunkNumber() {
        return this.chunkNumber;
    }

    public UploadPartChunkRequest setPartUploadId(String partUploadId) {
        this.partUploadId = partUploadId;
        return this;
    }
    public String getPartUploadId() {
        return this.partUploadId;
    }

    public UploadPartChunkRequest setUploadId(String uploadId) {
        this.uploadId = uploadId;
        return this;
    }
    public String getUploadId() {
        return this.uploadId;
    }

}
