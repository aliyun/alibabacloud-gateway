// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class StartPartUploadRequest extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("chunkSize")
    public Long chunkSize;

    @NameInMap("encoding-type")
    public String encodingType;

    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("partNumber")
    public Long partNumber;

    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("partSize")
    public Long partSize;

    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("uploadId")
    public String uploadId;

    public static StartPartUploadRequest build(java.util.Map<String, ?> map) throws Exception {
        StartPartUploadRequest self = new StartPartUploadRequest();
        return TeaModel.build(map, self);
    }

    public StartPartUploadRequest setChunkSize(Long chunkSize) {
        this.chunkSize = chunkSize;
        return this;
    }
    public Long getChunkSize() {
        return this.chunkSize;
    }

    public StartPartUploadRequest setEncodingType(String encodingType) {
        this.encodingType = encodingType;
        return this;
    }
    public String getEncodingType() {
        return this.encodingType;
    }

    public StartPartUploadRequest setPartNumber(Long partNumber) {
        this.partNumber = partNumber;
        return this;
    }
    public Long getPartNumber() {
        return this.partNumber;
    }

    public StartPartUploadRequest setPartSize(Long partSize) {
        this.partSize = partSize;
        return this;
    }
    public Long getPartSize() {
        return this.partSize;
    }

    public StartPartUploadRequest setUploadId(String uploadId) {
        this.uploadId = uploadId;
        return this;
    }
    public String getUploadId() {
        return this.uploadId;
    }

}
