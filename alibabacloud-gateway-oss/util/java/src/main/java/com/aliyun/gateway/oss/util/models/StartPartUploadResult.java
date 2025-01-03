// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class StartPartUploadResult extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>bucket1</p>
     */
    @NameInMap("Bucket")
    public String bucket;

    /**
     * <strong>example:</strong>
     * <p>url</p>
     */
    @NameInMap("EncodingType")
    public String encodingType;

    /**
     * <strong>example:</strong>
     * <p>multipart.data</p>
     */
    @NameInMap("Key")
    public String key;

    /**
     * <strong>example:</strong>
     * <p>89714B59EF29136807096C6AEB0382730EDA80099D22216E4FEDE45E2B4EC622FC91ED6717B413A9B0C2A4**********</p>
     */
    @NameInMap("PartUploadId")
    public String partUploadId;

    /**
     * <strong>example:</strong>
     * <p>8C108F2EDDCE4C7E946441**********</p>
     */
    @NameInMap("UploadId")
    public String uploadId;

    public static StartPartUploadResult build(java.util.Map<String, ?> map) throws Exception {
        StartPartUploadResult self = new StartPartUploadResult();
        return TeaModel.build(map, self);
    }

    public StartPartUploadResult setBucket(String bucket) {
        this.bucket = bucket;
        return this;
    }
    public String getBucket() {
        return this.bucket;
    }

    public StartPartUploadResult setEncodingType(String encodingType) {
        this.encodingType = encodingType;
        return this;
    }
    public String getEncodingType() {
        return this.encodingType;
    }

    public StartPartUploadResult setKey(String key) {
        this.key = key;
        return this;
    }
    public String getKey() {
        return this.key;
    }

    public StartPartUploadResult setPartUploadId(String partUploadId) {
        this.partUploadId = partUploadId;
        return this;
    }
    public String getPartUploadId() {
        return this.partUploadId;
    }

    public StartPartUploadResult setUploadId(String uploadId) {
        this.uploadId = uploadId;
        return this;
    }
    public String getUploadId() {
        return this.uploadId;
    }

}
